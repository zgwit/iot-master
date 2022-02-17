package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/connect"
	"github.com/zgwit/iot-master/database"
	"github.com/zgwit/iot-master/model"
)

func linkRoutes(app *gin.RouterGroup) {
	app.POST("list", linkList)
	app.POST("create", linkCreate)

	app.Use(parseParamId)
	app.POST(":id/update", linkUpdate)
	app.GET(":id/delete", linkDelete)
	app.GET(":id/close", linkClose)
	app.GET(":id/enable", linkEnable)
	app.GET(":id/disable", linkDisable)

}

func linkList(ctx *gin.Context) {
	var links []model.Link
	cnt, err := normalSearch(ctx, database.Master, &links)
	if err != nil {
		replyError(ctx, err)
		return
	}
	replyList(ctx, links, cnt)
}

func linkCreate(ctx *gin.Context) {
	var link model.Link
	err := ctx.ShouldBindJSON(&link)
	if err != nil {
		replyError(ctx, err)
		return
	}

	err = database.Master.Save(&link)
	if err != nil {
		replyError(ctx, err)
		return
	}

	//TODO 启动

	replyOk(ctx, link)
}

func linkUpdate(ctx *gin.Context) {
	var pid paramID
	err := ctx.ShouldBindUri(&pid)
	if err != nil {
		replyError(ctx, err)
		return
	}

	var link model.Link
	err = ctx.ShouldBindJSON(&link)
	if err != nil {
		replyError(ctx, err)
		return
	}
	link.ID = pid.ID

	err = database.Master.Update(&link)
	if err != nil {
		replyError(ctx, err)
		return
	}

	//TODO 重新启动

	replyOk(ctx, link)
}

func linkDelete(ctx *gin.Context) {
	var pid paramID
	err := ctx.ShouldBindUri(&pid)
	if err != nil {
		replyError(ctx, err)
		return
	}
	link := model.Link{ID: pid.ID}
	err = database.Master.DeleteStruct(&link)
	if err != nil {
		replyError(ctx, err)
		return
	}

	//TODO 重新启动

	replyOk(ctx, link)
}


func linkClose(ctx *gin.Context) {
	var lnk model.Link
	err := database.Master.One("ID", ctx.GetInt("id"), &lnk)
	if err != nil {
		replyError(ctx, err)
		return
	}

	tunnel := connect.GetTunnel(lnk.TunnelID)
	if tunnel == nil {
		replyFail(ctx, "tunnel not found")
		return
	}

	link := tunnel.GetLink(ctx.GetInt("id"))
	if link == nil {
		replyFail(ctx, "link not found")
		return
	}
	err = link.Close()
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
	//TODO 关闭
	replyOk(ctx, nil)
}
