package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/v4/project"
	"github.com/zgwit/iot-master/v4/web/curd"
)

// @Summary 查询项目
// @Schemes
// @Description 查询项目
// @Tags project
// @Param search body curd.ParamSearch true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyList[project.Project] 返回项目信息
// @Router /project/search [post]
func noopProjectSearch() {}

// @Summary 查询项目
// @Schemes
// @Description 查询项目
// @Tags project
// @Param search query curd.ParamList true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyList[project.Project] 返回项目信息
// @Router /project/list [get]
func noopProjectList() {}

// @Summary 创建项目
// @Schemes
// @Description 创建项目
// @Tags project
// @Param search body project.Project true "项目信息"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[project.Project] 返回项目信息
// @Router /project/create [post]
func noopProjectCreate() {}

// @Summary 修改项目
// @Schemes
// @Description 修改项目
// @Tags project
// @Param id path int true "项目ID"
// @Param project body project.Project true "项目信息"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[project.Project] 返回项目信息
// @Router /project/{id} [post]
func noopProjectUpdate() {}

// @Summary 删除项目
// @Schemes
// @Description 删除项目
// @Tags project
// @Param id path int true "项目ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[project.Project] 返回项目信息
// @Router /project/{id}/delete [get]
func noopProjectDelete() {}

// @Summary 启用项目
// @Schemes
// @Description 启用项目
// @Tags project
// @Param id path int true "项目ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[project.Project] 返回项目信息
// @Router /project/{id}/enable [get]
func noopProjectEnable() {}

// @Summary 禁用项目
// @Schemes
// @Description 禁用项目
// @Tags project
// @Param id path int true "项目ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[project.Project] 返回项目信息
// @Router /project/{id}/disable [get]
func noopProjectDisable() {}

// @Summary 启动项目
// @Schemes
// @Description 启动项目
// @Tags project
// @Param id path int true "项目ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[project.Project] 返回项目信息
// @Router /project/{id}/start [get]
func noopProjectStart() {}

// @Summary 停止项目
// @Schemes
// @Description 停止项目
// @Tags project
// @Param id path int true "项目ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[project.Project] 返回项目信息
// @Router /project/{id}/stop [get]
func noopProjectStop() {}

func projectRouter(app *gin.RouterGroup) {

	app.POST("/count", curd.ApiCount[project.Project]())

	app.POST("/search", curd.ApiSearchHook[project.Project](func(datum []*project.Project) error {
		for _, v := range datum {
			p := project.Get(v.Id)
			if p != nil {
				//v.Running = p.Running
			}
		}
		return nil
	}))

	app.GET("/list", curd.ApiListHook[project.Project](func(datum []*project.Project) error {
		for _, v := range datum {
			p := project.Get(v.Id)
			if p != nil {
				//v.Running = p.Running
			}
		}
		return nil
	}))
	app.POST("/create", curd.ApiCreateHook[project.Project](curd.GenerateID[project.Project](), nil))

	app.GET("/:id", curd.ParseParamStringId, curd.ApiGetHook[project.Project](func(m *project.Project) error {
		p := project.Get(m.Id)
		if p != nil {
			//m.Running = p.Running
		}
		return nil
	}))

	app.POST("/:id", curd.ParseParamStringId, curd.ApiUpdateHook[project.Project](nil, nil,
		"id", "name", "icon", "description", "keywords", "disabled"))

	app.GET("/:id/delete", curd.ParseParamStringId, curd.ApiDeleteHook[project.Project](nil, nil))

	app.GET(":id/disable", curd.ParseParamStringId, curd.ApiDisableHook[project.Project](true, nil, func(id any) error {
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

	app.GET("/:id/enable", curd.ParseParamStringId, curd.ApiDisableHook[project.Project](false, nil, func(id any) error {
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
}
