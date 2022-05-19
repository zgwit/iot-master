package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/database"
	"github.com/zgwit/iot-master/model"
)

func hmiRoutes(app *gin.RouterGroup) {
	app.POST("list", hmiList)
	app.POST("create", hmiCreate)

	app.Use(parseParamStringId)

	app.GET(":id", hmiDetail)
	app.POST(":id", hmiUpdate)
	app.GET(":id/delete", hmiDelete)
}


func hmiList(ctx *gin.Context) {
	hs, cnt, err := normalSearch(ctx, database.Master, &model.HMI{})
	if err != nil {
		replyError(ctx, err)
		return
	}
	replyList(ctx, hs, cnt)
}

func hmiCreate(ctx *gin.Context) {
	var hmi model.HMI
	err := ctx.ShouldBindJSON(&hmi)
	if err != nil {
		replyError(ctx, err)
		return
	}

	err = database.Master.Save(&hmi)
	if err != nil {
		replyError(ctx, err)
		return
	}

	replyOk(ctx, hmi)
}

func hmiDetail(ctx *gin.Context) {
	var hmi model.HMI
	err := database.Master.One("Id", ctx.GetString("id"), &hmi)
	if err != nil {
		replyError(ctx, err)
		return
	}
	replyOk(ctx, hmi)
}

func hmiUpdate(ctx *gin.Context) {
	var hmi model.HMI
	err := ctx.ShouldBindJSON(&hmi)
	if err != nil {
		replyError(ctx, err)
		return
	}
	hmi.Id = ctx.GetString("id")

	err = database.Master.Update(&hmi)
	if err != nil {
		replyError(ctx, err)
		return
	}

	replyOk(ctx, hmi)
}

func hmiDelete(ctx *gin.Context) {
	hmi := model.HMI{Id: ctx.GetString("id")}
	err := database.Master.DeleteStruct(&hmi)
	if err != nil {
		replyError(ctx, err)
		return
	}

	replyOk(ctx, hmi)
}
