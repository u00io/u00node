package system

import (
	"time"

	unit00base "github.com/u00io/u00node/unit/unit_00_base"
)

type System struct {
	client *U00

	units []unit00base.IUnit
}

var Instance *System

func NewSystem() *System {
	var c System
	c.client = NewU00()
	c.units = make([]unit00base.IUnit, 0)
	return &c
}

func (c *System) Start() {
	c.client.Run()
	c.LoadDefaultConfig()
	c.startAllUnits()
	go c.thWork()
}

func (c *System) Stop() {
}

func (c *System) startAllUnits() {
	for _, unit := range c.units {
		unit.Start()
	}
}

func (c *System) LoadDefaultConfig() {
	u01 := createUnitByType("unit01filecontent")
	c.units = append(c.units, u01)

}

func (c *System) Test() {
	c.client.WriteValue("Test Value")
}

func (c *System) thWork() {
	for {
		c.SendValues()
		time.Sleep(1 * time.Second)
	}
}

func (c *System) SendValues() {
	for _, unit := range c.units {
		value := unit.GetValue("value")
		if value != "" {
			c.client.WriteValue(value)
		}
	}
}
