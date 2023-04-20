package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/v3/model"
	"github.com/zgwit/iot-master/v3/pkg/curd"
)

func appRouter(app *gin.RouterGroup) {

	app.POST("/search", curd.ApiSearch[model.App]())

	app.GET("/list", curd.ApiList[model.App]())

	app.POST("/create", curd.ApiCreate[model.App](curd.GenerateRandomId[model.App](8), nil))

	app.GET("/:id", curd.ParseParamStringId, curd.ApiGet[model.App]())

	app.POST("/:id", curd.ParseParamStringId, curd.ApiModify[model.App](nil,
		nil, "id", "name", "type", "address", "desc", "disabled"))

	app.GET("/:id/delete", curd.ParseParamStringId, curd.ApiDelete[model.App](nil, nil))
}
