package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/v4/project"
	"github.com/zgwit/iot-master/v4/types"
	"github.com/zgwit/iot-master/v4/web/attach"
	"github.com/zgwit/iot-master/v4/web/curd"
	"github.com/zgwit/iot-master/v4/web/export"
)

// @Summary 查询项目
// @Schemes
// @Description 查询项目
// @Tags project
// @Param search body curd.ParamSearch true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyList[types.Project] 返回项目信息
// @Router /project/search [post]
func noopProjectSearch() {}

// @Summary 查询项目
// @Schemes
// @Description 查询项目
// @Tags project
// @Param search query curd.ParamList true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyList[types.Project] 返回项目信息
// @Router /project/list [get]
func noopProjectList() {}

// @Summary 创建项目
// @Schemes
// @Description 创建项目
// @Tags project
// @Param search body types.Project true "项目信息"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[types.Project] 返回项目信息
// @Router /project/create [post]
func noopProjectCreate() {}

// @Summary 修改项目
// @Schemes
// @Description 修改项目
// @Tags project
// @Param id path int true "项目ID"
// @Param project body types.Project true "项目信息"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[types.Project] 返回项目信息
// @Router /project/{id} [post]
func noopProjectUpdate() {}

// @Summary 删除项目
// @Schemes
// @Description 删除项目
// @Tags project
// @Param id path int true "项目ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[types.Project] 返回项目信息
// @Router /project/{id}/delete [get]
func noopProjectDelete() {}

// @Summary 导出项目
// @Schemes
// @Description 导出项目
// @Tags project
// @Accept json
// @Produce octet-stream
// @Router /project/export [get]
func noopProjectExport() {}

// @Summary 导入项目
// @Schemes
// @Description 导入项目
// @Tags project
// @Param file formData file true "压缩包"
// @Accept mpfd
// @Produce json
// @Success 200 {object} curd.ReplyData[int64] 返回项目数量
// @Router /project/import [post]
func noopProjectImport() {}

// @Summary 启用项目
// @Schemes
// @Description 启用项目
// @Tags project
// @Param id path int true "项目ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[types.Project] 返回项目信息
// @Router /project/{id}/enable [get]
func noopProjectEnable() {}

// @Summary 禁用项目
// @Schemes
// @Description 禁用项目
// @Tags project
// @Param id path int true "项目ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[types.Project] 返回项目信息
// @Router /project/{id}/disable [get]
func noopProjectDisable() {}

// @Summary 启动项目
// @Schemes
// @Description 启动项目
// @Tags project
// @Param id path int true "项目ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[types.Project] 返回项目信息
// @Router /project/{id}/start [get]
func noopProjectStart() {}

// @Summary 停止项目
// @Schemes
// @Description 停止项目
// @Tags project
// @Param id path int true "项目ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[types.Project] 返回项目信息
// @Router /project/{id}/stop [get]
func noopProjectStop() {}

func projectRouter(app *gin.RouterGroup) {

	app.POST("/count", curd.ApiCount[types.Project]())

	app.POST("/search", curd.ApiSearchHook[types.Project](func(datum []*types.Project) error {
		for _, v := range datum {
			p := project.Get(v.Id)
			if p != nil {
				//v.Running = p.Running
			}
		}
		return nil
	}))

	app.GET("/list", curd.ApiListHook[types.Project](func(datum []*types.Project) error {
		for _, v := range datum {
			p := project.Get(v.Id)
			if p != nil {
				//v.Running = p.Running
			}
		}
		return nil
	}))
	app.POST("/create", curd.ApiCreateHook[types.Project](curd.GenerateKSUID[types.Project](), nil))

	app.GET("/:id", curd.ParseParamStringId, curd.ApiGetHook[types.Project](func(m *types.Project) error {
		p := project.Get(m.Id)
		if p != nil {
			//m.Running = p.Running
		}
		return nil
	}))

	app.POST("/:id", curd.ParseParamStringId, curd.ApiUpdateHook[types.Project](nil, nil,
		"id", "name", "icon", "description", "keywords", "disabled"))

	app.GET("/:id/delete", curd.ParseParamStringId, curd.ApiDeleteHook[types.Project](nil, nil))

	app.GET("/export", export.ApiExport("project", "项目"))

	app.POST("/import", export.ApiImport("project"))

	app.GET(":id/disable", curd.ParseParamStringId, curd.ApiDisableHook[types.Project](true, nil, func(id any) error {
		p := project.Get(id.(string))
		if p == nil {
			return errors.New("项目未加载")
		}
		//err := p.Close()
		//if err != nil {
		//	return err
		//}
		return nil
	}))

	app.GET("/:id/enable", curd.ParseParamStringId, curd.ApiDisableHook[types.Project](false, nil, func(id any) error {
		return project.Load(id.(string))
	}))

	app.GET("/:id/start", curd.ParseParamStringId, func(ctx *gin.Context) {
		err := project.Load(ctx.GetString("id"))
		if err != nil {
			curd.Error(ctx, err)
			return
		}
		curd.OK(ctx, nil)
	})

	app.GET("/:id/stop", curd.ParseParamStringId, func(ctx *gin.Context) {
		p := project.Get(ctx.GetString("id"))
		if p == nil {
			curd.Fail(ctx, "项目未加载")
			return
		}
		//err := p.Close()
		//if err != nil {
		//	curd.Error(ctx, err)
		//	return
		//}
		curd.OK(ctx, nil)
	})

	//附件
	attach.ObjectRouters("project", app)
}
