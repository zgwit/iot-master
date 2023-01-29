package model

type Model struct {
	ID         string
	Name       string
	Desc       string
	Version    string
	Author     string
	Email      string
	Properties []Property
	Functions  []Function
	Events     []Event
}

type Property struct {
	ID   string
	Name string
	Desc string
	Type string //int float ....
	Unit string
	Mode string //r w rw
}

type Function struct {
	ID     string
	Name   string
	Desc   string
	Async  bool
	Input  []Argument
	Output []Argument
}

type Argument struct {
	Name string
	Desc string
	Type string
	Unit string
}

type Event struct {
	ID     string
	Name   string
	Desc   string
	Type   string //info alert error
	Output []Argument
}
