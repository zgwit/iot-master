package rpc

import (
	"context"
	"errors"
	"github.com/zgwit/iot-master/internal/core"
	"github.com/zgwit/iot-master/internal/db"

	"github.com/zgwit/iot-master/model"
	"github.com/zgwit/iot-master/plugin"
)

type deviceServer struct{}

func (ds *deviceServer) List(ctx context.Context, list *plugin.List) (*plugin.Buffer, error) {
	datum, err := db.Search[model.Device]([]string{"Name"}, list.Keyword, int(list.Skip), int(list.Limit))
	if err != nil {
		return nil, err
	}
	return jsonMarshalBuffer(datum)
}

func (ds *deviceServer) Get(ctx context.Context, id *plugin.Id) (*plugin.Buffer, error) {
	var device model.Device
	err := db.Store().Get(id.Value, &device)
	if err != nil {
		return nil, err
	}
	return jsonMarshalBuffer(&device)
}

func (ds *deviceServer) Open(ctx context.Context, id *plugin.Id) (*plugin.Empty, error) {
	dev := core.GetDevice(id.Value)
	if dev == nil {
		return nil, errors.New("找不到设备")
	}
	return nil, dev.Start()
}

func (ds *deviceServer) Close(ctx context.Context, id *plugin.Id) (*plugin.Empty, error) {
	dev := core.GetDevice(id.Value)
	if dev == nil {
		return nil, errors.New("找不到设备")
	}
	return nil, dev.Stop()
}

func (ds *deviceServer) Enable(ctx context.Context, id *plugin.Id) (*plugin.Empty, error) {
	var device model.Device
	err := db.Store().Get(id.Value, &device)
	if err != nil {
		return nil, err
	}
	device.Disabled = false
	err = db.Store().Update(device.Id, &device)
	if err != nil {
		return nil, err
	}

	go func() {
		_, _ = ds.Open(ctx, id)
	}()

	return nil, err
}

func (ds *deviceServer) Disable(ctx context.Context, id *plugin.Id) (*plugin.Empty, error) {
	var device model.Device
	err := db.Store().Get(id.Value, &device)
	if err != nil {
		return nil, err
	}
	device.Disabled = true
	err = db.Store().Update(device.Id, &device)
	if err != nil {
		return nil, err
	}

	go func() {
		_, _ = ds.Close(ctx, id)
	}()

	return nil, err
}

func (ds *deviceServer) Refresh(ctx context.Context, id *plugin.Id) (*plugin.Empty, error) {
	dev := core.GetDevice(id.Value)
	if dev == nil {
		return nil, errors.New("找不到设备")
	}
	return nil, dev.Refresh()
}

func (ds *deviceServer) Context(ctx context.Context, id *plugin.Id) (*plugin.Buffer, error) {
	dev := core.GetDevice(id.Value)
	if dev == nil {
		return nil, errors.New("找不到设备")
	}
	return jsonMarshalBuffer(dev.Context)
}

func (ds *deviceServer) Execute(ctx context.Context, command *plugin.DeviceCommand) (*plugin.Empty, error) {
	dev := core.GetDevice(command.Id)
	if dev == nil {
		return nil, errors.New("找不到设备")
	}
	//参数类型转化
	args := make([]interface{}, 0)
	for _, v := range command.Arguments {
		args = append(args, v)
	}
	return nil, dev.Execute(command.Command, args)
}
