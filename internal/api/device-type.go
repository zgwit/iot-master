package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/v3/model"
	"github.com/zgwit/iot-master/v3/pkg/curd"
)

// @Summary 查询设备类型数量
// @Schemes
// @Description 查询设备类型数量
// @Tags device-type
// @Param search body curd.ParamSearch true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[int64] 返回设备类型数量
// @Router /device/type/count [post]
func noopDeviceTypeCount() {}

// @Summary 查询设备类型
// @Schemes
// @Description 查询设备类型
// @Tags device-type
// @Param search body curd.ParamSearch true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyList[model.DeviceType] 返回设备类型信息
// @Router /device/type/search [post]
func noopDeviceTypeSearch() {}

// @Summary 查询设备类型
// @Schemes
// @Description 查询设备类型
// @Tags device-type
// @Param search query curd.ParamList true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyList[model.DeviceType] 返回设备类型信息
// @Router /device/type/list [get]
func noopDeviceTypeList() {}

// @Summary 创建设备类型
// @Schemes
// @Description 创建设备类型
// @Tags device-type
// @Param search body model.DeviceType true "设备类型信息"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[model.DeviceType] 返回设备类型信息
// @Router /device/type/create [post]
func noopDeviceTypeCreate() {}

// @Summary 修改设备类型
// @Schemes
// @Description 修改设备类型
// @Tags device-type
// @Param id path string true "设备类型ID"
// @Param device-type body model.DeviceType true "设备类型信息"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[model.DeviceType] 返回设备类型信息
// @Router /device/type/{id} [post]
func noopDeviceTypeUpdate() {}

// @Summary 获取设备类型
// @Schemes
// @Description 获取设备类型
// @Tags device-type
// @Param id path string true "设备类型ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[model.DeviceType] 返回设备类型信息
// @Router /device/type/{id} [get]
func noopDeviceTypeGet() {}

// @Summary 删除设备类型
// @Schemes
// @Description 删除设备类型
// @Tags device-type
// @Param id path string true "设备类型ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[model.DeviceType] 返回设备类型信息
// @Router /device/type/{id}/delete [get]
func noopDeviceTypeDelete() {}

// @Summary 导出设备类型
// @Schemes
// @Description 导出设备类型
// @Tags device-type
// @Accept json
// @Produce octet-stream
// @Router /device/type/export [get]
func noopDeviceTypeExport() {}

// @Summary 导入设备类型
// @Schemes
// @Description 导入设备类型
// @Tags device-type
// @Param file formData file true "压缩包"
// @Accept mpfd
// @Produce json
// @Success 200 {object} curd.ReplyData[int64] 返回设备类型数量
// @Router /device/type/import [post]
func noopDeviceTypeImport() {}

func deviceTypeRouter(app *gin.RouterGroup) {

	app.POST("/count", curd.ApiCount[model.DeviceType]())
	app.POST("/search", curd.ApiSearch[model.DeviceType]())
	app.GET("/list", curd.ApiList[model.DeviceType]())
	app.POST("/create", curd.ApiCreateHook[model.DeviceType](curd.GenerateRandomId[model.DeviceType](8), nil))
	app.GET("/:id", curd.ParseParamStringId, curd.ApiGet[model.DeviceType]())
	app.POST("/:id", curd.ParseParamStringId, curd.ApiUpdateHook[model.DeviceType](nil, nil,
		"name", "desc"))
	app.GET("/:id/delete", curd.ParseParamStringId, curd.ApiDeleteHook[model.DeviceType](nil, nil))
	app.GET("/export", curd.ApiExport[model.DeviceType]("device-type"))
	app.POST("/import", curd.ApiImport[model.DeviceType]())
}
