package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/v4/pkg/db"
	"github.com/zgwit/iot-master/v4/pkg/web/curd"
	"github.com/zgwit/iot-master/v4/types"
	"xorm.io/xorm/schemas"
)

// @Summary 项目设备列表
// @Schemes
// @Description 项目设备列表
// @Tags project-device
// @Param id path int true "项目ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[[]types.ProjectDevice] 返回项目设备信息
// @Router /project/{id}/device/{device} [get]
func projectDeviceList(ctx *gin.Context) {
	var pds []types.ProjectDevice
	err := db.Engine.Where("project_id=?", ctx.Param("id")).Find(&pds)
	if err != nil {
		curd.Error(ctx, err)
		return
	}
	curd.OK(ctx, pds)
}

// @Summary 绑定项目设备
// @Schemes
// @Description 绑定项目设备
// @Tags project-device
// @Param id path int true "项目ID"
// @Param device path int true "设备ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[int]
// @Router /project/{id}/device/{device}/bind [get]
func projectDeviceBind(ctx *gin.Context) {
	pd := types.ProjectDevice{
		ProjectId: ctx.Param("id"),
		DeviceId:  ctx.Param("device"),
		Name:      ctx.Query("name"),
	}
	_, err := db.Engine.InsertOne(&pd)
	if err != nil {
		curd.Error(ctx, err)
		return
	}
	curd.OK(ctx, nil)
}

// @Summary 删除项目设备
// @Schemes
// @Description 删除项目设备
// @Tags project-device
// @Param id path int true "项目ID"
// @Param device path int true "设备ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[int]
// @Router /project/{id}/device/{device}/unbind [get]
func projectDeviceUnbind(ctx *gin.Context) {
	_, err := db.Engine.ID(schemas.PK{ctx.Param("id"), ctx.Param("device")}).Delete(new(types.ProjectDevice))
	if err != nil {
		curd.Error(ctx, err)
		return
	}
	curd.OK(ctx, nil)
}

// @Summary 修改项目设备
// @Schemes
// @Description 修改项目设备
// @Tags project-device
// @Param id path int true "项目ID"
// @Param device path int true "设备ID"
// @Param project-device body types.ProjectDevice true "项目设备信息"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[int]
// @Router /project/{id}/device/{device} [post]
func projectDeviceUpdate(ctx *gin.Context) {
	var pd types.ProjectDevice
	err := ctx.ShouldBindJSON(&pd)
	if err != nil {
		curd.Error(ctx, err)
		return
	}
	_, err = db.Engine.ID(schemas.PK{ctx.Param("id"), ctx.Param("device")}).
		Cols("device_id", "name").
		Update(&pd)
	if err != nil {
		curd.Error(ctx, err)
		return
	}
	curd.OK(ctx, nil)
}

func projectDeviceRouter(app *gin.RouterGroup) {
	app.GET("", projectDeviceList)
	app.GET("/:device/bind", projectDeviceBind)
	app.GET("/:device/unbind", projectDeviceUnbind)
	app.POST("/:device", projectDeviceUpdate)
}
