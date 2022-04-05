package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/database"
	"github.com/zgwit/iot-master/master"
	"github.com/zgwit/iot-master/model"
)

func tunnelRoutes(app *gin.RouterGroup) {
	app.POST("list", tunnelList)
	app.POST("create", tunnelCreate)

	app.Use(parseParamId)
	app.POST(":id/update", tunnelUpdate)
	app.GET(":id/delete", tunnelDelete)
	app.GET(":id/start", tunnelStart)
	app.GET(":id/stop", tunnelStop)
	app.GET(":id/enable", tunnelEnable)
	app.GET(":id/disable", tunnelDisable)

}

func tunnelList(ctx *gin.Context) {
	var tunnels []model.Tunnel
	cnt, err := normalSearch(ctx, database.Master, &tunnels)
	if err != nil {
		replyError(ctx, err)
		return
	}
	replyList(ctx, tunnels, cnt)
}

func tunnelCreate(ctx *gin.Context) {
	var tunnel model.Tunnel
	err := ctx.ShouldBindJSON(&tunnel)
	if err != nil {
		replyError(ctx, err)
		return
	}

	err = database.Master.Save(&tunnel)
	if err != nil {
		replyError(ctx, err)
		return
	}

	//TODO 启动

	replyOk(ctx, tunnel)
}

func tunnelUpdate(ctx *gin.Context) {
	var tunnel model.Tunnel
	err := ctx.ShouldBindJSON(&tunnel)
	if err != nil {
		replyError(ctx, err)
		return
	}
	tunnel.ID = ctx.GetInt("id")

	err = database.Master.Update(&tunnel)
	if err != nil {
		replyError(ctx, err)
		return
	}

	//TODO 重新启动

	replyOk(ctx, tunnel)
}

func tunnelDelete(ctx *gin.Context) {
	tunnel := model.Tunnel{ID: ctx.GetInt("id")}
	err := database.Master.DeleteStruct(&tunnel)
	if err != nil {
		replyError(ctx, err)
		return
	}

	//TODO 关闭

	replyOk(ctx, tunnel)
}

func tunnelStart(ctx *gin.Context) {
	tunnel := master.GetTunnel(ctx.GetInt("id"))
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
	tunnel := master.GetTunnel(ctx.GetInt("id"))
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


func tunnelEnable(ctx *gin.Context) {
	err := database.Master.UpdateField(&model.Tunnel{ID: ctx.GetInt("id")}, "Disabled", false)
	if err != nil {
		replyError(ctx, err)
		return
	}
	//TODO 启动
	replyOk(ctx, nil)
}

func tunnelDisable(ctx *gin.Context) {
	err := database.Master.UpdateField(&model.Tunnel{ID: ctx.GetInt("id")}, "Disabled", true)
	if err != nil {
		replyError(ctx, err)
		return
	}
	//TODO 关闭
	replyOk(ctx, nil)
}