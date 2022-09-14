package core

import (
	"encoding/json"
	"github.com/zgwit/iot-master/model"
)

func NewDevice(id string) *Device {
	return &Device{
		Id:     id,
		Values: make(map[string]any),
	}
}

type Device struct {
	Id     string
	Values map[string]any
	Status model.Status
}

func (d *Device) Assign(points map[string]any) error {
	data, _ := json.Marshal(points)
	err := Publish("/device/"+d.Id+"/command/assign", data)
	if err != nil {
		return err
	}
	return nil
}

func (d *Device) Refresh() error {
	err := Publish("/device/"+d.Id+"/command/refresh", []byte(""))
	if err != nil {
		return err
	}
	return nil
}

//func (d *Device) Status() error {
//	broker.MQTT.Publish("/device/"+d.Id+"/command/status", 0, false, "")
//	return nil
//}
