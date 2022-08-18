package rpc

import (
	"context"
	"encoding/json"

	"iot-master/model"
	"iot-master/plugin"
)

type deviceServer struct{}

func (*deviceServer) List(ctx context.Context, list *plugin.List) (*plugin.Buffer, error) {
	//TODO implement me

	var datum []model.Device
	err := db.Engine.Limit(int(list.Limit), int(list.Skip)).Find(&datum)
	if err != nil {
		return nil, err
	}

	buf, err := json.Marshal(datum)
	if err != nil {
		return nil, err
	}

	return &plugin.Buffer{Value: buf}, nil
}

func (*deviceServer) Get(ctx context.Context, i *plugin.Int64) (*plugin.Buffer, error) {
	//TODO implement me
	panic("implement me")
}

func (*deviceServer) Open(ctx context.Context, i *plugin.Int64) (*plugin.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (*deviceServer) Close(ctx context.Context, i *plugin.Int64) (*plugin.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (*deviceServer) Execute(ctx context.Context, command *plugin.Command) (*plugin.Empty, error) {
	//TODO implement me
	panic("implement me")
}
