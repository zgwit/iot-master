package alarm

import (
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/v4/api"
	"github.com/zgwit/iot-master/v4/pkg/db"
	"github.com/zgwit/iot-master/v4/web/curd"
)

func init() {
	api.Register("POST", "/alarm/count", curd.ApiCount[Alarm]())

	api.Register("POST", "/alarm/search", curd.ApiSearchWith[Alarm]([]*curd.With{
		{"product", "product_id", "id", "name", "product"},
		{"project", "project_id", "id", "name", "project"},
		{"space", "space_id", "id", "name", "space"},
		{"device", "device_id", "id", "name", "device"},
	}))

	api.Register("GET", "/alarm/list", curd.ApiList[Alarm]())

	api.Register("GET", "/alarm/:id", curd.ParseParamId, curd.ApiGet[Alarm]())

	api.Register("GET", "/alarm/:id/delete", curd.ParseParamId, curd.ApiDelete[Alarm]())

	api.Register("GET", "/alarm/:id/read", curd.ParseParamId, alarmRead)
}

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
// @Success 200 {object} curd.ReplyList[Alarm] 返回报警信息
// @Router /alarm/search [post]
func noopAlarmSearch() {}

// @Summary 查询报警
// @Schemes
// @Description 查询报警
// @Tags alarm
// @Param search query curd.ParamList true "查询参数"
// @Produce json
// @Success 200 {object} curd.ReplyList[Alarm] 返回报警信息
// @Router /alarm/list [get]
func noopAlarmList() {}

// @Summary 删除报警
// @Schemes
// @Description 删除报警
// @Tags alarm
// @Param id path int true "报警ID"
// @Produce json
// @Success 200 {object} curd.ReplyData[Alarm] 返回报警信息
// @Router /alarm/{id}/delete [get]
func noopAlarmDelete() {}

// @Summary 阅读报警
// @Schemes
// @Description 阅读报警
// @Tags alarm
// @Param id path int true "报警ID"
// @Produce json
// @Success 200 {object} curd.ReplyData[Alarm] 返回报警信息
// @Router /alarm/{id}/read [get]
func noopAlarmRead() {}

func alarmRead(ctx *gin.Context) {
	a := Alarm{
		Read: true,
	}
	cnt, err := db.Engine.ID(ctx.GetInt64("id")).Cols("read").Update(a)
	if err != nil {
		curd.Error(ctx, err)
		return
	}
	curd.OK(ctx, cnt)
}
