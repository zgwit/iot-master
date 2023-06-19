package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/v3/model"
	"github.com/zgwit/iot-master/v3/pkg/curd"
)

// @Summary 查询通知
// @Schemes
// @Description 查询通知
// @Tags notification
// @Param search body curd.ParamSearch true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[int64] 返回通知信息
// @Router /notification/count [post]
func noopNotificationCount() {}

// @Summary 查询通知
// @Schemes
// @Description 查询通知
// @Tags notification
// @Param search body curd.ParamSearch true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyList[model.Notification] 返回通知信息
// @Router /notification/search [post]
func noopNotificationSearch() {}

// @Summary 查询通知
// @Schemes
// @Description 查询通知
// @Tags notification
// @Param search query curd.ParamList true "查询参数"
// @Produce json
// @Success 200 {object} curd.ReplyList[model.Notification] 返回通知信息
// @Router /notification/list [get]
func noopNotificationList() {}

// @Summary 删除通知
// @Schemes
// @Description 删除通知
// @Tags notification
// @Param id path int true "通知ID"
// @Produce json
// @Success 200 {object} curd.ReplyData[model.Notification] 返回通知信息
// @Router /notification/{id}/delete [get]
func noopNotificationDelete() {}

func notificationRouter(app *gin.RouterGroup) {

	app.POST("/count", curd.ApiCount[model.Notification]())

	app.POST("/search", curd.ApiSearchWith[model.Notification](
		"notification", []*curd.Join{{
			Table:        "device",
			LocaleField:  "device_id",
			ForeignField: "id",
			Field:        "name",
			As:           "device",
		}},
	))

	app.GET("/list", curd.ApiList[model.Notification]())

	app.GET("/:id", curd.ParseParamId, curd.ApiGet[model.Notification]())

	app.GET("/:id/delete", curd.ParseParamId, curd.ApiDelete[model.Notification]())
}
