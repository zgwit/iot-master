package rpc

import (
	"context"
	"iot-master/plugin"
)

type tunnelServer struct{}

func (*tunnelServer) List(ctx context.Context, list *plugin.List) (*plugin.Buffer, error) {
	//TODO implement me
	panic("implement me")
}

func (*tunnelServer) Get(ctx context.Context, i *plugin.Int64) (*plugin.Buffer, error) {
	//TODO implement me
	panic("implement me")
}

func (*tunnelServer) Open(ctx context.Context, i *plugin.Int64) (*plugin.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (*tunnelServer) Close(ctx context.Context, i *plugin.Int64) (*plugin.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (*tunnelServer) Pipe(server plugin.Tunnel_PipeServer) error {
	//TODO implement me
	panic("implement me")
}
