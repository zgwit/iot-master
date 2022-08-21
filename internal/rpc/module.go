package rpc

import (
	"context"
	"github.com/zgwit/iot-master/plugin"
)

type moduleServer struct {
}

func (*moduleServer) Register(ctx context.Context, b *plugin.Buffer) (*plugin.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (*moduleServer) Unregister(ctx context.Context, s *plugin.String) (*plugin.Empty, error) {
	//TODO implement me
	panic("implement me")
}
