package api

import (
	"github.com/gin-gonic/gin"
	curd2 "github.com/zgwit/iot-master/v4/curd"
	"github.com/zgwit/iot-master/v4/model"
)

// @Summary 查询历史数量
// @Schemes
// @Description 查询历史
// @Tags history
// @Param search body curd.ParamSearch true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[int64] 返回历史数量
// @Router /history/count [post]
func noopHistoryCount() {}

// @Summary 查询历史
// @Schemes
// @Description 查询历史
// @Tags history
// @Param search body curd.ParamSearch true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyList[model.History] 返回历史信息
// @Router /history/search [post]
func noopHistorySearch() {}

// @Summary 查询历史
// @Schemes
// @Description 查询历史
// @Tags history
// @Param search query curd.ParamList true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyList[model.History] 返回历史信息
// @Router /history/list [get]
func noopHistoryList() {}

// @Summary 删除历史
// @Schemes
// @Description 删除历史
// @Tags history
// @Param id path int true "历史ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[model.History] 返回历史信息
// @Router /history/{id}/delete [get]
func noopHistoryDelete() {}

func historyExport(ctx *gin.Context) {

}

func historyRouter(app *gin.RouterGroup) {

	app.POST("/count", curd2.ApiCount[model.History]())

	app.POST("/search", curd2.ApiSearchWith[model.HistoryEx]([]*curd2.Join{{
		Table:        "device",
		LocaleField:  "device_id",
		ForeignField: "id",
		Field:        "name",
		As:           "device",
	}}))

	app.GET("/list", curd2.ApiList[model.History]())

	app.GET("/:id/delete", curd2.ParseParamId, curd2.ApiDelete[model.History]())

	app.POST("/export", historyExport)
}
