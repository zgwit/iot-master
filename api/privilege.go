package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/v3/model"
	"github.com/zgwit/iot-master/v3/pkg/curd"
)

// @Summary 查询权限
// @Schemes
// @Description 这里写描述 get privileges
// @Tags privilege
// @Param search body curd.ParamSearch true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyList[model.Privilege] 返回权限信息
// @Router /privilege/search [post]
func noopPrivilegeSearch() {}

// @Summary 查询权限
// @Schemes
// @Description 查询权限
// @Tags privilege
// @Param search query curd.ParamList true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyList[model.Privilege] 返回权限信息
// @Router /privilege/list [get]
func noopPrivilegeList() {}

// @Summary 创建权限
// @Schemes
// @Description 创建权限
// @Tags privilege
// @Param search body model.Privilege true "权限信息"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[model.Privilege] 返回权限信息
// @Router /privilege/create [post]
func noopPrivilegeCreate() {}

// @Summary 修改权限
// @Schemes
// @Description 修改权限
// @Tags privilege
// @Param id path string true "权限ID"
// @Param privilege body model.Privilege true "权限信息"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[model.Privilege] 返回权限信息
// @Router /privilege/{id} [post]
func noopPrivilegeUpdate() {}

// @Summary 删除权限
// @Schemes
// @Description 删除权限
// @Tags privilege
// @Param id path string true "权限ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[model.Privilege] 返回权限信息
// @Router /privilege/{id}/delete [get]
func noopPrivilegeDelete() {}

func privilegeRouter(app *gin.RouterGroup) {
	app.POST("/search", curd.ApiSearch[model.Privilege]())

	app.GET("/list", curd.ApiList[model.Privilege]())

	app.POST("/create", curd.ParseParamStringId, curd.ApiCreate[model.Privilege](nil, nil))

	app.GET("/:id", curd.ParseParamStringId, curd.ApiGet[model.Privilege]())

	app.POST("/:id", curd.ParseParamStringId, curd.ApiModify[model.Privilege](nil, nil, "name"))

	app.GET("/:id/delete", curd.ParseParamStringId, curd.ApiDelete[model.Privilege](nil, nil))

}
