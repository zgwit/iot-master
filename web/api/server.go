package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/db"
	"github.com/zgwit/iot-master/master"
	"github.com/zgwit/iot-master/model"
	"golang.org/x/net/websocket"
)

func serverList(ctx *gin.Context) {
	var servers []*model.ServerEx

	var body paramSearchEx
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		replyError(ctx, err)
		return
	}

	query := body.toQuery()
	query.Select("server.*, " + //TODO 只返回需要的字段
		" 0 as running")
	cnt, err := query.FindAndCount(&servers)
	if err != nil {
		replyError(ctx, err)
		return
	}
	for _, tnl := range servers {
		d := master.GetServer(tnl.Id)
		if d != nil {
			tnl.Running = d.Instance.Running()
		}
	}

	replyList(ctx, servers, cnt)
}

func afterServerCreate(data interface{}) error {
	server := data.(*model.Server)
	if !server.Disabled {
		return master.LoadServer(server.Id)
	}
	return nil
}

func serverDetail(ctx *gin.Context) {
	var server model.ServerEx
	has, err := db.Engine.ID(ctx.GetInt64("id")).Get(&server.Server)
	if err != nil {
		replyError(ctx, err)
		return
	}
	if !has {
		replyFail(ctx, "记录不存在")
		return
	}
	d := master.GetServer(server.Id)
	if d != nil {
		server.Running = d.Instance.Running()
	}
	replyOk(ctx, server)
}

func afterServerUpdate(data interface{}) error {
	server := data.(*model.Server)
	_ = master.RemoveServer(server.Id)
	return master.LoadServer(server.Id)
}

func afterServerDelete(id int64) error {
	return master.RemoveServer(id)
}

func serverStart(ctx *gin.Context) {
	server := master.GetServer(ctx.GetInt64("id"))
	if server == nil {
		replyFail(ctx, "not found")
		return
	}
	err := server.Instance.Open()
	if err != nil {
		replyError(ctx, err)
		return
	}

	replyOk(ctx, nil)
}

func serverStop(ctx *gin.Context) {
	server := master.GetServer(ctx.GetInt64("id"))
	if server == nil {
		replyFail(ctx, "not found")
		return
	}
	err := server.Instance.Close()
	if err != nil {
		replyError(ctx, err)
		return
	}

	replyOk(ctx, nil)
}

func afterServerEnable(id int64) error {
	_ = master.RemoveServer(id)
	return master.LoadServer(id)
}

func afterServerDisable(id int64) error {
	return master.RemoveServer(id)
}

func serverWatch(ctx *gin.Context) {
	server := master.GetServer(ctx.GetInt64("id"))
	if server == nil {
		replyFail(ctx, "找不到通道")
		return
	}
	websocket.Handler(func(ws *websocket.Conn) {
		watchAllEvents(ws, server.Instance)
	}).ServeHTTP(ctx.Writer, ctx.Request)
}
