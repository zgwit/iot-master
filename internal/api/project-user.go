package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/v4/pkg/web/curd"
	"github.com/zgwit/iot-master/v4/types"
)

// @Summary 查询项目用户数量
// @Schemes
// @Description 查询项目用户数量
// @Tags project-user
// @Param search body curd.ParamSearch true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[int64] 返回项目用户数量
// @Router /project/user/count [post]
func noopProjectUserCount() {}

// @Summary 查询项目用户
// @Schemes
// @Description 这里写描述 get project-users
// @Tags project-user
// @Param search body curd.ParamSearch true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyList[types.ProjectUser] 返回项目用户信息
// @Router /project/user/search [post]
func noopProjectUserSearch() {}

// @Summary 查询项目用户
// @Schemes
// @Description 查询项目用户
// @Tags project-user
// @Param search query curd.ParamList true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyList[types.ProjectUser] 返回项目用户信息
// @Router /project/user/list [get]
func noopProjectUserList() {}

// @Summary 创建项目用户
// @Schemes
// @Description 创建项目用户
// @Tags project-user
// @Param search body types.ProjectUser true "项目用户信息"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[types.ProjectUser] 返回项目用户信息
// @Router /project/user/create [post]
func noopProjectUserCreate() {}

// @Summary 修改项目用户
// @Schemes
// @Description 修改项目用户
// @Tags project-user
// @Param id path int true "项目用户ID"
// @Param project-user body types.ProjectUser true "项目用户信息"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[types.ProjectUser] 返回项目用户信息
// @Router /project/user/{id} [post]
func noopProjectUserUpdate() {}

// @Summary 获取项目用户
// @Schemes
// @Description 获取项目用户
// @Tags project-user
// @Param id path int true "项目用户ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[types.ProjectUser] 返回项目用户信息
// @Router /project/user/{id} [get]
func noopProjectUserGet() {}

// @Summary 删除项目用户
// @Schemes
// @Description 删除项目用户
// @Tags project-user
// @Param id path int true "项目用户ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[types.ProjectUser] 返回项目用户信息
// @Router /project/user/{id}/delete [get]
func noopProjectUserDelete() {}

func projectUserRouter(app *gin.RouterGroup) {

	app.POST("/count", curd.ApiCount[types.ProjectUser]())

	app.POST("/search", curd.ApiSearch[types.ProjectUser]())

	app.GET("/list", curd.ApiList[types.ProjectUser]())

	app.POST("/create", curd.ApiCreate[types.ProjectUser]())

	app.GET("/:id", curd.ParseParamId, curd.ApiGet[types.ProjectUser]())

	app.POST("/:id", curd.ParseParamId, curd.ApiUpdateHook[types.ProjectUser](nil, nil,
		"id", "project_id", "user_id", "admin", "disabled"))

	app.GET("/:id/delete", curd.ParseParamId, curd.ApiDeleteHook[types.ProjectUser](nil, nil))

}
