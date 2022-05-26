package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/db"
	"github.com/zgwit/iot-master/master"
	"github.com/zgwit/iot-master/model"
	"golang.org/x/net/websocket"
)

func linkList(ctx *gin.Context) {
	var links []*model.LinkEx

	var body paramSearchEx
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		replyError(ctx, err)
		return
	}

	query := body.toQuery()
	query.Select("link.*, " + //TODO 只返回需要的字段
		" 0 as running, tunnel.name as tunnel")
	query.Join("LEFT", "tunnel", "link.tunnel_id=tunnel.id")

	cnt, err := query.FindAndCount(&links)
	if err != nil {
		replyError(ctx, err)
		return
	}
	for _, lnk := range links {
		d := master.GetLink(lnk.Id)
		if d != nil {
			lnk.Running = d.Instance.Running()
		}
	}

	replyList(ctx, links, cnt)
}


func linkDetail(ctx *gin.Context) {
	var link model.LinkEx
	has, err := db.Engine.ID(ctx.GetInt64("id")).Get(&link.Link)
	if err != nil {
		replyError(ctx, err)
		return
	}
	if !has {
		replyFail(ctx, "记录不存在")
		return
	}
	d := master.GetLink(link.Id)
	if d != nil {
		link.Running = d.Instance.Running()
	}
	replyOk(ctx, link)
}

func afterLinkDelete(id int64) error {
	return master.RemoveLink(id)
}

func afterLinkDisable(id int64) error {
	return master.RemoveLink(id)
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
