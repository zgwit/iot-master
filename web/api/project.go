package api

import (
	"github.com/gin-gonic/gin"
	"iot-master/internal/core"
	"iot-master/model"
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
	return core.RemoveProject(id.(uint64))
}

func projectStart(ctx *gin.Context) {
	project := core.GetProject(ctx.GetUint64("id"))
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
	project := core.GetProject(ctx.GetUint64("id"))
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
	_ = core.RemoveProject(id.(uint64))
	_, err := core.LoadProject(id.(uint64))
	return err
}

func afterProjectDisable(id interface{}) error {
	return core.RemoveProject(id.(uint64))
}

func projectContext(ctx *gin.Context) {
	project := core.GetProject(ctx.GetUint64("id"))
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

	project := core.GetProject(ctx.GetUint64("id"))
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
