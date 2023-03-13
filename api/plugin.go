package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/v3/model"
	"github.com/zgwit/iot-master/v3/pkg/curd"
)

func pluginRouter(app *gin.RouterGroup) {

	app.POST("/search", curd.ApiSearch[model.Plugin]())
	app.GET("/list", curd.ApiList[model.Plugin]())
	app.POST("/create", curd.ApiCreate[model.Plugin](curd.GenerateUuidKey, nil))
	app.GET("/:id", curd.ParseParamStringId, curd.ApiGet[model.Plugin]())
	app.POST("/:id", curd.ParseParamStringId, curd.ApiModify[model.Plugin](nil, nil,
		"id", "name", "version", "command", "dependencies"))
	app.GET("/:id/delete", curd.ParseParamStringId, curd.ApiDelete[model.Plugin](nil, nil))
}
