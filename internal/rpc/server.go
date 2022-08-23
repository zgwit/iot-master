package rpc

import (
	"context"
	"errors"
	"github.com/zgwit/iot-master/internal/core"
	"github.com/zgwit/iot-master/internal/db"

	"github.com/zgwit/iot-master/model"
	"github.com/zgwit/iot-master/plugin"
)

type serverServer struct{}

func (ss *serverServer) List(ctx context.Context, list *plugin.List) (*plugin.Buffer, error) {
	datum, err := db.Search[model.Server]([]string{"Name"}, list.Keyword, int(list.Skip), int(list.Limit))
	if err != nil {
		return nil, err
	}
	return jsonMarshalBuffer(datum)
}

func (ss *serverServer) Get(ctx context.Context, id *plugin.Id) (*plugin.Buffer, error) {
	var server model.Server
	err := db.Store().Get(id.Value, &server)
	if err != nil {
		return nil, err
	}
	return jsonMarshalBuffer(&server)
}

func (ss *serverServer) Open(ctx context.Context, id *plugin.Id) (*plugin.Empty, error) {
	svr := core.GetServer(id.Value)
	if svr == nil {
		return nil, errors.New("找不到服务器")
	}
	return nil, svr.Instance.Open()
}

func (ss *serverServer) Close(ctx context.Context, id *plugin.Id) (*plugin.Empty, error) {
	dev := core.GetServer(id.Value)
	if dev == nil {
		return nil, errors.New("找不到服务器")
	}
	return nil, dev.Instance.Close()
}

func (ss *serverServer) Enable(ctx context.Context, id *plugin.Id) (*plugin.Empty, error) {
	var server model.Server
	err := db.Store().Get(id.Value, &server)
	if err != nil {
		return nil, err
	}
	server.Disabled = false
	err = db.Store().Update(server.Id, &server)
	if err != nil {
		return nil, err
	}

	go func() {
		_, _ = ss.Open(ctx, id)
	}()

	return nil, err
}

func (ss *serverServer) Disable(ctx context.Context, id *plugin.Id) (*plugin.Empty, error) {
	var server model.Server
	err := db.Store().Get(id.Value, &server)
	if err != nil {
		return nil, err
	}
	server.Disabled = true
	err = db.Store().Update(server.Id, &server)
	if err != nil {
		return nil, err
	}

	go func() {
		_, _ = ss.Close(ctx, id)
	}()

	return nil, err
}
