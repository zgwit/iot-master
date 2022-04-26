package api

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/zgwit/iot-master/database"
	"github.com/zgwit/iot-master/model"
)

func templateRoutes(app *gin.RouterGroup) {
	app.POST("list", templateList)
	app.POST("create", templateCreate)

	app.Use(parseParamStringId)
	app.GET(":id", templateDetail)
	app.POST(":id", templateUpdate)
	app.GET(":id/delete", templateDelete)

}

func templateList(ctx *gin.Context) {
	templates, cnt, err := normalSearch(ctx, database.Master, &model.Template{})
	if err != nil {
		replyError(ctx, err)
		return
	}
	replyList(ctx, templates, cnt)
}

func templateCreate(ctx *gin.Context) {
	var template model.Template
	err := ctx.ShouldBindJSON(&template)
	if err != nil {
		replyError(ctx, err)
		return
	}

	//使用UUID作为ID
	template.ID = uuid.NewString()
	//保存
	err = database.Master.Save(&template)
	if err != nil {
		replyError(ctx, err)
		return
	}

	replyOk(ctx, template)
}

func templateDetail(ctx *gin.Context) {
	var template model.Template
	err := database.Master.One("ID", ctx.GetString("id"), &template)
	if err != nil {
		replyError(ctx, err)
		return
	}
	replyOk(ctx, template)
}

func templateUpdate(ctx *gin.Context) {
	var template model.Template
	err := ctx.ShouldBindJSON(&template)
	if err != nil {
		replyError(ctx, err)
		return
	}
	template.ID = ctx.GetString("id")

	err = database.Master.Update(&template)
	if err != nil {
		replyError(ctx, err)
		return
	}

	replyOk(ctx, template)
}

func templateDelete(ctx *gin.Context) {
	template := model.Template{ID: ctx.GetString("id")}
	err := database.Master.DeleteStruct(&template)
	if err != nil {
		replyError(ctx, err)
		return
	}

	replyOk(ctx, template)
}
