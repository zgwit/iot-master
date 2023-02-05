package core

var Devices Map[Device]

type Device struct {
	Id     string
	Values map[string]any
	Status map[string]any
}
