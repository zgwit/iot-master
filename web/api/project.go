package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/db"
	"github.com/zgwit/iot-master/master"
	"github.com/zgwit/iot-master/model"
	"golang.org/x/net/websocket"
)

func projectList(ctx *gin.Context) {
	//补充信息
	var projects []model.ProjectEx

	var body paramSearchEx
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		replyError(ctx, err)
		return
	}

	query := body.toQuery()

	query.Table("project")
	query.Select("project.*, " + //TODO 只返回需要的字段
		" 0 as running, template.name as template")
	query.Join("LEFT", "template", "project.template_id=template.id")

	cnt, err := query.FindAndCount(&projects)
	if err != nil {
		replyError(ctx, err)
		return
	}

	//补充running状态
	for _, dev := range projects {
		d := master.GetProject(dev.Id)
		if d != nil {
			dev.Running = d.Running()
		}
	}

	replyList(ctx, projects, cnt)
}

func afterProjectCreate(data interface{}) error {
	project := data.(*model.Project)
	_, err := master.LoadProject(project.Id)
	return err
}

func projectDetail(ctx *gin.Context) {
	var project model.ProjectEx

	has, err := db.Engine.ID(ctx.GetInt64("id")).Get(&project.Project)
	if err != nil {
		replyError(ctx, err)
		return
	}
	if !has {
		replyFail(ctx, "项目不存在")
		return
	}

	if project.TemplateId != "" {
		var template model.Template
		has, err := db.Engine.ID(project.Template).Get(&template)
		if has && err != nil {
			project.ProjectContent = template.ProjectContent
		}
	}

	d := master.GetProject(project.Id)
	if d != nil {
		project.Running = d.Running()
	}

	replyOk(ctx, project)
}

func afterProjectUpdate(data interface{}) error {
	project := data.(*model.Project)
	_ = master.RemoveProject(project.Id)
	_, err := master.LoadProject(project.Id)
	return err
}

func afterProjectDelete(data interface{}) error {
	project := data.(*model.Project)
	return master.RemoveProject(project.Id)
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

func afterProjectEnable(data interface{}) error {
	project := data.(*model.Project)
	_ = master.RemoveProject(project.Id)
	_, err := master.LoadProject(project.Id)
	return err
}

func afterProjectDisable(data interface{}) error {
	project := data.(*model.Project)
	return master.RemoveProject(project.Id)
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
