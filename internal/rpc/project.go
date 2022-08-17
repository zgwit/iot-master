package rpc

import (
	"context"
	"iot-master/plugin"
)

type projectServer struct{}

func (*projectServer) List(ctx context.Context, list *plugin.List) (*plugin.Buffer, error) {
	//TODO implement me
	panic("implement me")
}

func (*projectServer) Get(ctx context.Context, i *plugin.Int64) (*plugin.Buffer, error) {
	//TODO implement me
	panic("implement me")
}

func (*projectServer) Open(ctx context.Context, i *plugin.Int64) (*plugin.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (*projectServer) Close(ctx context.Context, i *plugin.Int64) (*plugin.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (*projectServer) Import(server plugin.Project_ImportServer) error {
	//TODO implement me
	panic("implement me")
}

func (*projectServer) Export(s *plugin.String, server plugin.Project_ExportServer) error {
	//TODO implement me
	panic("implement me")
}
