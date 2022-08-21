package rpc

import (
	"context"
	"github.com/zgwit/iot-master/plugin"
)

type sessionServer struct{}

func (*sessionServer) Create(ctx context.Context, s *plugin.String) (*plugin.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (*sessionServer) Get(ctx context.Context, s *plugin.String) (*plugin.Buffer, error) {
	//TODO implement me
	panic("implement me")
}

func (*sessionServer) Delete(ctx context.Context, s *plugin.String) (*plugin.Empty, error) {
	//TODO implement me
	panic("implement me")
}
