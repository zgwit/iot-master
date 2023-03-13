package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/v3/model"
	"github.com/zgwit/iot-master/v3/pkg/curd"
)

func serverRouter(app *gin.RouterGroup) {

	app.POST("/search", curd.ApiSearch[model.Server]())
	app.GET("/list", curd.ApiList[model.Server]())
	app.POST("/create", curd.ApiCreate[model.Server](nil, nil))
	app.GET("/:id", curd.ParseParamId, curd.ApiGet[model.Server]())
	app.POST("/:id", curd.ParseParamId, curd.ApiModify[model.Server](nil, nil,
		"name", "type", "port", "desc", "disabled"))
	app.GET("/:id/delete", curd.ParseParamId, curd.ApiDelete[model.Server](nil, nil))
}

func afterServerCreate(data interface{}) error {
	//server := data.(*model.Server)

	//TODO start server
	return nil
}

func afterServerUpdate(data interface{}) error {
	//server := data.(*model.Server)

	//TODO restart server
	return nil
}

func afterServerDelete(id interface{}) error {
	//gid := id.(string)

	//todo stop server
	return nil
}
