package unit02currenttime

import unit00base "github.com/u00io/u00node/unit/unit_00_base"

type Unit02CurrentTime struct {
	unit00base.Unit
}

func New() unit00base.IUnit {
	var c Unit02CurrentTime
	return &c
}
