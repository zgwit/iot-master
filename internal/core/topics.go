package core

import (
	"encoding/json"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/zgwit/iot-master/internal/db"
	"github.com/zgwit/iot-master/model"
	"strings"
)

func subscribeTopics(client mqtt.Client) {
	//网关事件
	client.Subscribe("/gateway/+/event", 0, func(client mqtt.Client, message mqtt.Message) {
		gateway := strings.Split(message.Topic(), "/")[2]
		var event model.Event
		_ = json.Unmarshal(message.Payload(), &event)
		_, _ = db.Engine.InsertOne(&model.GatewayEvent{Event: event, GatewayId: gateway})
	})

	//网关状态
	client.Subscribe("/gateway/+/status", 0, func(client mqtt.Client, message mqtt.Message) {
		//strings.IndexRune(message.Topic(), '/')
		//gateway := strings.Split(message.Topic(), "/")[2]

	})

	//通道事件
	client.Subscribe("/tunnel/+/event", 0, func(client mqtt.Client, message mqtt.Message) {
		tunnel := strings.Split(message.Topic(), "/")[2]
		var event model.Event
		_ = json.Unmarshal(message.Payload(), &event)
		_, _ = db.Engine.InsertOne(&model.TunnelEvent{Event: event, TunnelId: tunnel})
	})

	//通道状态
	client.Subscribe("/tunnel/+/status", 0, func(client mqtt.Client, message mqtt.Message) {
		tunnel := strings.Split(message.Topic(), "/")[2]
		var status model.Status
		_ = json.Unmarshal(message.Payload(), &status)
		TunnelStatus.Store(tunnel, &status)
	})

	//服务事件
	client.Subscribe("/server/+/event", 0, func(client mqtt.Client, message mqtt.Message) {
		server := strings.Split(message.Topic(), "/")[2]
		var event model.Event
		_ = json.Unmarshal(message.Payload(), &event)
		_, _ = db.Engine.InsertOne(&model.ServerEvent{Event: event, ServerId: server})
	})

	//服务状态
	client.Subscribe("/server/+/status", 0, func(client mqtt.Client, message mqtt.Message) {
		server := strings.Split(message.Topic(), "/")[2]
		var status model.Status
		_ = json.Unmarshal(message.Payload(), &status)
		ServerStatus.Store(server, &status)
	})

	//新通道
	client.Subscribe("/server/+/tunnel", 0, func(client mqtt.Client, message mqtt.Message) {
		server := strings.Split(message.Topic(), "/")[2]
		var tunnel model.Tunnel
		tunnel.ServerId = server
		_ = json.Unmarshal(message.Payload(), &tunnel)
		//ServerStatus.Store(server, &tunnel)
		_, _ = db.Engine.InsertOne(&tunnel)
	})

	//设备事件
	client.Subscribe("/device/+/event", 0, func(client mqtt.Client, message mqtt.Message) {
		device := strings.Split(message.Topic(), "/")[2]
		var event model.Event
		_ = json.Unmarshal(message.Payload(), &event)
		_, _ = db.Engine.InsertOne(&model.DeviceEvent{Event: event, DeviceId: device})
	})

	//设备状态
	client.Subscribe("/device/+/status", 0, func(client mqtt.Client, message mqtt.Message) {
		device := strings.Split(message.Topic(), "/")[2]
		var status model.Status
		_ = json.Unmarshal(message.Payload(), &status)
		if dev := Devices.Load(device); dev != nil {
			dev.Status = status
		} else {
			Devices.Store(device, &Device{
				Id:     device,
				Values: make(map[string]any),
				Status: status,
			})
		}
	})

	//查询项目状态
	client.Subscribe("/project/+/command/status", 0, func(client mqtt.Client, message mqtt.Message) {
		project := strings.Split(message.Topic(), "/")[2]
		if prj := Projects.Load(project); prj != nil {
			payload, _ := json.Marshal(prj.Status)
			client.Publish("/project/"+project+"/status", 0, false, payload)
		}
	})

	//刷新项目数据
	client.Subscribe("/project/+/command/refresh", 0, func(client mqtt.Client, message mqtt.Message) {
		project := strings.Split(message.Topic(), "/")[2]
		if prj := Projects.Load(project); prj != nil {
			_ = prj.Refresh()
		}
	})

	//赋值项目数据
	client.Subscribe("/project/+/command/assign", 0, func(client mqtt.Client, message mqtt.Message) {
		project := strings.Split(message.Topic(), "/")[2]
		points := make(map[string]any)
		_ = json.Unmarshal(message.Payload(), &points)
		if prj := Projects.Load(project); prj != nil {
			_ = prj.Assign(points)
		}
	})

}
