package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/master"
	"github.com/zgwit/iot-master/model"
	"golang.org/x/net/websocket"
)

func projectList(ctx *gin.Context) {
	//补充信息
	projects := make([]*model.ProjectEx, 0)

	var body paramSearchEx
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		replyError(ctx, err)
		return
	}

	query := body.toQuery()
	query.Join("LEFT", "template", "project.template_id=template.id")

	cnt, err := query.FindAndCount(projects)
	if err != nil {
		replyError(ctx, err)
		return
	}

	//prj.Running = d.Running()
	//prj.Template = template.Name
	//prj.ProjectContent = template.ProjectContent

	replyList(ctx, projects, cnt)
}

func afterProjectCreate(data interface{}) error {
	project := data.(*model.Project)
	_, err := master.LoadProject(project.Id)
	return err
}

func projectDetail(ctx *gin.Context) {
	prj := model.ProjectEx{}
	var template model.Template
	prj.Template = template.Name
	prj.ProjectContent = template.ProjectContent

	replyOk(ctx, prj)
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
