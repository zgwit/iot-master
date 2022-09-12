package api

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/net/websocket"
	"iot-master/db"
	"iot-master/master"
	"iot-master/model"
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
		if has && err == nil {
			project.Template = template.Name
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

func afterProjectDelete(id interface{}) error {
	return master.RemoveProject(id.(int64))
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

func afterProjectEnable(id interface{}) error {
	_ = master.RemoveProject(id.(int64))
	_, err := master.LoadProject(id.(int64))
	return err
}

func afterProjectDisable(id interface{}) error {
	return master.RemoveProject(id.(int64))
}

func projectContext(ctx *gin.Context) {
	project := master.GetProject(ctx.GetInt64("id"))
	if project == nil {
		replyFail(ctx, "找不到项目")
		return
	}
	replyOk(ctx, project.Context)
}

func projectContextUpdate(ctx *gin.Context) {
	var values map[string]interface{}
	err := ctx.ShouldBindJSON(values)
	if err != nil {
		replyError(ctx, err)
		return
	}

	project := master.GetProject(ctx.GetInt64("id"))
	if project == nil {
		replyFail(ctx, "找不到项目")
		return
	}

	for k, v := range values {
		err := project.Set(k, v)
		if err != nil {
			replyError(ctx, err)
			return
		}
	}

	replyOk(ctx, nil)
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

func projectTargets(ctx *gin.Context) {
	var project model.Project
	has, err := db.Engine.ID(ctx.GetInt64("id")).Cols("devices").Get(&project)
	if err != nil {
		replyError(ctx, err)
		return
	} else if !has {
		replyFail(ctx, "记录不存在")
		return
	}

	set := make(map[string]bool)
	for _, d := range project.Devices {
		set[d.Name] = true

		var dev model.Device
		has, err = db.Engine.ID(d.Id).Cols("product_id", "tags").Get(&dev)
		if has && err == nil {
			//查询产品
			if dev.ProductId != "" {
				var p model.Product
				has, err = db.Engine.ID(dev.ProductId).Cols("tags").Get(&p)
				if has && err == nil {
					dev.Tags = p.Tags
				}
			}
			for _, t := range dev.Tags {
				set[t] = true
			}
		}
	}

	targets := make([]string, 0)
	for k, _ := range set {
		targets = append(targets, k)
	}

	replyOk(ctx, targets)
}

func templateTargets(ctx *gin.Context) {
	var template model.Template
	has, err := db.Engine.ID(ctx.GetInt64("id")).Cols("products").Get(&template)
	if err != nil {
		replyError(ctx, err)
		return
	} else if !has {
		replyFail(ctx, "记录不存在")
		return
	}

	set := make(map[string]bool)
	for _, d := range template.Products {
		set[d.Name] = true

		var dev model.Device
		has, err = db.Engine.ID(d.Id).Cols("tags").Get(&dev)
		if has && err == nil {
			for _, t := range dev.Tags {
				set[t] = true
			}
		}
	}

	targets := make([]string, 0)
	for k, _ := range set {
		targets = append(targets, k)
	}

	replyOk(ctx, targets)
}
