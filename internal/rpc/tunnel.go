package rpc

import (
	"context"
	"errors"
	"github.com/zgwit/iot-master/internal/core"
	"github.com/zgwit/iot-master/internal/db"

	"github.com/zgwit/iot-master/model"
	"github.com/zgwit/iot-master/plugin"
)

type tunnelServer struct{}

func (ts *tunnelServer) List(ctx context.Context, list *plugin.List) (*plugin.Buffer, error) {
	datum, err := db.Search[model.Tunnel]([]string{"Name"}, list.Keyword, int(list.Skip), int(list.Limit))
	if err != nil {
		return nil, err
	}
	return jsonMarshalBuffer(datum)
}

func (ts *tunnelServer) Get(ctx context.Context, id *plugin.Id) (*plugin.Buffer, error) {
	var tunnel model.Tunnel
	err := db.Store().Get(id.Value, &tunnel)
	if err != nil {
		return nil, err
	}
	return jsonMarshalBuffer(&tunnel)
}

func (ts *tunnelServer) Open(ctx context.Context, id *plugin.Id) (*plugin.Empty, error) {
	svr := core.GetTunnel(id.Value)
	if svr == nil {
		return nil, errors.New("找不到通道")
	}
	return nil, svr.Instance.Open()
}

func (ts *tunnelServer) Close(ctx context.Context, id *plugin.Id) (*plugin.Empty, error) {
	dev := core.GetTunnel(id.Value)
	if dev == nil {
		return nil, errors.New("找不到通道")
	}
	return nil, dev.Instance.Close()
}

func (ts *tunnelServer) Enable(ctx context.Context, id *plugin.Id) (*plugin.Empty, error) {
	var tunnel model.Tunnel
	err := db.Store().Get(id.Value, &tunnel)
	if err != nil {
		return nil, err
	}
	tunnel.Disabled = false
	err = db.Store().Update(tunnel.Id, &tunnel)
	if err != nil {
		return nil, err
	}

	go func() {
		_, _ = ts.Open(ctx, id)
	}()

	return nil, err
}

func (ts *tunnelServer) Disable(ctx context.Context, id *plugin.Id) (*plugin.Empty, error) {
	var tunnel model.Tunnel
	err := db.Store().Get(id.Value, &tunnel)
	if err != nil {
		return nil, err
	}
	tunnel.Disabled = true
	err = db.Store().Update(tunnel.Id, &tunnel)
	if err != nil {
		return nil, err
	}

	go func() {
		_, _ = ts.Close(ctx, id)
	}()

	return nil, err
}

func (tunnelServer) Pipe(server plugin.Tunnel_PipeServer) error {
	//TODO implement me
	panic("implement me")
}
