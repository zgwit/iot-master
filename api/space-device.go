package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/v4/pkg/db"
	"github.com/zgwit/iot-master/v4/types"
	"github.com/zgwit/iot-master/v4/web/curd"
	"xorm.io/xorm/schemas"
)

// @Summary 空间设备列表
// @Schemes
// @Description 空间设备列表
// @Tags space-device
// @Param id path int true "项目ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[[]types.SpaceDevice] 返回空间设备信息
// @Router /space/{id}/device [get]
func spaceDeviceList(ctx *gin.Context) {
	var pds []types.SpaceDevice
	err := db.Engine.
		Select("space_device.space_id, space_device.device_id, space_device.name, space_device.created, device.name as device").
		Join("INNER", "device", "device.id=space_device.device_id").
		Where("space_device.space_id=?", ctx.Param("id")).
		Find(&pds)
	if err != nil {
		curd.Error(ctx, err)
		return
	}
	curd.OK(ctx, pds)
}

// @Summary 绑定空间设备
// @Schemes
// @Description 绑定空间设备
// @Tags space-device
// @Param id path int true "项目ID"
// @Param device path int true "设备ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[int]
// @Router /space/{id}/device/{device}/bind [get]
func spaceDeviceBind(ctx *gin.Context) {
	pd := types.SpaceDevice{
		SpaceId:  ctx.Param("id"),
		DeviceId: ctx.Param("device"),
	}
	_, err := db.Engine.InsertOne(&pd)
	if err != nil {
		curd.Error(ctx, err)
		return
	}
	curd.OK(ctx, nil)
}

// @Summary 删除空间设备
// @Schemes
// @Description 删除空间设备
// @Tags space-device
// @Param id path int true "项目ID"
// @Param device path int true "设备ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[int]
// @Router /space/{id}/device/{device}/unbind [get]
func spaceDeviceUnbind(ctx *gin.Context) {
	_, err := db.Engine.ID(schemas.PK{ctx.Param("id"), ctx.Param("device")}).Delete(new(types.SpaceDevice))
	if err != nil {
		curd.Error(ctx, err)
		return
	}
	curd.OK(ctx, nil)
}

// @Summary 修改空间设备
// @Schemes
// @Description 修改空间设备
// @Tags space-device
// @Param id path int true "项目ID"
// @Param device path int true "设备ID"
// @Param space-device body types.SpaceDevice true "空间设备信息"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[int]
// @Router /space/{id}/device/{device} [post]
func spaceDeviceUpdate(ctx *gin.Context) {
	var pd types.SpaceDevice
	err := ctx.ShouldBindJSON(&pd)
	if err != nil {
		curd.Error(ctx, err)
		return
	}
	_, err = db.Engine.ID(schemas.PK{ctx.Param("id"), ctx.Param("device")}).
		Cols("device_id", "name", "disabled").
		Update(&pd)
	if err != nil {
		curd.Error(ctx, err)
		return
	}
	curd.OK(ctx, nil)
}

func spaceDeviceRouter(app *gin.RouterGroup) {
	app.GET("", spaceDeviceList)
	app.GET("/:device/bind", spaceDeviceBind)
	app.GET("/:device/unbind", spaceDeviceUnbind)
	app.POST("/:device", spaceDeviceUpdate)
}
