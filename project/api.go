package project

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/v4/api"
	"github.com/zgwit/iot-master/v4/web/curd"
)

func init() {

	api.Register("POST", "/project/count", curd.ApiCount[Project]())

	api.Register("POST", "/project/search", curd.ApiSearchHook[Project](func(datum []*Project) error {
		for _, v := range datum {
			p := Get(v.Id)
			if p != nil {
				//v.running = p.running
			}
		}
		return nil
	}))

	api.Register("GET", "/project/list", curd.ApiListHook[Project](func(datum []*Project) error {
		for _, v := range datum {
			p := Get(v.Id)
			if p != nil {
				//v.running = p.running
			}
		}
		return nil
	}))
	api.Register("POST", "/project/create", curd.ApiCreateHook[Project](curd.GenerateID[Project](), nil))

	api.Register("GET", "/project/:id", curd.ParseParamStringId, curd.ApiGetHook[Project](func(m *Project) error {
		p := Get(m.Id)
		if p != nil {
			//m.running = p.running
		}
		return nil
	}))

	api.Register("POST", "/project/:id", curd.ParseParamStringId, curd.ApiUpdate[Project]())

	api.Register("GET", "/project/:id/delete", curd.ParseParamStringId, curd.ApiDeleteHook[Project](nil, nil))

	api.Register("GET", "/project/:id/disable", curd.ParseParamStringId, curd.ApiDisableHook[Project](true, nil, func(id any) error {
		p := Get(id.(string))
		if p == nil {
			return errors.New("项目未加载")
		}
		//err := p.Close()
		//if err != nil {
		//	return err
		//}
		return nil
	}))

	api.Register("GET", "/project/:id/enable", curd.ParseParamStringId, curd.ApiDisableHook[Project](false, nil, func(id any) error {
		return Load(id.(string))
	}))

	api.Register("GET", "/project/:id/start", curd.ParseParamStringId, func(ctx *gin.Context) {
		err := Load(ctx.GetString("id"))
		if err != nil {
			curd.Error(ctx, err)
			return
		}
		curd.OK(ctx, nil)
	})

	api.Register("GET", "/project/:id/stop", curd.ParseParamStringId, func(ctx *gin.Context) {
		p := Get(ctx.GetString("id"))
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

// @Summary 查询项目
// @Schemes
// @Description 查询项目
// @Tags project
// @Param search body curd.ParamSearch true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyList[Project] 返回项目信息
// @Router /project/search [post]
func noopProjectSearch() {}

// @Summary 查询项目
// @Schemes
// @Description 查询项目
// @Tags project
// @Param search query curd.ParamList true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyList[Project] 返回项目信息
// @Router /project/list [get]
func noopProjectList() {}

// @Summary 创建项目
// @Schemes
// @Description 创建项目
// @Tags project
// @Param search body Project true "项目信息"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[Project] 返回项目信息
// @Router /project/create [post]
func noopProjectCreate() {}

// @Summary 修改项目
// @Schemes
// @Description 修改项目
// @Tags project
// @Param id path int true "项目ID"
// @Param project body Project true "项目信息"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[Project] 返回项目信息
// @Router /project/{id} [post]
func noopProjectUpdate() {}

// @Summary 删除项目
// @Schemes
// @Description 删除项目
// @Tags project
// @Param id path int true "项目ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[Project] 返回项目信息
// @Router /project/{id}/delete [get]
func noopProjectDelete() {}

// @Summary 启用项目
// @Schemes
// @Description 启用项目
// @Tags project
// @Param id path int true "项目ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[Project] 返回项目信息
// @Router /project/{id}/enable [get]
func noopProjectEnable() {}

// @Summary 禁用项目
// @Schemes
// @Description 禁用项目
// @Tags project
// @Param id path int true "项目ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[Project] 返回项目信息
// @Router /project/{id}/disable [get]
func noopProjectDisable() {}

// @Summary 启动项目
// @Schemes
// @Description 启动项目
// @Tags project
// @Param id path int true "项目ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[Project] 返回项目信息
// @Router /project/{id}/start [get]
func noopProjectStart() {}

// @Summary 停止项目
// @Schemes
// @Description 停止项目
// @Tags project
// @Param id path int true "项目ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[Project] 返回项目信息
// @Router /project/{id}/stop [get]
func noopProjectStop() {}
