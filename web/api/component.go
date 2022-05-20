package api

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/zgwit/iot-master/database"
	"github.com/zgwit/iot-master/model"
)

func  componentRoutes(app *gin.RouterGroup) {
	app.POST("list",  componentList)
	app.POST("create",  componentCreate)

	app.Use(parseParamStringId)

	app.GET(":id",  componentDetail)
	app.POST(":id",  componentUpdate)
	app.GET(":id/delete",  componentDelete)
}


func  componentList(ctx *gin.Context) {
	hs, cnt, err := normalSearch(ctx, database.Master, &model.Component{})
	if err != nil {
		replyError(ctx, err)
		return
	}
	replyList(ctx, hs, cnt)
}

func  componentCreate(ctx *gin.Context) {
	var  component model.Component
	err := ctx.ShouldBindJSON(& component)
	if err != nil {
		replyError(ctx, err)
		return
	}

	//使用UUId作为Id
	if  component.Id == "" {
		 component.Id = uuid.NewString()
	}
	//保存
	err = database.Master.Save(& component)
	if err != nil {
		replyError(ctx, err)
		return
	}

	replyOk(ctx,  component)
}

func  componentDetail(ctx *gin.Context) {
	var  component model.Component
	err := database.Master.One("Id", ctx.GetString("id"), & component)
	if err != nil {
		replyError(ctx, err)
		return
	}
	replyOk(ctx,  component)
}

func  componentUpdate(ctx *gin.Context) {
	var  component model.Component
	err := ctx.ShouldBindJSON(& component)
	if err != nil {
		replyError(ctx, err)
		return
	}
	 component.Id = ctx.GetString("id")

	err = database.Master.Update(& component)
	if err != nil {
		replyError(ctx, err)
		return
	}

	replyOk(ctx,  component)
}

func  componentDelete(ctx *gin.Context) {
	 component := model.Component{Id: ctx.GetString("id")}
	err := database.Master.DeleteStruct(& component)
	if err != nil {
		replyError(ctx, err)
		return
	}

	replyOk(ctx,  component)
}
