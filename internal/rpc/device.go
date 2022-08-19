package rpc

import (
	"context"
	"encoding/json"
	"iot-master/internal/db"

	"iot-master/model"
	"iot-master/plugin"
)

type deviceServer struct{}

func (*deviceServer) List(ctx context.Context, list *plugin.List) (*plugin.Buffer, error) {
	datum, err := db.Search[model.Device]([]string{"Name"}, list.Keyword, int(list.Skip), int(list.Limit))
	if err != nil {
		return nil, err
	}

	buf, err := json.Marshal(datum)
	if err != nil {
		return nil, err
	}

	return &plugin.Buffer{Value: buf}, nil
}

func (*deviceServer) Get(ctx context.Context, i *plugin.Id) (*plugin.Buffer, error) {
	//TODO implement me
	panic("implement me")
}

func (*deviceServer) Open(ctx context.Context, i *plugin.Id) (*plugin.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (*deviceServer) Close(ctx context.Context, i *plugin.Id) (*plugin.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (*deviceServer) Execute(ctx context.Context, command *plugin.DeviceCommand) (*plugin.Empty, error) {
	//TODO implement me
	panic("implement me")
}
