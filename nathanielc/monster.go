package nathanielc

import (
	"github.com/go-gl/mathgl/mgl64"
	"github.com/golang/glog"
	"github.com/nathanielc/avi"
	"github.com/nathanielc/avi/nav"
)

func init() {
	avi.RegisterPilot("monster", NewMonster)
}

type MonsterPilot struct {
	avi.GenericPilot
	dir           mgl64.Vec3
	fired         bool
	navComputer   *nav.Nav
	cooldownTicks int64
	target        avi.ID
	targetI       velPoint
	targetF       velPoint
	ctlp          avi.ID
}

func NewMonster() avi.Pilot {
	return &MonsterPilot{
		dir:           mgl64.Vec3{1, 1, 1},
		cooldownTicks: 1,
		target:        avi.NilID,
		ctlp:          avi.NilID,
	}
}

func (self *MonsterPilot) Tick(tick int64) {
	if self.navComputer == nil {
		self.navComputer = nav.NewNav(self.Thrusters)
	}
	for _, engine := range self.Engines {
		err := engine.PowerOn(1.0)
		if err != nil {
			if glog.V(4) {
				glog.Infoln("Failed to power engines", err)
			}
		}
	}
	scan, err := self.Sensors[0].Scan()
	if err != nil {
		if glog.V(4) {
			glog.Infoln("Failed to scan", err)
		}
		return
	}
	defer scan.Done()
	err = self.navComputer.Tick(scan.Position, scan.Velocity)
	if err != nil {
		if glog.V(4) {
			glog.Infoln("Failed to navigate", err)
		}
	}
	if glog.V(4) {
		glog.Infoln("Monster", scan.Health, scan.Position, scan.Velocity.Len())
	}
	self.navCtlP(scan)

	self.fire(tick, scan)

}

func (self *MonsterPilot) navCtlP(scan avi.ScanResult) {

	// Find Control Point
	if !ctlpExists(self.ctlp, scan.ControlPoints) {
		points := 0.0
		for id, ctlp := range scan.ControlPoints {
			p := ctlp.Points
			if p > points {
				points = p
				self.ctlp = id
			}
		}
	}

	if !ctlpExists(self.ctlp, scan.ControlPoints) {
		return
	}

	ctlp := scan.ControlPoints[self.ctlp]

	distance := ctlp.Position.Sub(scan.Position).Len()
	tolerance := (distance - ctlp.Influence) / 2.0
	if tolerance < 10 {
		tolerance = 30
	}
	bias := mgl64.Vec3{0, 0, 1}
	wp := nav.Waypoint{
		Position:  ctlp.Position.Add(bias.Mul(ctlp.Radius + scan.Radius + tolerance)),
		MaxSpeed:  tolerance * 0.4,
		Tolerance: tolerance,
	}
	self.navComputer.SetWaypoint(wp)
}

func (self *MonsterPilot) fire(tick int64, scan avi.ScanResult) {

	//Find target ship
	if !shipExists(self.target, scan.Ships) {
		distance := 0.0
		for id, ship := range scan.Ships {
			if ship.Fleet == self.Fleet {
				continue
			}
			d := ship.Position.Sub(scan.Position).Len()
			if d < distance || distance == 0 {
				distance = d
				self.target = id
				self.targetI = velPoint{}
			}
		}
	}
	if !shipExists(self.target, scan.Ships) {
		return
	}

	target := scan.Ships[self.target]
	targetPos := target.Position
	targetVel := target.Velocity

	if avi.LengthSq(targetPos.Sub(scan.Position)) > 1e6 {
		self.target = avi.NilID
		if glog.V(3) {
			glog.Infoln("Target is too far away choosing another target")
		}
	}

	if tick%self.cooldownTicks == 0 {
		self.targetF.tick = tick
		self.targetF.velocity = targetVel
		for _, weapon := range self.Weapons {
			vel := weapon.GetAmmoVel()

			acc := self.targetF.velocity.
				Sub(self.targetI.velocity).
				Mul(1.0 / (avi.SecondsPerTick * float64(self.targetF.tick-self.targetI.tick)))

			if glog.V(4) {
				glog.Infoln("Acc: ", acc)
			}

			deltaPos := targetPos.Sub(scan.Position)
			deltaVel := targetVel.Sub(scan.Velocity)
			time := calcT(deltaPos, deltaVel, vel)
			if time < 0 {
				if glog.V(3) {
					glog.Infoln("Target out of range", time)
				}
				continue
			}

			dir := deltaVel.Add(deltaPos.Mul(1 / time)).Add(acc.Mul(time * 0.5))

			if glog.V(3) {
				glog.Infoln(dir, dir.Len())
			}

			err := weapon.Fire(dir)
			if err != nil {
				if glog.V(3) {
					glog.Infoln("Failed to fire", err)
				}
			}
			self.cooldownTicks = weapon.GetCoolDownTicks()
		}
		self.targetI.tick = tick
		self.targetI.velocity = targetVel
	}
}
