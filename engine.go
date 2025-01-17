package avi

import (
	"errors"

	"github.com/go-gl/mathgl/mgl64"
)

type Engine struct {
	partT
	energy        float64
	currentOutput float64
}

// Conf format for loading engines from a file
type EngineConf struct {
	Mass   float64 `yaml:"mass" json:"mass"`
	Radius float64 `yaml:"radius" json:"radius"`
	Energy float64 `yaml:"energy" json:"energy"`
}

func NewEngine001(pos mgl64.Vec3) *Engine {
	return &Engine{
		partT: partT{
			objectT: objectT{
				position: pos,
				mass:     2000,
				radius:   5,
			},
		},
		energy: 100,
	}
}

func NewEngineFromConf(pos mgl64.Vec3, conf EngineConf) *Engine {
	return &Engine{
		partT: partT{
			objectT: objectT{
				position: pos,
				mass:     conf.Mass,
				radius:   conf.Radius,
			},
		},
		energy: conf.Energy,
	}
}

func (self *Engine) getOutput() float64 {
	return self.currentOutput
}

func (self *Engine) PowerOn(power float64) error {
	if power > 1 || power < 0 {
		return errors.New("Power must be between 0 and 1")
	}

	self.currentOutput = self.energy * power
	return nil
}
