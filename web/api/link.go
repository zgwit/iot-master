package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/master"
	"github.com/zgwit/iot-master/model"
	"golang.org/x/net/websocket"
)

func linkList(ctx *gin.Context) {
	links := make([]*model.LinkEx, 0)

	var body paramSearchEx
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		replyError(ctx, err)
		return
	}

	query := body.toQuery()

	query.Join("LEFT", "tunnel", "link.tunnel_id=tunnel.id")

	cnt, err := query.FindAndCount(links)
	if err != nil {
		replyError(ctx, err)
		return
	}
	replyList(ctx, links, cnt)
}

func afterLinkDelete(data interface{}) error {
	link := data.(*model.Link)
	return master.RemoveLink(link.Id)
}

func afterLinkDisable(data interface{}) error {
	link := data.(*model.Link)
	return master.RemoveLink(link.Id)
}

func linkClose(ctx *gin.Context) {
	link := master.GetLink(ctx.GetInt64("id"))
	if link == nil {
		replyFail(ctx, "link not found")
		return
	}
	err := link.Instance.Close()
	if err != nil {
		replyError(ctx, err)
		return
	}

	replyOk(ctx, nil)
}

func linkWatch(ctx *gin.Context) {
	link := master.GetLink(ctx.GetInt64("id"))
	if link == nil {
		replyFail(ctx, "找不到链接")
		return
	}
	websocket.Handler(func(ws *websocket.Conn) {
		watchAllEvents(ws, link.Instance)
	}).ServeHTTP(ctx.Writer, ctx.Request)
}
