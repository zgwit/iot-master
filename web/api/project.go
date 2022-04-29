package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/database"
	"github.com/zgwit/iot-master/log"
	"github.com/zgwit/iot-master/master"
	"github.com/zgwit/iot-master/model"
	"github.com/zgwit/storm/v3/q"
	"golang.org/x/net/websocket"
)

func projectRoutes(app *gin.RouterGroup) {
	app.POST("list", projectList)
	app.POST("create", projectCreate)

	app.GET("event/clear", projectEventClearAll)
	app.GET("alarm/clear", projectAlarmClearAll)

	app.Use(parseParamId)
	app.GET(":id", projectDetail)
	app.POST(":id", projectUpdate)
	app.GET(":id/delete", projectDelete)
	app.GET(":id/start", projectStart)
	app.GET(":id/stop", projectStop)
	app.GET(":id/enable", projectEnable)
	app.GET(":id/disable", projectDisable)
	app.GET(":id/context", projectContext)
	app.GET(":id/watch", projectWatch)
	app.POST(":id/event/list", projectEvent)
	app.GET(":id/event/clear", projectEventClear)
	app.POST(":id/alarm/list", projectAlarm)
	app.GET(":id/alarm/clear", projectAlarmClear)

}

func projectList(ctx *gin.Context) {
	projects, cnt, err := normalSearch(ctx, database.Master, &model.Project{})
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

func projectDetail(ctx *gin.Context) {
	var project model.Project
	err := database.Master.One("Id", ctx.GetInt("id"), &project)
	if err != nil {
		replyError(ctx, err)
		return
	}
	replyOk(ctx, project)
}

func projectUpdate(ctx *gin.Context) {
	var project model.Project
	err := ctx.ShouldBindJSON(&project)
	if err != nil {
		replyError(ctx, err)
		return
	}
	project.Id = ctx.GetInt("id")

	err = database.Master.Update(&project)
	if err != nil {
		replyError(ctx, err)
		return
	}

	replyOk(ctx, project)


	//重新启动
	go func() {
		prj := master.GetProject(ctx.GetInt("id"))
		if prj == nil {
			return
		}
		err = prj.Stop()
		if err != nil {
			log.Error(err)
			return
		}
		err = prj.Start()
		if err != nil {
			log.Error(err)
			return
		}
	}()
}

func projectDelete(ctx *gin.Context) {
	project := model.Project{Id: ctx.GetInt("id")}
	err := database.Master.DeleteStruct(&project)
	if err != nil {
		replyError(ctx, err)
		return
	}



	replyOk(ctx, project)

	//关闭
	go func() {
		project := master.GetProject(ctx.GetInt("id"))
		if project == nil {
			return
		}
		err := project.Stop()
		if err != nil {
			log.Error(err)
			return
		}
	}()
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
	err := database.Master.UpdateField(&model.Project{Id: ctx.GetInt("id")}, "Disabled", false)
	if err != nil {
		replyError(ctx, err)
		return
	}
	replyOk(ctx, nil)

	//启动
	go func() {
		project := master.GetProject(ctx.GetInt("id"))
		if project == nil {
			return
		}
		err := project.Start()
		if err != nil {
			log.Error(err)
			return
		}
	}()
}

func projectDisable(ctx *gin.Context) {
	err := database.Master.UpdateField(&model.Project{Id: ctx.GetInt("id")}, "Disabled", true)
	if err != nil {
		replyError(ctx, err)
		return
	}
	replyOk(ctx, nil)

	//关闭
	go func() {
		project := master.GetProject(ctx.GetInt("id"))
		if project == nil {
			return
		}
		err := project.Stop()
		if err != nil {
			log.Error(err)
			return
		}
	}()
}

func projectContext(ctx *gin.Context) {
	project := master.GetProject(ctx.GetInt("id"))
	if project == nil {
		replyFail(ctx, "找不到项目")
		return
	}
	replyOk(ctx, project.Context)
}

func projectWatch(ctx *gin.Context) {
	project := master.GetProject(ctx.GetInt("id"))
	if project == nil {
		replyFail(ctx, "找不到项目")
		return
	}
	websocket.Handler(func(ws *websocket.Conn) {
		watchAllEvents(ws, project)
	}).ServeHTTP(ctx.Writer, ctx.Request)
}

func projectEvent(ctx *gin.Context) {
	events, cnt, err := normalSearchById(ctx, database.History, "ProjectId", ctx.GetInt("id"), &model.ProjectEvent{})
	if err != nil {
		replyError(ctx, err)
		return
	}
	replyList(ctx, events, cnt)
}

func projectEventClear(ctx *gin.Context) {
	err := database.History.Select(q.Eq("ProjectId", ctx.GetInt("id"))).Delete(&model.ProjectEvent{})
	if err != nil {
		replyError(ctx, err)
		return
	}

	replyOk(ctx, nil)
}

func projectEventClearAll(ctx *gin.Context) {
	err := database.History.Drop(&model.ProjectEvent{})
	if err != nil {
		replyError(ctx, err)
		return
	}

	replyOk(ctx, nil)
}

func projectAlarm(ctx *gin.Context) {
	alarms, cnt, err := normalSearchById(ctx, database.History, "ProjectId", ctx.GetInt("id"), &model.ProjectAlarm{})
	if err != nil {
		replyError(ctx, err)
		return
	}
	replyList(ctx, alarms, cnt)
}

func projectAlarmClear(ctx *gin.Context) {
	err := database.History.Select(q.Eq("ProjectId", ctx.GetInt("id"))).Delete(&model.ProjectAlarm{})
	if err != nil {
		replyError(ctx, err)
		return
	}

	replyOk(ctx, nil)
}


func projectAlarmClearAll(ctx *gin.Context) {
	err := database.History.Drop(&model.ProjectAlarm{})
	if err != nil {
		replyError(ctx, err)
		return
	}

	replyOk(ctx, nil)
}
