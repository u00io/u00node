package system

type UnitState struct {
	Id       string
	UnitType string
	Value    string
}

type State struct {
	Units []UnitState
}
