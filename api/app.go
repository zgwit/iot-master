package api

import (
	"github.com/gin-gonic/gin"
	curd2 "github.com/zgwit/iot-master/v4/curd"
	"github.com/zgwit/iot-master/v4/model"
)

func appRouter(app *gin.RouterGroup) {

	app.POST("/search", curd2.ApiSearch[model.App]())

	app.GET("/list", curd2.ApiList[model.App]())

	app.POST("/create", curd2.ApiCreateHook[model.App](curd2.GenerateRandomId[model.App](8), nil))

	app.GET("/:id", curd2.ParseParamStringId, curd2.ApiGet[model.App]())

	app.POST("/:id", curd2.ParseParamStringId, curd2.ApiUpdateHook[model.App](nil,
		nil, "id", "name", "type", "address", "desc", "disabled"))

	app.GET("/:id/delete", curd2.ParseParamStringId, curd2.ApiDeleteHook[model.App](nil, nil))
}
