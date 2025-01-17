package nathanielc

import (
	"github.com/go-gl/mathgl/mgl64"
	"github.com/golang/glog"
	"github.com/nathanielc/avi"
)

func init() {
	avi.RegisterPilot("constance", NewConstance)
}

type ConstancePilot struct {
	avi.GenericPilot
	moving int
}

func NewConstance() avi.Pilot {
	return &ConstancePilot{}
}

func (self *ConstancePilot) Tick(tick int64) {
	if self.moving > 20000 {
		return
	}

	for _, engine := range self.Engines {
		err := engine.PowerOn(1.0)
		if err != nil {
			if glog.V(3) {
				glog.Infoln("Failed to power engines", err)
			}
		}
	}

	for _, thruster := range self.Thrusters {
		err := thruster.Thrust(mgl64.Vec3{0, 0, 1e4})
		if err != nil {
			if glog.V(3) {
				glog.Infoln("Failed to thrust", err)
			}
		} else {
			self.moving++
		}
	}

}
