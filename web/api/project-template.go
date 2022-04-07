package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/database"
	"github.com/zgwit/iot-master/model"
)

func projectTemplateRoutes(app *gin.RouterGroup) {
	app.POST("list", projectTemplateList)
	app.POST("create", projectTemplateCreate)

	app.Use(parseParamId)
	app.POST(":id/update", projectTemplateUpdate)
	app.GET(":id/delete", projectTemplateDelete)

}

func projectTemplateList(ctx *gin.Context) {
	var projectTemplates []model.ProjectTemplate
	cnt, err := normalSearch(ctx, database.Master, &projectTemplates)
	if err != nil {
		replyError(ctx, err)
		return
	}
	replyList(ctx, projectTemplates, cnt)
}

func projectTemplateCreate(ctx *gin.Context) {
	var projectTemplate model.ProjectTemplate
	err := ctx.ShouldBindJSON(&projectTemplate)
	if err != nil {
		replyError(ctx, err)
		return
	}

	err = database.Master.Save(&projectTemplate)
	if err != nil {
		replyError(ctx, err)
		return
	}

	replyOk(ctx, projectTemplate)
}

func projectTemplateUpdate(ctx *gin.Context) {
	var projectTemplate model.ProjectTemplate
	err := ctx.ShouldBindJSON(&projectTemplate)
	if err != nil {
		replyError(ctx, err)
		return
	}
	projectTemplate.ID = ctx.GetInt("id")

	err = database.Master.Update(&projectTemplate)
	if err != nil {
		replyError(ctx, err)
		return
	}

	replyOk(ctx, projectTemplate)
}

func projectTemplateDelete(ctx *gin.Context) {
	projectTemplate := model.ProjectTemplate{ID: ctx.GetInt("id")}
	err := database.Master.DeleteStruct(&projectTemplate)
	if err != nil {
		replyError(ctx, err)
		return
	}

	replyOk(ctx, projectTemplate)
}
