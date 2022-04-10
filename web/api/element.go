package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/database"
	"github.com/zgwit/iot-master/model"
)

func elementRoutes(app *gin.RouterGroup) {
	app.POST("list", elementList)
	app.POST("create", elementCreate)

	app.Use(parseParamId)
	app.POST(":id/update", elementUpdate)
	app.GET(":id/delete", elementDelete)

}

func elementList(ctx *gin.Context) {
	elements, cnt, err := normalSearch(ctx, database.Master, &model.Element{})
	if err != nil {
		replyError(ctx, err)
		return
	}
	replyList(ctx, elements, cnt)
}

func elementCreate(ctx *gin.Context) {
	var element model.Element
	err := ctx.ShouldBindJSON(&element)
	if err != nil {
		replyError(ctx, err)
		return
	}

	err = database.Master.Save(&element)
	if err != nil {
		replyError(ctx, err)
		return
	}

	replyOk(ctx, element)
}

func elementUpdate(ctx *gin.Context) {
	var element model.Element
	err := ctx.ShouldBindJSON(&element)
	if err != nil {
		replyError(ctx, err)
		return
	}
	element.ID = ctx.GetInt("id")

	err = database.Master.Update(&element)
	if err != nil {
		replyError(ctx, err)
		return
	}

	replyOk(ctx, element)
}

func elementDelete(ctx *gin.Context) {
	element := model.Element{ID: ctx.GetInt("id")}
	err := database.Master.DeleteStruct(&element)
	if err != nil {
		replyError(ctx, err)
		return
	}

	replyOk(ctx, element)
}
