package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/v3/model"
	"github.com/zgwit/iot-master/v3/pkg/curd"
)

// @Summary 查询分组
// @Schemes
// @Description 查询分组
// @Tags group
// @Param search body curd.ParamSearch true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyList[model.Group] 返回分组信息
// @Router /group/search [post]
func noopGroupSearch() {}

// @Summary 查询分组
// @Schemes
// @Description 查询分组
// @Tags group
// @Param search query curd.ParamList true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyList[model.Group] 返回分组信息
// @Router /group/list [get]
func noopGroupList() {}

// @Summary 创建分组
// @Schemes
// @Description 创建分组
// @Tags group
// @Param search body model.Group true "分组信息"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[model.Group] 返回分组信息
// @Router /group/create [post]
func noopGroupCreate() {}

// @Summary 修改分组
// @Schemes
// @Description 修改分组
// @Tags group
// @Param id path int true "分组ID"
// @Param group body model.Group true "分组信息"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[model.Group] 返回分组信息
// @Router /group/{id} [post]
func noopGroupUpdate() {}

// @Summary 获取分组
// @Schemes
// @Description 获取分组
// @Tags group
// @Param id path int true "分组ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[model.Group] 返回分组信息
// @Router /group/{id} [get]
func noopGroupGet() {}

// @Summary 删除分组
// @Schemes
// @Description 删除分组
// @Tags group
// @Param id path int true "分组ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[model.Group] 返回分组信息
// @Router /group/{id}/delete [get]
func noopGroupDelete() {}

func groupRouter(app *gin.RouterGroup) {

	app.POST("/search", curd.ApiSearch[model.Group]())
	app.GET("/list", curd.ApiList[model.Group]())
	app.POST("/create", curd.ApiCreate[model.Group](nil, nil))
	app.GET("/:id", curd.ParseParamId, curd.ApiGet[model.Group]())
	app.POST("/:id", curd.ParseParamId, curd.ApiModify[model.Group](nil, nil,
		"name", "desc"))
	app.GET("/:id/delete", curd.ParseParamId, curd.ApiDelete[model.Group](nil, nil))
}
