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
	app.POST(":id/alarm/list", projectAlarm)
	app.GET(":id/alarm/clear", projectAlarmClear)

}

func projectList(ctx *gin.Context) {
	records, cnt, err := normalSearch(ctx, database.Master, &model.Project{})
	if err != nil {
		replyError(ctx, err)
		return
	}

	//补充信息
	projects := records.(*[]*model.Project)
	ps := make([]*model.ProjectEx, 0) //len(projects)

	for _, d := range *projects {
		prj := &model.ProjectEx{Project: *d}
		ps = append(ps, prj)
		d := master.GetProject(prj.Id)
		if d != nil {
			prj.Running = d.Running()
		}
		if prj.TemplateId != "" {
			var template model.Template
			err := database.Master.One("Id", prj.TemplateId, &template)
			if err == nil {
				prj.Template = template.Name
				prj.ProjectContent = template.ProjectContent
			} // else err
		}
	}

	replyList(ctx, ps, cnt)
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
	err := database.Master.One("Id", ctx.GetInt64("id"), &project)
	if err != nil {
		replyError(ctx, err)
		return
	}

	//补充信息
	prj := model.ProjectEx{Project: project}
	d := master.GetProject(prj.Id)
	if d != nil {
		prj.Running = d.Running()
	}
	if prj.TemplateId != "" {
		var template model.Template
		err := database.Master.One("Id", prj.TemplateId, &template)
		if err == nil {
			prj.Template = template.Name
			prj.ProjectContent = template.ProjectContent
		}
	}

	replyOk(ctx, prj)
}

func projectUpdate(ctx *gin.Context) {
	var project model.Project
	err := ctx.ShouldBindJSON(&project)
	if err != nil {
		replyError(ctx, err)
		return
	}
	project.Id = ctx.GetInt64("id")

	err = database.Master.Update(&project)
	if err != nil {
		replyError(ctx, err)
		return
	}

	replyOk(ctx, project)


	//重新启动
	go func() {
		_ = master.RemoveProject(project.Id)
		_, err := master.LoadProject(project.Id)
		if err != nil {
			log.Error(err)
			return
		}
	}()
}

func projectDelete(ctx *gin.Context) {
	project := model.Project{Id: ctx.GetInt64("id")}
	err := database.Master.DeleteStruct(&project)
	if err != nil {
		replyError(ctx, err)
		return
	}

	replyOk(ctx, project)

	//关闭
	go func() {
		err := master.RemoveProject(project.Id)
		if err != nil {
			log.Error(err)
		}
	}()
}

func projectStart(ctx *gin.Context) {
	project := master.GetProject(ctx.GetInt64("id"))
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
	project := master.GetProject(ctx.GetInt64("id"))
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
	err := database.Master.UpdateField(&model.Project{Id: ctx.GetInt64("id")}, "Disabled", false)
	if err != nil {
		replyError(ctx, err)
		return
	}
	replyOk(ctx, nil)

	//启动
	go func() {
		_, err := master.LoadProject(ctx.GetInt64("id"))
		if err != nil {
			log.Error(err)
			return
		}
	}()
}

func projectDisable(ctx *gin.Context) {
	err := database.Master.UpdateField(&model.Project{Id: ctx.GetInt64("id")}, "Disabled", true)
	if err != nil {
		replyError(ctx, err)
		return
	}
	replyOk(ctx, nil)

	//关闭
	go func() {
		project := master.GetProject(ctx.GetInt64("id"))
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
	project := master.GetProject(ctx.GetInt64("id"))
	if project == nil {
		replyFail(ctx, "找不到项目")
		return
	}
	replyOk(ctx, project.Context)
}

func projectWatch(ctx *gin.Context) {
	project := master.GetProject(ctx.GetInt64("id"))
	if project == nil {
		replyFail(ctx, "找不到项目")
		return
	}
	websocket.Handler(func(ws *websocket.Conn) {
		watchAllEvents(ws, project)
	}).ServeHTTP(ctx.Writer, ctx.Request)
}

func projectAlarm(ctx *gin.Context) {
	alarms, cnt, err := normalSearchById(ctx, database.History, "ProjectId", ctx.GetInt64("id"), &model.ProjectAlarm{})
	if err != nil {
		replyError(ctx, err)
		return
	}
	replyList(ctx, alarms, cnt)
}

func projectAlarmClear(ctx *gin.Context) {
	err := database.History.Select(q.Eq("ProjectId", ctx.GetInt64("id"))).Delete(&model.ProjectAlarm{})
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
