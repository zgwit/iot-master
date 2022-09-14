package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/model"
)

func afterProjectCreate(data interface{}) error {
	project := data.(*model.Project)
	_, err := core.LoadProject(project.Id)
	return err
}

func afterProjectUpdate(data interface{}) error {
	project := data.(*model.Project)
	_ = core.RemoveProject(project.Id)
	_, err := core.LoadProject(project.Id)
	return err
}

func afterProjectDelete(id interface{}) error {
	return core.RemoveProject(id.(int64))
}

func afterProjectEnable(id interface{}) error {
	_ = core.RemoveProject(id.(int64))
	_, err := core.LoadProject(id.(int64))
	return err
}

func afterProjectDisable(id interface{}) error {
	return core.RemoveProject(id.(int64))
}

func projectContext(ctx *gin.Context) {
	project := core.GetProject(ctx.GetInt64("id"))
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

	project := core.GetProject(ctx.GetInt64("id"))
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
