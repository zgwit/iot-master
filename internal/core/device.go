package core

import (
	"encoding/json"
)

var Devices Map[Device]

type Device struct {
	Id     string
	Values map[string]any
	Status map[string]any
}

func NewDevice(id string) *Device {
	return &Device{
		Id: id,
	}
}

func (d *Device) Assign(points map[string]any) error {
	data, err := json.Marshal(points)
	if err != nil {
		return err
	}
	return Publish("/device/"+d.Id+"/command/assign", data)
}

func (d *Device) Refresh() error {
	return Publish("/device/"+d.Id+"/command/refresh", []byte(""))
}

//func (d *Device) Status() error {
//	broker.MQTT.Publish("/device/"+d.Id+"/command/status", 0, false, "")
//	return nil
//}
