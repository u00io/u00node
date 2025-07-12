package system

import (
	unit00base "github.com/u00io/u00node/unit/unit_00_base"
	unit01filecontent "github.com/u00io/u00node/unit/unit_01_file_content"
	unit02currenttime "github.com/u00io/u00node/unit/unit_02_current_time"
)

type UnitTypeRecord struct {
	TypeName    string
	Constructor func() unit00base.IUnit
}

type UnitsRegistry struct {
	unitTypes map[string]*UnitTypeRecord
}

func (r *UnitsRegistry) RegisterUnitType(unitType string, constructor func() unit00base.IUnit) {
	var record UnitTypeRecord
	record.TypeName = unitType
	record.Constructor = constructor
	r.unitTypes[unitType] = &record
}

var registry UnitsRegistry

func init() {
	registry.unitTypes = make(map[string]*UnitTypeRecord)
	registry.RegisterUnitType("unit01filecontent", unit01filecontent.New)
	registry.RegisterUnitType("unit02currenttime", unit02currenttime.New)
}

func createUnitByType(unitType string) unit00base.IUnit {
	if record, exists := registry.unitTypes[unitType]; exists {
		return record.Constructor()
	}
	return nil
}
