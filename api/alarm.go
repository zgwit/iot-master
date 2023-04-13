package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/v3/model"
	"github.com/zgwit/iot-master/v3/pkg/curd"
	"github.com/zgwit/iot-master/v3/pkg/db"
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
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyList[model.Alarm] 返回报警信息
// @Router /alarm/list [get]
func noopAlarmList() {}

// @Summary 删除报警
// @Schemes
// @Description 删除报警
// @Tags alarm
// @Param id path int true "报警ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[model.Alarm] 返回报警信息
// @Router /alarm/{id}/delete [get]
func noopAlarmDelete() {}

// @Summary 阅读报警
// @Schemes
// @Description 阅读报警
// @Tags alarm
// @Param id path int true "报警ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[model.Alarm] 返回报警信息
// @Router /alarm/{id}/read [get]
func noopAlarmRead() {}

func alarmRouter(app *gin.RouterGroup) {

	app.POST("/count", curd.ApiCount[model.Alarm]())

	app.POST("/search", curd.ApiSearch[model.Alarm]())

	app.GET("/list", curd.ApiList[model.Alarm]())

	app.GET("/:id", curd.ParseParamId, curd.ApiGet[model.Alarm]())

	app.GET("/:id/delete", curd.ParseParamId, curd.ApiDelete[model.Alarm](nil, nil))

	app.GET("/:id/read", curd.ParseParamId, alarmRead)
}

func alarmRead(ctx *gin.Context) {
	alarm := model.Alarm{
		Read: true,
	}
	cnt, err := db.Engine.ID(ctx.GetInt64("id")).Cols("read").Update(alarm)
	if err != nil {
		curd.Error(ctx, err)
		return
	}
	curd.OK(ctx, cnt)
}
