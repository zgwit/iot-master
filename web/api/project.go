package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/database"
	"github.com/zgwit/iot-master/master"
	"github.com/zgwit/iot-master/model"
	"golang.org/x/net/websocket"
)

func projectRoutes(app *gin.RouterGroup) {
	app.POST("list", projectList)
	app.POST("create", projectCreate)

	app.Use(parseParamId)
	app.POST(":id/update", projectUpdate)
	app.GET(":id/delete", projectDelete)
	app.GET(":id/start", projectStart)
	app.GET(":id/stop", projectStop)
	app.GET(":id/enable", projectEnable)
	app.GET(":id/disable", projectDisable)
	app.GET(":id/watch", projectWatch)

}

func projectList(ctx *gin.Context) {
	var projects []model.Project
	cnt, err := normalSearch(ctx, database.Master, &projects)
	if err != nil {
		replyError(ctx, err)
		return
	}
	replyList(ctx, projects, cnt)
}

func projectCreate(ctx *gin.Context) {
	var project model.Project
	err := ctx.ShouldBindJSON(&project)
	if err != nil {
		replyError(ctx, err)
		return
	}

	err = database.Master.Save(&project)
	if err != nil {
		replyError(ctx, err)
		return
	}

	//TODO 启动

	replyOk(ctx, project)
}

func projectUpdate(ctx *gin.Context) {
	var project model.Project
	err := ctx.ShouldBindJSON(&project)
	if err != nil {
		replyError(ctx, err)
		return
	}
	project.ID = ctx.GetInt("id")

	err = database.Master.Update(&project)
	if err != nil {
		replyError(ctx, err)
		return
	}

	//TODO 重新启动

	replyOk(ctx, project)
}

func projectDelete(ctx *gin.Context) {
	project := model.Project{ID: ctx.GetInt("id")}
	err := database.Master.DeleteStruct(&project)
	if err != nil {
		replyError(ctx, err)
		return
	}

	//TODO 关闭

	replyOk(ctx, project)
}

func projectStart(ctx *gin.Context) {
	project := master.GetProject(ctx.GetInt("id"))
	if project == nil {
		replyFail(ctx, "not found")
		return
	}
	err := project.Start()
	if err != nil {
		replyError(ctx, err)
		return
	}

	replyOk(ctx, nil)
}

func projectStop(ctx *gin.Context) {
	project := master.GetProject(ctx.GetInt("id"))
	if project == nil {
		replyFail(ctx, "not found")
		return
	}
	err := project.Stop()
	if err != nil {
		replyError(ctx, err)
		return
	}

	replyOk(ctx, nil)
}


func projectEnable(ctx *gin.Context) {
	err := database.Master.UpdateField(&model.Project{ID: ctx.GetInt("id")}, "Disabled", false)
	if err != nil {
		replyError(ctx, err)
		return
	}
	//TODO 启动
	replyOk(ctx, nil)
}

func projectDisable(ctx *gin.Context) {
	err := database.Master.UpdateField(&model.Project{ID: ctx.GetInt("id")}, "Disabled", true)
	if err != nil {
		replyError(ctx, err)
		return
	}
	//TODO 关闭
	replyOk(ctx, nil)
}

func projectWatch(ctx *gin.Context) {
	project := master.GetProject(ctx.GetInt("id"))
	if project == nil {
		replyFail(ctx, "找不到链接")
		return
	}
	websocket.Handler(func(ws *websocket.Conn) {
		watchAllEvents(ws, project)
	}).ServeHTTP(ctx.Writer, ctx.Request)
}