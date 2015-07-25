package main

import (
	"compress/gzip"
	"flag"
	"github.com/golang/glog"
	"github.com/nathanielc/avi"
	_ "github.com/nathanielc/avi/ships"
	"os"
	"runtime"
)

var partsFile = flag.String("parts", "parts.yaml", "YAML file where available parts are defined")
var mapFile = flag.String("map", "map.yaml", "YAML file that defines the map")
var saveFile = flag.String("save", "save.avi", "Where to save the simulation data")

func main() {

	flag.Parse()

	runtime.GOMAXPROCS(runtime.NumCPU())

	// Load parts
	pf, err := os.Open(*partsFile)
	if err != nil {
		glog.Error(err)
		return
	}
	defer pf.Close()
	parts, err := avi.LoadPartsFromFile(pf)
	if err != nil {
		glog.Error(err)
		return
	}

	// Load map
	mf, err := os.Open(*mapFile)
	if err != nil {
		glog.Error(err)
		return
	}
	defer mf.Close()
	mp, err := avi.LoadMapFromFile(mf)
	if err != nil {
		glog.Error(err)
		return
	}

	// Load fleets
	fleetFiles := flag.Args()
	fleets := make([]*avi.FleetConf, 0, len(fleetFiles))
	for _, fleetFile := range fleetFiles {
		ff, err := os.Open(fleetFile)
		if err != nil {
			glog.Error(err)
			return
		}
		defer ff.Close()
		fleet, err := avi.LoadFleetFromFile(ff)
		if err != nil {
			glog.Error(err)
			return
		}

		fleets = append(fleets, fleet)
	}

	f, _ := os.Create(*saveFile)
	defer f.Close()
	g := gzip.NewWriter(f)
	defer g.Close()
	stream := avi.NewStream(g)

	sim, err := avi.NewSimulation(mp, parts, fleets, stream)
	if err != nil {
		glog.Error(err)
		return
	}
	sim.Start()
}
