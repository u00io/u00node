package unit01filecontent

import (
	"fmt"

	"github.com/u00io/gomisc/logger"
	unit00base "github.com/u00io/u00node/unit/unit_00_base"
)

type Unit01FileContent struct {
	unit00base.Unit

	counter int
}

func New() unit00base.IUnit {
	var c Unit01FileContent
	c.SetType("unit01filecontent")
	c.Init(&c)
	return &c
}

func (c *Unit01FileContent) Tick() {
	logger.Println("Unit01FileContent Tick")
	c.SetValue("value", "File Content: "+fmt.Sprint(c.counter))
	c.counter++
}
