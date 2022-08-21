package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/internal/core"
	"github.com/zgwit/iot-master/model"
	"golang.org/x/net/websocket"
)

func afterTunnelCreate(data interface{}) error {
	tunnel := data.(*model.Tunnel)
	if !tunnel.Disabled {
		return core.LoadTunnel(tunnel.Id)
	}
	return nil
}

func afterTunnelDelete(id interface{}) error {
	return core.RemoveTunnel(id.(uint64))
}

func afterTunnelEnable(id interface{}) error {
	_ = core.RemoveTunnel(id.(uint64))
	err := core.LoadTunnel(id.(uint64))
	return err
}

func afterTunnelDisable(id interface{}) error {
	return core.RemoveTunnel(id.(uint64))
}

func tunnelStart(ctx *gin.Context) {
	tunnel := core.GetTunnel(ctx.GetUint64("id"))
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
	tunnel := core.GetTunnel(ctx.GetUint64("id"))
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

func tunnelTransfer(ctx *gin.Context) {
	tunnel := core.GetTunnel(ctx.GetUint64("id"))
	if tunnel == nil {
		replyFail(ctx, "找不到通道")
		return
	}
	websocket.Handler(func(ws *websocket.Conn) {
		ws.PayloadType = websocket.BinaryFrame

		tunnel.Instance.Pipe(ws)
	}).ServeHTTP(ctx.Writer, ctx.Request)
}
