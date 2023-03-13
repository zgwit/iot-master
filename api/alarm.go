package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/v3/model"
	"github.com/zgwit/iot-master/v3/pkg/curd"
	"github.com/zgwit/iot-master/v3/pkg/db"
)

func alarmRouter(app *gin.RouterGroup) {

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
