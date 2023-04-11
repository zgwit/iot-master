package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/v3/model"
	"github.com/zgwit/iot-master/v3/pkg/curd"
)

// @Summary 查询分组数量
// @Schemes
// @Description 查询分组数量
// @Tags device-group
// @Param search body curd.ParamSearch true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[int64] 返回分组数量
// @Router /device/group/count [post]
func noopDeviceGroupCount() {}

// @Summary 查询分组
// @Schemes
// @Description 查询分组
// @Tags device-group
// @Param search body curd.ParamSearch true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyList[model.DeviceGroup] 返回分组信息
// @Router /device/group/search [post]
func noopDeviceGroupSearch() {}

// @Summary 查询分组
// @Schemes
// @Description 查询分组
// @Tags device-group
// @Param search query curd.ParamList true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyList[model.DeviceGroup] 返回分组信息
// @Router /device/group/list [get]
func noopDeviceGroupList() {}

// @Summary 创建分组
// @Schemes
// @Description 创建分组
// @Tags device-group
// @Param search body model.DeviceGroup true "分组信息"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[model.DeviceGroup] 返回分组信息
// @Router /device/group/create [post]
func noopDeviceGroupCreate() {}

// @Summary 修改分组
// @Schemes
// @Description 修改分组
// @Tags device-group
// @Param id path int true "分组ID"
// @Param device-group body model.DeviceGroup true "分组信息"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[model.DeviceGroup] 返回分组信息
// @Router /device/group/{id} [post]
func noopDeviceGroupUpdate() {}

// @Summary 获取分组
// @Schemes
// @Description 获取分组
// @Tags device-group
// @Param id path int true "分组ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[model.DeviceGroup] 返回分组信息
// @Router /device/group/{id} [get]
func noopDeviceGroupGet() {}

// @Summary 删除分组
// @Schemes
// @Description 删除分组
// @Tags device-group
// @Param id path int true "分组ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[model.DeviceGroup] 返回分组信息
// @Router /device/group/{id}/delete [get]
func noopDeviceGroupDelete() {}

// @Summary 导出分组
// @Schemes
// @Description 导出分组
// @Tags device-group
// @Accept json
// @Produce octet-stream
// @Router /device/group/export [get]
func noopDeviceGroupExport() {}

// @Summary 导入分组
// @Schemes
// @Description 导入分组
// @Tags device-group
// @Param file formData file true "压缩包"
// @Accept mpfd
// @Produce json
// @Success 200 {object} curd.ReplyData[int64] 返回分组数量
// @Router /device/group/import [post]
func noopDeviceGroupImport() {}

func deviceGroupRouter(app *gin.RouterGroup) {

	app.POST("/count", curd.ApiCount[model.DeviceGroup]())
	app.POST("/search", curd.ApiSearch[model.DeviceGroup]())
	app.GET("/list", curd.ApiList[model.DeviceGroup]())
	app.POST("/create", curd.ApiCreate[model.DeviceGroup](nil, nil))
	app.GET("/:id", curd.ParseParamId, curd.ApiGet[model.DeviceGroup]())
	app.POST("/:id", curd.ParseParamId, curd.ApiModify[model.DeviceGroup](nil, nil,
		"name", "desc"))
	app.GET("/:id/delete", curd.ParseParamId, curd.ApiDelete[model.DeviceGroup](nil, nil))
	app.GET("/export", curd.ApiExport[model.DeviceGroup]("device-group"))
	app.POST("/import", curd.ApiImport[model.DeviceGroup]())
}
