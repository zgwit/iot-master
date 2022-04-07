package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/database"
	"github.com/zgwit/iot-master/model"
)

func deviceTemplateRoutes(app *gin.RouterGroup) {
	app.POST("list", deviceTemplateList)
	app.POST("create", deviceTemplateCreate)

	app.Use(parseParamId)
	app.POST(":id/update", deviceTemplateUpdate)
	app.GET(":id/delete", deviceTemplateDelete)

}

func deviceTemplateList(ctx *gin.Context) {
	var deviceTemplates []model.DeviceTemplate
	cnt, err := normalSearch(ctx, database.Master, &deviceTemplates)
	if err != nil {
		replyError(ctx, err)
		return
	}
	replyList(ctx, deviceTemplates, cnt)
}

func deviceTemplateCreate(ctx *gin.Context) {
	var deviceTemplate model.DeviceTemplate
	err := ctx.ShouldBindJSON(&deviceTemplate)
	if err != nil {
		replyError(ctx, err)
		return
	}

	err = database.Master.Save(&deviceTemplate)
	if err != nil {
		replyError(ctx, err)
		return
	}

	replyOk(ctx, deviceTemplate)
}

func deviceTemplateUpdate(ctx *gin.Context) {
	var deviceTemplate model.DeviceTemplate
	err := ctx.ShouldBindJSON(&deviceTemplate)
	if err != nil {
		replyError(ctx, err)
		return
	}
	deviceTemplate.ID = ctx.GetInt("id")

	err = database.Master.Update(&deviceTemplate)
	if err != nil {
		replyError(ctx, err)
		return
	}

	replyOk(ctx, deviceTemplate)
}

func deviceTemplateDelete(ctx *gin.Context) {
	deviceTemplate := model.DeviceTemplate{ID: ctx.GetInt("id")}
	err := database.Master.DeleteStruct(&deviceTemplate)
	if err != nil {
		replyError(ctx, err)
		return
	}

	replyOk(ctx, deviceTemplate)
}
