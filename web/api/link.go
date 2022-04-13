package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/database"
	"github.com/zgwit/iot-master/log"
	"github.com/zgwit/iot-master/master"
	"github.com/zgwit/iot-master/model"
	"github.com/zgwit/storm/v3/q"
	"golang.org/x/net/websocket"
)

func linkRoutes(app *gin.RouterGroup) {
	app.POST("list", linkList)

	app.GET("event/clear", linkEventClearAll)

	app.Use(parseParamId)
	app.POST(":id/update", linkUpdate)
	app.GET(":id/delete", linkDelete)
	app.GET(":id/close", linkClose)
	app.GET(":id/enable", linkEnable)
	app.GET(":id/disable", linkDisable)
	app.GET(":id/watch", linkWatch)
	app.GET(":id/event", linkEvent)
	app.GET(":id/event/clear", linkEventClear)
}

func linkList(ctx *gin.Context) {
	links, cnt, err := normalSearch(ctx, database.Master, &model.Link{})
	if err != nil {
		replyError(ctx, err)
		return
	}
	replyList(ctx, links, cnt)
}

func linkUpdate(ctx *gin.Context) {
	var link model.Link
	err := ctx.ShouldBindJSON(&link)
	if err != nil {
		replyError(ctx, err)
		return
	}
	link.ID = ctx.GetInt("id")

	err = database.Master.Update(&link)
	if err != nil {
		replyError(ctx, err)
		return
	}

	replyOk(ctx, link)
}

func linkDelete(ctx *gin.Context) {
	link := model.Link{ID: ctx.GetInt("id")}
	err := database.Master.DeleteStruct(&link)
	if err != nil {
		replyError(ctx, err)
		return
	}

	replyOk(ctx, link)
	//关闭
	go func() {
		link := master.GetLink(ctx.GetInt("id"))
		if link == nil {
			return
		}
		err := link.Instance.Close()
		if err != nil {
			log.Error(err)
			return
		}
	}()
}

func linkClose(ctx *gin.Context) {
	link := master.GetLink(ctx.GetInt("id"))
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

func linkEnable(ctx *gin.Context) {
	err := database.Master.UpdateField(&model.Link{ID: ctx.GetInt("id")}, "Disabled", false)
	if err != nil {
		replyError(ctx, err)
		return
	}
	replyOk(ctx, nil)
}

func linkDisable(ctx *gin.Context) {
	err := database.Master.UpdateField(&model.Link{ID: ctx.GetInt("id")}, "Disabled", true)
	if err != nil {
		replyError(ctx, err)
		return
	}
	replyOk(ctx, nil)

	//关闭
	go func() {
		link := master.GetLink(ctx.GetInt("id"))
		if link == nil {
			return
		}
		err := link.Instance.Close()
		if err != nil {
			log.Error(err)
			return
		}
	}()
}

func linkWatch(ctx *gin.Context) {
	link := master.GetLink(ctx.GetInt("id"))
	if link == nil {
		replyFail(ctx, "找不到链接")
		return
	}
	websocket.Handler(func(ws *websocket.Conn) {
		watchAllEvents(ws, link.Instance)
	}).ServeHTTP(ctx.Writer, ctx.Request)
}

func linkEvent(ctx *gin.Context) {
	events, cnt, err := normalSearchById(ctx, database.History, "LinkID", ctx.GetInt("id"), &model.LinkEvent{})
	if err != nil {
		replyError(ctx, err)
		return
	}
	replyList(ctx, events, cnt)
}

func linkEventClear(ctx *gin.Context) {
	err := database.History.Select(q.Eq("LinkID", ctx.GetInt("id"))).Delete(&model.LinkEvent{})
	if err != nil {
		replyError(ctx, err)
		return
	}

	replyOk(ctx, nil)
}

func linkEventClearAll(ctx *gin.Context) {
	err := database.History.Drop(&model.LinkEvent{})
	if err != nil {
		replyError(ctx, err)
		return
	}

	replyOk(ctx, nil)
}
