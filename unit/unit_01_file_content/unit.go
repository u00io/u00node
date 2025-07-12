package unit01filecontent

import (
	"time"

	"github.com/u00io/gomisc/logger"
	unit00base "github.com/u00io/u00node/unit/unit_00_base"
)

type Unit01FileContent struct {
	unit00base.Unit
}

func New() unit00base.IUnit {
	var c Unit01FileContent
	c.Init(&c)
	return &c
}

func (c *Unit01FileContent) Tick() {
	logger.Println("Unit01FileContent Tick")
	c.SetValue("value", time.Now().Format(time.RFC3339))
}
