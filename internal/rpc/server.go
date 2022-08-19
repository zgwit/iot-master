package rpc

import (
	"context"
	"iot-master/plugin"
)

type serverServer struct{}

func (*serverServer) List(ctx context.Context, list *plugin.List) (*plugin.Buffer, error) {
	//TODO implement me
	panic("implement me")
}

func (*serverServer) Get(ctx context.Context, i *plugin.Id) (*plugin.Buffer, error) {
	//TODO implement me
	panic("implement me")
}

func (*serverServer) Open(ctx context.Context, i *plugin.Id) (*plugin.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (*serverServer) Close(ctx context.Context, i *plugin.Id) (*plugin.Empty, error) {
	//TODO implement me
	panic("implement me")
}
