package rpc

import (
	"context"
	"github.com/zgwit/iot-master/plugin"
)

type userServer struct{}

func (us *userServer) List(ctx context.Context, list *plugin.List) (*plugin.Buffer, error) {
	//TODO implement me
	panic("implement me")
}

func (us *userServer) Get(ctx context.Context, id *plugin.Id) (*plugin.Buffer, error) {
	//TODO implement me
	panic("implement me")
}

func (us *userServer) Enable(ctx context.Context, id *plugin.Id) (*plugin.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (us *userServer) Disable(ctx context.Context, id *plugin.Id) (*plugin.Empty, error) {
	//TODO implement me
	panic("implement me")
}
