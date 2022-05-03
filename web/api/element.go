package api

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/zgwit/iot-master/database"
	"github.com/zgwit/iot-master/model"
)

func elementRoutes(app *gin.RouterGroup) {
	app.POST("list", elementList)
	app.POST("create", elementCreate)

	app.Use(parseParamStringId)
	app.GET(":id", elementDetail)
	app.POST(":id", elementUpdate)
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

	//使用UUId作为Id
	if element.Id == "" {
		element.Id = uuid.NewString()
	}
	//保存
	err = database.Master.Save(&element)
	if err != nil {
		replyError(ctx, err)
		return
	}

	replyOk(ctx, element)
}

func elementDetail(ctx *gin.Context) {
	var element model.Element
	err := database.Master.One("Id", ctx.GetString("id"), &element)
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
	element.Id = ctx.GetString("id")

	err = database.Master.Update(&element)
	if err != nil {
		replyError(ctx, err)
		return
	}

	replyOk(ctx, element)
}

func elementDelete(ctx *gin.Context) {
	element := model.Element{Id: ctx.GetString("id")}
	err := database.Master.DeleteStruct(&element)
	if err != nil {
		replyError(ctx, err)
		return
	}

	replyOk(ctx, element)
}
