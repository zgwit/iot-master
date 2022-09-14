package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/internal/core"
	"github.com/zgwit/iot-master/model"
)

func afterProjectCreate(data interface{}) error {
	project := data.(*model.Project)
	core.Projects.Store(project.Id, core.NewProject(project.Id))
	return nil
}

func afterProjectUpdate(data interface{}) error {
	//project := data.(*model.Project)

	return nil
}

func afterProjectDelete(id interface{}) error {
	core.Projects.Delete(id.(string))
	return nil
}

func projectValues(ctx *gin.Context) {
	project := core.Projects.Load(ctx.GetString("id"))
	if project == nil {
		replyFail(ctx, "找不到项目")
		return
	}
	replyOk(ctx, project.Values)
}

func projectAssign(ctx *gin.Context) {
	var values map[string]interface{}
	err := ctx.ShouldBindJSON(values)
	if err != nil {
		replyError(ctx, err)
		return
	}

	project := core.Projects.Load(ctx.GetString("id"))
	if project == nil {
		replyFail(ctx, "找不到项目")
		return
	}

	err = project.Assign(values)
	if err != nil {
		replyError(ctx, err)
		return
	}

	replyOk(ctx, nil)
}
