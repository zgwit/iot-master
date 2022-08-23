package rpc

import (
	"context"
	"github.com/zgwit/iot-master/internal/db"

	"github.com/zgwit/iot-master/model"
	"github.com/zgwit/iot-master/plugin"
)

type productServer struct{}

func (*productServer) List(ctx context.Context, list *plugin.List) (*plugin.Buffer, error) {
	datum, err := db.Search[model.Product]([]string{"Name"}, list.Keyword, int(list.Skip), int(list.Limit))
	if err != nil {
		return nil, err
	}
	return jsonMarshalBuffer(datum)
}

func (*productServer) Get(ctx context.Context, s *plugin.String) (*plugin.Buffer, error) {
	var product model.Product
	err := db.Store().Get(s.Value, &product)
	if err != nil {
		return nil, err
	}
	return jsonMarshalBuffer(&product)
}

func (*productServer) Import(server plugin.Product_ImportServer) error {
	//TODO implement me
	panic("implement me")
}

func (*productServer) Export(s *plugin.String, server plugin.Product_ExportServer) error {
	//TODO implement me
	panic("implement me")
}
