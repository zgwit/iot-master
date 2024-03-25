package space

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/v4/api"
	"github.com/zgwit/iot-master/v4/web/curd"
)

func init() {

	api.Register("POST", "/space/count", curd.ApiCount[Space]())

	api.Register("POST", "/space/search", curd.ApiSearchWith[Space]([]*curd.With{
		{"project", "project_id", "id", "name", "project"},
	}, "id", "name", "project_id", "description", "disabled", "created"))

	api.Register("GET", "/space/list", curd.ApiList[Space]())

	api.Register("POST", "/space/create", curd.ApiCreateHook[Space](curd.GenerateID[Space](), nil))

	api.Register("GET", "/space/:id", curd.ParseParamStringId, curd.ApiGet[Space]())

	api.Register("POST", "/space/:id", curd.ParseParamStringId, curd.ApiUpdateHook[Space](nil, nil,
		"id", "name", "project_id", "description", "disabled"))

	api.Register("GET", "/space/:id/delete", curd.ParseParamStringId, curd.ApiDeleteHook[Space](nil, nil))

	api.Register("GET", "/space/:id/disable", curd.ParseParamStringId, curd.ApiDisableHook[Space](true, nil, func(id any) error {
		p := Get(id.(string))
		if p == nil {
			return errors.New("空间未加载")
		}
		//err := p.Close()
		//if err != nil {
		//	return err
		//}
		return nil
	}))

	api.Register("GET", "/space/:id/enable", curd.ParseParamStringId, curd.ApiDisableHook[Space](false, nil, func(id any) error {
		return Load(id.(string))
	}))

	api.Register("GET", "/space/:id/start", curd.ParseParamStringId, func(ctx *gin.Context) {
		err := Load(ctx.GetString("id"))
		if err != nil {
			curd.Error(ctx, err)
			return
		}
		curd.OK(ctx, nil)
	})

	api.Register("GET", "/space/:id/stop", curd.ParseParamStringId, func(ctx *gin.Context) {
		p := Get(ctx.GetString("id"))
		if p == nil {
			curd.Fail(ctx, "空间未加载")
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

// @Summary 查询空间
// @Schemes
// @Description 查询空间
// @Tags space
// @Param search body curd.ParamSearch true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyList[Space] 返回空间信息
// @Router /space/search [post]
func noopSpaceSearch() {}

// @Summary 查询空间
// @Schemes
// @Description 查询空间
// @Tags space
// @Param search query curd.ParamList true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyList[Space] 返回空间信息
// @Router /space/list [get]
func noopSpaceList() {}

// @Summary 创建空间
// @Schemes
// @Description 创建空间
// @Tags space
// @Param search body Space true "空间信息"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[Space] 返回空间信息
// @Router /space/create [post]
func noopSpaceCreate() {}

// @Summary 修改空间
// @Schemes
// @Description 修改空间
// @Tags space
// @Param id path int true "空间ID"
// @Param space body Space true "空间信息"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[Space] 返回空间信息
// @Router /space/{id} [post]
func noopSpaceUpdate() {}

// @Summary 删除空间
// @Schemes
// @Description 删除空间
// @Tags space
// @Param id path int true "空间ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[Space] 返回空间信息
// @Router /space/{id}/delete [get]
func noopSpaceDelete() {}

// @Summary 启用空间
// @Schemes
// @Description 启用空间
// @Tags space
// @Param id path int true "空间ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[Space] 返回空间信息
// @Router /space/{id}/enable [get]
func noopSpaceEnable() {}

// @Summary 禁用空间
// @Schemes
// @Description 禁用空间
// @Tags space
// @Param id path int true "空间ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[Space] 返回空间信息
// @Router /space/{id}/disable [get]
func noopSpaceDisable() {}

// @Summary 启动空间
// @Schemes
// @Description 启动空间
// @Tags space
// @Param id path int true "空间ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[Space] 返回空间信息
// @Router /space/{id}/start [get]
func noopSpaceStart() {}

// @Summary 停止空间
// @Schemes
// @Description 停止空间
// @Tags space
// @Param id path int true "空间ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[Space] 返回空间信息
// @Router /space/{id}/stop [get]
func noopSpaceStop() {}
