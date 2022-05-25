package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/database"
	"github.com/zgwit/iot-master/model"
)

func eventRoutes(app *gin.RouterGroup) {
	app.POST("list", eventList)
	app.GET("clear", eventClear)

	app.Use(parseParamId)
	app.GET(":id/delete", eventDelete)
}

func eventList(ctx *gin.Context) {
	records, cnt, err := normalSearch(ctx, database.History, &model.Event{})
	if err != nil {
		replyError(ctx, err)
		return
	}

	//补充信息
	events := records.(*[]*model.Event)

	replyList(ctx, events, cnt)
}


func eventDelete(ctx *gin.Context) {
	event := model.Event{Id: ctx.GetInt64("id")}
	err := database.History.DeleteStruct(&event)
	if err != nil {
		replyError(ctx, err)
		return
	}

	replyOk(ctx, event)
}

func eventClear(ctx *gin.Context) {
	err := database.History.Drop(&model.Event{})
	if err != nil {
		replyError(ctx, err)
		return
	}

	replyOk(ctx, nil)
}
