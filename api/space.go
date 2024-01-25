package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/v4/space"
	"github.com/zgwit/iot-master/v4/types"
	"github.com/zgwit/iot-master/v4/web/attach"
	"github.com/zgwit/iot-master/v4/web/curd"
	"github.com/zgwit/iot-master/v4/web/export"
)

// @Summary 查询空间
// @Schemes
// @Description 查询空间
// @Tags space
// @Param search body curd.ParamSearch true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyList[types.Space] 返回空间信息
// @Router /space/search [post]
func noopSpaceSearch() {}

// @Summary 查询空间
// @Schemes
// @Description 查询空间
// @Tags space
// @Param search query curd.ParamList true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyList[types.Space] 返回空间信息
// @Router /space/list [get]
func noopSpaceList() {}

// @Summary 创建空间
// @Schemes
// @Description 创建空间
// @Tags space
// @Param search body types.Space true "空间信息"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[types.Space] 返回空间信息
// @Router /space/create [post]
func noopSpaceCreate() {}

// @Summary 修改空间
// @Schemes
// @Description 修改空间
// @Tags space
// @Param id path int true "空间ID"
// @Param space body types.Space true "空间信息"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[types.Space] 返回空间信息
// @Router /space/{id} [post]
func noopSpaceUpdate() {}

// @Summary 删除空间
// @Schemes
// @Description 删除空间
// @Tags space
// @Param id path int true "空间ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[types.Space] 返回空间信息
// @Router /space/{id}/delete [get]
func noopSpaceDelete() {}

// @Summary 导出空间
// @Schemes
// @Description 导出空间
// @Tags space
// @Accept json
// @Produce octet-stream
// @Router /space/export [get]
func noopSpaceExport() {}

// @Summary 导入空间
// @Schemes
// @Description 导入空间
// @Tags space
// @Param file formData file true "压缩包"
// @Accept mpfd
// @Produce json
// @Success 200 {object} curd.ReplyData[int64] 返回空间数量
// @Router /space/import [post]
func noopSpaceImport() {}

// @Summary 启用空间
// @Schemes
// @Description 启用空间
// @Tags space
// @Param id path int true "空间ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[types.Space] 返回空间信息
// @Router /space/{id}/enable [get]
func noopSpaceEnable() {}

// @Summary 禁用空间
// @Schemes
// @Description 禁用空间
// @Tags space
// @Param id path int true "空间ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[types.Space] 返回空间信息
// @Router /space/{id}/disable [get]
func noopSpaceDisable() {}

// @Summary 启动空间
// @Schemes
// @Description 启动空间
// @Tags space
// @Param id path int true "空间ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[types.Space] 返回空间信息
// @Router /space/{id}/start [get]
func noopSpaceStart() {}

// @Summary 停止空间
// @Schemes
// @Description 停止空间
// @Tags space
// @Param id path int true "空间ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[types.Space] 返回空间信息
// @Router /space/{id}/stop [get]
func noopSpaceStop() {}

func spaceRouter(app *gin.RouterGroup) {

	app.POST("/search", curd.ApiSearchWith[types.Space]([]*curd.Join{
		{"project", "project_id", "id", "name", "project"},
	}, "id", "name", "project_id", "description", "disabled", "created"))

	app.GET("/list", curd.ApiList[types.Space]())

	app.POST("/create", curd.ApiCreateHook[types.Space](curd.GenerateKSUID[types.Space](), nil))

	app.GET("/:id", curd.ParseParamStringId, curd.ApiGet[types.Space]())

	app.POST("/:id", curd.ParseParamStringId, curd.ApiUpdateHook[types.Space](nil, nil,
		"id", "name", "project_id", "description", "disabled"))

	app.GET("/:id/delete", curd.ParseParamStringId, curd.ApiDeleteHook[types.Space](nil, nil))

	app.GET("/export", export.ApiExport("space", "空间"))

	app.POST("/import", export.ApiImport("space"))

	app.GET(":id/disable", curd.ParseParamStringId, curd.ApiDisableHook[types.Space](true, nil, func(id any) error {
		p := space.Get(id.(string))
		if p == nil {
			return errors.New("空间未加载")
		}
		//err := p.Close()
		//if err != nil {
		//	return err
		//}
		return nil
	}))

	app.GET("/:id/enable", curd.ParseParamStringId, curd.ApiDisableHook[types.Space](false, nil, func(id any) error {
		return space.Load(id.(string))
	}))

	app.GET("/:id/start", curd.ParseParamStringId, func(ctx *gin.Context) {
		err := space.Load(ctx.GetString("id"))
		if err != nil {
			curd.Error(ctx, err)
			return
		}
		curd.OK(ctx, nil)
	})

	app.GET("/:id/stop", curd.ParseParamStringId, func(ctx *gin.Context) {
		p := space.Get(ctx.GetString("id"))
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

	//附件
	attach.ObjectRouters("space", app)
}
