package core

import (
	"encoding/json"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/zgwit/iot-master/v3/internal/db"
	"github.com/zgwit/iot-master/v3/model"
	"strings"
)

func subscribeTopics(client mqtt.Client) {
	//网关事件
	client.Subscribe("/gateway/+/event", 0, func(client mqtt.Client, message mqtt.Message) {
		id := strings.Split(message.Topic(), "/")[2]
		var event model.Event
		_ = json.Unmarshal(message.Payload(), &event)
		_, _ = db.Engine.InsertOne(&model.GatewayEvent{Event: event, GatewayId: id})
	})

	//网关状态
	client.Subscribe("/gateway/+/status", 0, func(client mqtt.Client, message mqtt.Message) {
		id := strings.Split(message.Topic(), "/")[2]
		var status model.Status
		_ = json.Unmarshal(message.Payload(), &status)
		GatewayStatus.Store(id, &status)
	})

	//通道事件
	client.Subscribe("/tunnel/+/event", 0, func(client mqtt.Client, message mqtt.Message) {
		id := strings.Split(message.Topic(), "/")[2]
		var event model.Event
		_ = json.Unmarshal(message.Payload(), &event)
		_, _ = db.Engine.InsertOne(&model.TunnelEvent{Event: event, TunnelId: id})
	})

	//通道状态
	client.Subscribe("/tunnel/+/status", 0, func(client mqtt.Client, message mqtt.Message) {
		id := strings.Split(message.Topic(), "/")[2]
		var status model.Status
		_ = json.Unmarshal(message.Payload(), &status)
		TunnelStatus.Store(id, &status)
	})

	//服务事件
	client.Subscribe("/server/+/event", 0, func(client mqtt.Client, message mqtt.Message) {
		id := strings.Split(message.Topic(), "/")[2]
		var event model.Event
		_ = json.Unmarshal(message.Payload(), &event)
		_, _ = db.Engine.InsertOne(&model.ServerEvent{Event: event, ServerId: id})
	})

	//服务状态
	client.Subscribe("/server/+/status", 0, func(client mqtt.Client, message mqtt.Message) {
		id := strings.Split(message.Topic(), "/")[2]
		var status model.Status
		_ = json.Unmarshal(message.Payload(), &status)
		ServerStatus.Store(id, &status)
	})

	//新通道
	client.Subscribe("/server/+/tunnel", 0, func(client mqtt.Client, message mqtt.Message) {
		id := strings.Split(message.Topic(), "/")[2]
		var tunnel model.Tunnel
		tunnel.ServerId = id
		_ = json.Unmarshal(message.Payload(), &tunnel)
		//ServerStatus.Store(id, &tunnel)
		_, _ = db.Engine.InsertOne(&tunnel)
	})

	//设备事件
	client.Subscribe("/device/+/event", 0, func(client mqtt.Client, message mqtt.Message) {
		id := strings.Split(message.Topic(), "/")[2]
		var event model.Event
		_ = json.Unmarshal(message.Payload(), &event)
		_, _ = db.Engine.InsertOne(&model.DeviceEvent{Event: event, DeviceId: id})
	})

	//设备状态
	client.Subscribe("/device/+/status", 0, func(client mqtt.Client, message mqtt.Message) {
		id := strings.Split(message.Topic(), "/")[2]
		var status model.Status
		_ = json.Unmarshal(message.Payload(), &status)
		if dev := Devices.Load(id); dev != nil {
			dev.Status = status
		} else {
			Devices.Store(id, &Device{Id: id, Values: make(model.Values), Status: status})
		}
	})

	//设备状态
	client.Subscribe("/device/+/values", 0, func(client mqtt.Client, message mqtt.Message) {
		id := strings.Split(message.Topic(), "/")[2]
		var values model.Values
		_ = json.Unmarshal(message.Payload(), &values)
		if dev := Devices.Load(id); dev != nil {
			//dev.Values = values
			for k, v := range values {
				dev.Values[k] = v
			}
		} else {
			Devices.Store(id, &Device{Id: id, Values: values})
		}
	})

	//服务注册
	client.Subscribe("/service/register", 0, func(client mqtt.Client, message mqtt.Message) {
		var service model.Service
		_ = json.Unmarshal(message.Payload(), &service)
		Services.Store(service.Name, &service)
	})

}
