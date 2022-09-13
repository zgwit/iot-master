package core

import (
	"encoding/json"
	"github.com/zgwit/iot-master/internal/broker"
	"github.com/zgwit/iot-master/model"
)

type Device struct {
	Id     string
	Values map[string]any
	Status model.Status
}

func (d *Device) Assign(points map[string]any) error {
	data, _ := json.Marshal(points)
	broker.MQTT.Publish("/device/"+d.Id+"/command/assign", 0, false, data)
	return nil
}

func (d *Device) Refresh() error {
	broker.MQTT.Publish("/device/"+d.Id+"/command/refresh", 0, false, "")
	return nil
}

//func (d *Device) Status() error {
//	broker.MQTT.Publish("/device/"+d.Id+"/command/status", 0, false, "")
//	return nil
//}
