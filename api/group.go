package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/v3/model"
	"github.com/zgwit/iot-master/v3/pkg/curd"
)

func groupRouter(app *gin.RouterGroup) {

	app.POST("/search", curd.ApiSearch[model.Group]())
	app.GET("/list", curd.ApiList[model.Group]())
	app.POST("/create", curd.ApiCreate[model.Group](nil, nil))
	app.GET("/:id", curd.ParseParamId, curd.ApiGet[model.Group]())
	app.POST("/:id", curd.ParseParamId, curd.ApiModify[model.Group](nil, nil,
		"name", "desc"))
	app.GET("/:id/delete", curd.ParseParamId, curd.ApiDelete[model.Group](nil, nil))
}
