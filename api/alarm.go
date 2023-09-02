package api

import (
	"github.com/gin-gonic/gin"
	curd2 "github.com/zgwit/iot-master/v4/curd"
	"github.com/zgwit/iot-master/v4/db"
	"github.com/zgwit/iot-master/v4/model"
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
// @Success 200 {object} curd.ReplyList[model.Alarm] 返回报警信息
// @Router /alarm/search [post]
func noopAlarmSearch() {}

// @Summary 查询报警
// @Schemes
// @Description 查询报警
// @Tags alarm
// @Param search query curd.ParamList true "查询参数"
// @Produce json
// @Success 200 {object} curd.ReplyList[model.Alarm] 返回报警信息
// @Router /alarm/list [get]
func noopAlarmList() {}

// @Summary 删除报警
// @Schemes
// @Description 删除报警
// @Tags alarm
// @Param id path int true "报警ID"
// @Produce json
// @Success 200 {object} curd.ReplyData[model.Alarm] 返回报警信息
// @Router /alarm/{id}/delete [get]
func noopAlarmDelete() {}

// @Summary 阅读报警
// @Schemes
// @Description 阅读报警
// @Tags alarm
// @Param id path int true "报警ID"
// @Produce json
// @Success 200 {object} curd.ReplyData[model.Alarm] 返回报警信息
// @Router /alarm/{id}/read [get]
func noopAlarmRead() {}

func alarmRead(ctx *gin.Context) {
	alarm := model.Alarm{
		Read: true,
	}
	cnt, err := db.Engine.ID(ctx.GetInt64("id")).Cols("read").Update(alarm)
	if err != nil {
		curd2.Error(ctx, err)
		return
	}
	curd2.OK(ctx, cnt)
}

func alarmRouter(app *gin.RouterGroup) {

	app.POST("/count", curd2.ApiCount[model.Alarm]())

	app.POST("/search", curd2.ApiSearchWith[model.AlarmEx]([]*curd2.Join{
		{"product", "product_id", "id", "name", "product"},
		{"device", "device_id", "id", "name", "device"},
	}))

	app.GET("/list", curd2.ApiList[model.Alarm]())

	app.GET("/:id", curd2.ParseParamId, curd2.ApiGet[model.Alarm]())

	app.GET("/:id/delete", curd2.ParseParamId, curd2.ApiDelete[model.Alarm]())

	app.GET("/:id/read", curd2.ParseParamId, alarmRead)
}
