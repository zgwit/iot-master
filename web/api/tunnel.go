package api

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/net/websocket"
	"iot-master/db"
	"iot-master/master"
	"iot-master/model"
)

func tunnelList(ctx *gin.Context) {
	var tunnels []*model.TunnelEx

	var body paramSearchEx
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		replyError(ctx, err)
		return
	}

	query := body.toQuery()
	query.Select("tunnel.*, " + //TODO 只返回需要的字段
		" 0 as running, server.name as server")
	query.Join("LEFT", "server", "tunnel.server_id=server.id")

	cnt, err := query.FindAndCount(&tunnels)
	if err != nil {
		replyError(ctx, err)
		return
	}
	for _, lnk := range tunnels {
		d := master.GetTunnel(lnk.Id)
		if d != nil {
			lnk.Running = d.Instance.Running()
		}
	}

	replyList(ctx, tunnels, cnt)
}

func afterTunnelCreate(data interface{}) error {
	tunnel := data.(*model.Tunnel)
	if !tunnel.Disabled {
		return master.LoadTunnel(tunnel.Id)
	}
	return nil
}

func tunnelDetail(ctx *gin.Context) {
	var tunnel model.TunnelEx
	has, err := db.Engine.ID(ctx.GetInt64("id")).Get(&tunnel.Tunnel)
	if err != nil {
		replyError(ctx, err)
		return
	}
	if !has {
		replyFail(ctx, "记录不存在")
		return
	}
	d := master.GetTunnel(tunnel.Id)
	if d != nil {
		tunnel.Running = d.Instance.Running()
	}
	replyOk(ctx, tunnel)
}

func afterTunnelDelete(id interface{}) error {
	return master.RemoveTunnel(id.(int64))
}

func afterTunnelEnable(id interface{}) error {
	_ = master.RemoveTunnel(id.(int64))
	err := master.LoadTunnel(id.(int64))
	return err
}

func afterTunnelDisable(id interface{}) error {
	return master.RemoveTunnel(id.(int64))
}

func tunnelStart(ctx *gin.Context) {
	tunnel := master.GetTunnel(ctx.GetInt64("id"))
	if tunnel == nil {
		replyFail(ctx, "tunnel not found")
		return
	}
	err := tunnel.Instance.Open()
	if err != nil {
		replyError(ctx, err)
		return
	}

	replyOk(ctx, nil)
}

func tunnelClose(ctx *gin.Context) {
	tunnel := master.GetTunnel(ctx.GetInt64("id"))
	if tunnel == nil {
		replyFail(ctx, "tunnel not found")
		return
	}
	err := tunnel.Instance.Close()
	if err != nil {
		replyError(ctx, err)
		return
	}

	replyOk(ctx, nil)
}

func tunnelWatch(ctx *gin.Context) {
	tunnel := master.GetTunnel(ctx.GetInt64("id"))
	if tunnel == nil {
		replyFail(ctx, "找不到链接")
		return
	}
	websocket.Handler(func(ws *websocket.Conn) {
		watchAllEvents(ws, tunnel.Instance)
	}).ServeHTTP(ctx.Writer, ctx.Request)
}
