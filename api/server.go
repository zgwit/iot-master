package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/v3/model"
	"github.com/zgwit/iot-master/v3/pkg/curd"
)

// @Summary 查询服务器
// @Schemes
// @Description 查询服务器
// @Tags server
// @Param search body curd.ParamSearch true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyList[model.Server] 返回服务器信息
// @Router /server/search [post]
func noopServerSearch() {}

// @Summary 查询服务器
// @Schemes
// @Description 查询服务器
// @Tags server
// @Param search query curd.ParamList true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyList[model.Server] 返回服务器信息
// @Router /server/list [get]
func noopServerList() {}

// @Summary 创建服务器
// @Schemes
// @Description 创建服务器
// @Tags server
// @Param search body model.Server true "服务器信息"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[model.Server] 返回服务器信息
// @Router /server/create [post]
func noopServerCreate() {}

// @Summary 修改服务器
// @Schemes
// @Description 修改服务器
// @Tags server
// @Param id path int true "服务器ID"
// @Param server body model.Server true "服务器信息"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[model.Server] 返回服务器信息
// @Router /server/{id} [post]
func noopServerUpdate() {}

// @Summary 删除服务器
// @Schemes
// @Description 删除服务器
// @Tags server
// @Param id path int true "服务器ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[model.Server] 返回服务器信息
// @Router /server/{id}/delete [get]
func noopServerDelete() {}

// @Summary 启用服务器
// @Schemes
// @Description 启用服务器
// @Tags server
// @Param id path int true "服务器ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[model.Server] 返回服务器信息
// @Router /server/{id}/enable [get]
func noopServerEnable() {}

// @Summary 禁用服务器
// @Schemes
// @Description 禁用服务器
// @Tags server
// @Param id path int true "服务器ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[model.Server] 返回服务器信息
// @Router /server/{id}/disable [get]
func noopServerDisable() {}

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
