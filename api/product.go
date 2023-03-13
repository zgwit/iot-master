package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/v3/model"
	"github.com/zgwit/iot-master/v3/pkg/curd"
)

func productRouter(app *gin.RouterGroup) {

	app.POST("/search", curd.ApiSearch[model.Product]())
	app.GET("/list", curd.ApiList[model.Product]())
	app.POST("/create", curd.ApiCreate[model.Product](curd.GenerateRandomKey(8), nil))
	app.GET("/:id", curd.ParseParamStringId, curd.ApiGet[model.Product]())
	app.POST("/:id", curd.ParseParamStringId, curd.ApiModify[model.Product](nil, nil,
		"id", "name", "version", "desc", "properties", "functions", "events", "parameters", "constraints"))
	app.GET("/:id/delete", curd.ParseParamStringId, curd.ApiDelete[model.Product](nil, nil))
}
