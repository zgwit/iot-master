package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/db"
	"github.com/zgwit/iot-master/master"
	"github.com/zgwit/iot-master/model"
	"golang.org/x/net/websocket"
)


func afterTunnelCreate(data interface{}) error {
	tunnel := data.(*model.Tunnel)
	if !tunnel.Disabled {
		return master.LoadTunnel(tunnel.Id)
	}
	return nil
}

func tunnelDetail(ctx *gin.Context) {
	var tunnel model.TunnelEx
	has, err := db.Engine.ID(ctx.GetInt64("id")).Exist(&tunnel)
	if err != nil {
		replyError(ctx, err)
		return
	}
	if !has {
		replyFail(ctx, "记录存在")
		return
	}
	d := master.GetTunnel(tunnel.Id)
	if d != nil {
		tunnel.Running = d.Instance.Running()
	}
	replyOk(ctx, tunnel)
}

func afterTunnelUpdate(data interface{}) error {
	tunnel := data.(*model.Tunnel)
	_ = master.RemoveTunnel(tunnel.Id)
	return master.LoadTunnel(tunnel.Id)
}

func afterTunnelDelete(data interface{}) error {
	tunnel := data.(*model.Tunnel)
	return master.RemoveTunnel(tunnel.Id)
}

func tunnelStart(ctx *gin.Context) {
	tunnel := master.GetTunnel(ctx.GetInt64("id"))
	if tunnel == nil {
		replyFail(ctx, "not found")
		return
	}
	err := tunnel.Instance.Open()
	if err != nil {
		replyError(ctx, err)
		return
	}

	replyOk(ctx, nil)
}

func tunnelStop(ctx *gin.Context) {
	tunnel := master.GetTunnel(ctx.GetInt64("id"))
	if tunnel == nil {
		replyFail(ctx, "not found")
		return
	}
	err := tunnel.Instance.Close()
	if err != nil {
		replyError(ctx, err)
		return
	}

	replyOk(ctx, nil)
}

func afterTunnelEnable(data interface{}) error {
	tunnel := data.(*model.Tunnel)
	_ = master.RemoveTunnel(tunnel.Id)
	return master.LoadTunnel(tunnel.Id)
}

func afterTunnelDisable(data interface{}) error {
	tunnel := data.(*model.Tunnel)
	return master.RemoveTunnel(tunnel.Id)
}

func tunnelWatch(ctx *gin.Context) {
	tunnel := master.GetTunnel(ctx.GetInt64("id"))
	if tunnel == nil {
		replyFail(ctx, "找不到通道")
		return
	}
	websocket.Handler(func(ws *websocket.Conn) {
		watchAllEvents(ws, tunnel.Instance)
	}).ServeHTTP(ctx.Writer, ctx.Request)
}
