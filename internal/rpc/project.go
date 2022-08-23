package rpc

import (
	"context"
	"errors"
	"github.com/zgwit/iot-master/internal/core"
	"github.com/zgwit/iot-master/internal/db"

	"github.com/zgwit/iot-master/model"
	"github.com/zgwit/iot-master/plugin"
)

type projectServer struct{}

func (ps *projectServer) List(ctx context.Context, list *plugin.List) (*plugin.Buffer, error) {
	datum, err := db.Search[model.Project]([]string{"Name"}, list.Keyword, int(list.Skip), int(list.Limit))
	if err != nil {
		return nil, err
	}
	return jsonMarshalBuffer(datum)
}

func (ps *projectServer) Get(ctx context.Context, id *plugin.Id) (*plugin.Buffer, error) {
	var project model.Project
	err := db.Store().Get(id.Value, &project)
	if err != nil {
		return nil, err
	}
	return jsonMarshalBuffer(&project)
}

func (ps *projectServer) Open(ctx context.Context, id *plugin.Id) (*plugin.Empty, error) {
	prj := core.GetProject(id.Value)
	if prj == nil {
		return nil, errors.New("找不到设备")
	}
	return nil, prj.Start()
}

func (ps *projectServer) Close(ctx context.Context, id *plugin.Id) (*plugin.Empty, error) {
	prj := core.GetProject(id.Value)
	if prj == nil {
		return nil, errors.New("找不到设备")
	}
	return nil, prj.Stop()
}

func (ps *projectServer) Enable(ctx context.Context, id *plugin.Id) (*plugin.Empty, error) {
	var project model.Project
	err := db.Store().Get(id.Value, &project)
	if err != nil {
		return nil, err
	}
	project.Disabled = false
	err = db.Store().Update(project.Id, &project)
	if err != nil {
		return nil, err
	}

	go func() {
		_, _ = ps.Open(ctx, id)
	}()

	return nil, err
}

func (ps *projectServer) Disable(ctx context.Context, id *plugin.Id) (*plugin.Empty, error) {
	var project model.Project
	err := db.Store().Get(id.Value, &project)
	if err != nil {
		return nil, err
	}
	project.Disabled = true
	err = db.Store().Update(project.Id, &project)
	if err != nil {
		return nil, err
	}

	go func() {
		_, _ = ps.Close(ctx, id)
	}()

	return nil, err
}

func (ps *projectServer) Context(ctx context.Context, id *plugin.Id) (*plugin.Buffer, error) {
	prj := core.GetProject(id.Value)
	if prj == nil {
		return nil, errors.New("找不到项目")
	}
	return jsonMarshalBuffer(prj.Context)
}

func (ps *projectServer) Execute(ctx context.Context, command *plugin.ProjectCommand) (*plugin.Empty, error) {
	prj := core.GetProject(command.Id)
	if prj == nil {
		return nil, errors.New("找不到项目")
	}
	//参数类型转化
	args := make([]interface{}, 0)
	for _, v := range command.Arguments {
		args = append(args, v)
	}
	return nil, prj.Execute(command.Targets, command.Command, args)
}

func (ps *projectServer) Import(server plugin.Project_ImportServer) error {
	//TODO implement me
	panic("implement me")
}

func (ps *projectServer) Export(s *plugin.String, server plugin.Project_ExportServer) error {
	//TODO implement me
	panic("implement me")
}
