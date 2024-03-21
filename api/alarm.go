package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/v4/alarm"
	"github.com/zgwit/iot-master/v4/pkg/db"
	"github.com/zgwit/iot-master/v4/web/curd"
)

// @Summary 查询报警
// @Schemes
// @Description 查询报警
// @Tags alarm
// @Param search body curd.ParamSearch true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[int64] 返回报警信息
// @Router /alarm/count [post]
func noopAlarmCount() {}

// @Summary 查询报警
// @Schemes
// @Description 查询报警
// @Tags alarm
// @Param search body curd.ParamSearch true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyList[alarm.Alarm] 返回报警信息
// @Router /alarm/search [post]
func noopAlarmSearch() {}

// @Summary 查询报警
// @Schemes
// @Description 查询报警
// @Tags alarm
// @Param search query curd.ParamList true "查询参数"
// @Produce json
// @Success 200 {object} curd.ReplyList[alarm.Alarm] 返回报警信息
// @Router /alarm/list [get]
func noopAlarmList() {}

// @Summary 删除报警
// @Schemes
// @Description 删除报警
// @Tags alarm
// @Param id path int true "报警ID"
// @Produce json
// @Success 200 {object} curd.ReplyData[alarm.Alarm] 返回报警信息
// @Router /alarm/{id}/delete [get]
func noopAlarmDelete() {}

// @Summary 阅读报警
// @Schemes
// @Description 阅读报警
// @Tags alarm
// @Param id path int true "报警ID"
// @Produce json
// @Success 200 {object} curd.ReplyData[alarm.Alarm] 返回报警信息
// @Router /alarm/{id}/read [get]
func noopAlarmRead() {}

func alarmRead(ctx *gin.Context) {
	a := alarm.Alarm{
		Read: true,
	}
	cnt, err := db.Engine.ID(ctx.GetInt64("id")).Cols("read").Update(a)
	if err != nil {
		curd.Error(ctx, err)
		return
	}
	curd.OK(ctx, cnt)
}

func alarmRouter(app *gin.RouterGroup) {

	app.POST("/count", curd.ApiCount[alarm.Alarm]())

	app.POST("/search", curd.ApiSearchWith[alarm.Alarm]([]*curd.With{
		{"product", "product_id", "id", "name", "product"},
		{"project", "project_id", "id", "name", "project"},
		{"space", "space_id", "id", "name", "space"},
		{"device", "device_id", "id", "name", "device"},
	}))

	app.GET("/list", curd.ApiList[alarm.Alarm]())

	app.GET("/:id", curd.ParseParamId, curd.ApiGet[alarm.Alarm]())

	app.GET("/:id/delete", curd.ParseParamId, curd.ApiDelete[alarm.Alarm]())

	app.GET("/:id/read", curd.ParseParamId, alarmRead)
}
