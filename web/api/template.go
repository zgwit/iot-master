package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/database"
	"github.com/zgwit/iot-master/model"
)

func templateRoutes(app *gin.RouterGroup) {
	app.POST("list", templateList)
	app.POST("create", templateCreate)

	app.Use(parseParamId)
	app.POST(":id/update", templateUpdate)
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

	err = database.Master.Save(&template)
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
	template.ID = ctx.GetInt("id")

	err = database.Master.Update(&template)
	if err != nil {
		replyError(ctx, err)
		return
	}

	replyOk(ctx, template)
}

func templateDelete(ctx *gin.Context) {
	template := model.Template{ID: ctx.GetInt("id")}
	err := database.Master.DeleteStruct(&template)
	if err != nil {
		replyError(ctx, err)
		return
	}

	replyOk(ctx, template)
}
