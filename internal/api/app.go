package api

import (
	"github.com/gin-gonic/gin"
	curd "github.com/zgwit/iot-master/v4/pkg/curd"
	"github.com/zgwit/iot-master/v4/types"
)

func appRouter(app *gin.RouterGroup) {

	app.POST("/search", curd.ApiSearch[types.App]())

	app.GET("/list", curd.ApiList[types.App]())

	app.POST("/create", curd.ApiCreateHook[types.App](curd.GenerateRandomId[types.App](8), nil))

	app.GET("/:id", curd.ParseParamStringId, curd.ApiGet[types.App]())

	app.POST("/:id", curd.ParseParamStringId, curd.ApiUpdateHook[types.App](nil,
		nil, "id", "name", "type", "address", "desc", "disabled"))

	app.GET("/:id/delete", curd.ParseParamStringId, curd.ApiDeleteHook[types.App](nil, nil))
}
