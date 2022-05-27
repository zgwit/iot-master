package master

import (
	"github.com/zgwit/iot-master/db"
	"github.com/zgwit/iot-master/model"
)

func CreateEvent(Target string, TargetId int64, Event string) {
	_, _ = db.Engine.InsertOne(&model.Event{
		Event:    Event,
		Target:   Target,
		TargetId: TargetId,
	})
}

func CreateUserEvent(UserId int64, Event string) {
	_, _ = db.Engine.InsertOne(&model.Event{
		Event:    Event,
		Target:   "user",
		TargetId: UserId,
	})
}

func CreateTunnelEvent(TunnelId int64, Event string) {
	_, _ = db.Engine.InsertOne(&model.Event{
		Event:    Event,
		Target:   "tunnel",
		TargetId: TunnelId,
	})
}

func CreateServerEvent(ServerId int64, Event string) {
	_, _ = db.Engine.InsertOne(&model.Event{
		Event:    Event,
		Target:   "server",
		TargetId: ServerId,
	})
}

func CreateDeviceEvent(DeviceId int64, Event string) {
	_, _ = db.Engine.InsertOne(&model.Event{
		Event:    Event,
		Target:   "device",
		TargetId: DeviceId,
	})
}

func CreateProjectEvent(ProjectId int64, Event string) {
	_, _ = db.Engine.InsertOne(&model.Event{
		Event:    Event,
		Target:   "project",
		TargetId: ProjectId,
	})
}
