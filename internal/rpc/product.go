package rpc

import (
	"context"
	"github.com/zgwit/iot-master/plugin"
)

type productServer struct{}

func (*productServer) List(ctx context.Context, list *plugin.List) (*plugin.Buffer, error) {
	//TODO implement me
	panic("implement me")
}

func (*productServer) Get(ctx context.Context, s *plugin.String) (*plugin.Buffer, error) {
	//TODO implement me
	panic("implement me")
}

func (*productServer) Import(server plugin.Product_ImportServer) error {
	//TODO implement me
	panic("implement me")
}

func (*productServer) Export(s *plugin.String, server plugin.Product_ExportServer) error {
	//TODO implement me
	panic("implement me")
}
