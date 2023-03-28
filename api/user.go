package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/v3/model"
	"github.com/zgwit/iot-master/v3/pkg/curd"
	"github.com/zgwit/iot-master/v3/pkg/db"
)

// @Summary 获取用户信息
// @Schemes
// @Description 这里写描述 get users
// @Tags user
// @Produce json
// @Success 200 {object} curd.ReplyData[model.User] 返回用户信息
// @Router /user/me [get]
func noopUserMe() {}

// @Summary 查询用户数量
// @Schemes
// @Description 查询用户数量
// @Tags user
// @Param search body curd.ParamSearch true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[int64] 返回用户数量
// @Router /user/count [post]
func noopUserCount() {}

// @Summary 查询用户
// @Schemes
// @Description 查询用户
// @Tags user
// @Param search body curd.ParamSearch true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyList[model.User] 返回用户信息
// @Router /user/search [post]
func noopUserSearch() {}

// @Summary 查询用户
// @Schemes
// @Description 查询用户
// @Tags user
// @Param search query curd.ParamList true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyList[model.User] 返回用户信息
// @Router /user/list [get]
func noopUserList() {}

// @Summary 创建用户
// @Schemes
// @Description 创建用户
// @Tags user
// @Param search body model.User true "用户信息"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[model.User] 返回用户信息
// @Router /user/create [post]
func noopUserCreate() {}

// @Summary 修改用户
// @Schemes
// @Description 修改用户
// @Tags user
// @Param id path int true "用户ID"
// @Param user body model.User true "用户信息"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[model.User] 返回用户信息
// @Router /user/{id} [post]
func noopUserUpdate() {}

// @Summary 获取用户
// @Schemes
// @Description 获取用户
// @Tags user
// @Param id path int true "用户ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[model.User] 返回用户信息
// @Router /user/{id} [get]
func noopUserGet() {}

// @Summary 删除用户
// @Schemes
// @Description 删除用户
// @Tags user
// @Param id path int true "用户ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[model.User] 返回用户信息
// @Router /user/{id}/delete [get]
func noopUserDelete() {}

// @Summary 修改密码
// @Schemes
// @Description 修改密码
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[model.User] 返回用户信息
// @Router /user/password [get]
func noopUserPassword() {}

// @Summary 启用用户
// @Schemes
// @Description 启用用户
// @Tags user
// @Param id path int true "用户ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[model.User] 返回用户信息
// @Router /user/{id}/enable [get]
func noopUserEnable() {}

// @Summary 禁用用户
// @Schemes
// @Description 禁用用户
// @Tags user
// @Param id path int true "用户ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[model.User] 返回用户信息
// @Router /user/{id}/disable [get]
func noopUserDisable() {}

func userRouter(app *gin.RouterGroup) {

	app.GET("/me", userMe)

	app.POST("/search", curd.ApiSearch[model.User]())

	app.GET("/list", curd.ApiList[model.User]())

	app.POST("/create", curd.ParseParamId, curd.ApiCreate[model.User](nil, nil))

	app.GET("/:id", curd.ParseParamId, curd.ApiGet[model.User]())

	app.POST("/:id", curd.ParseParamId, curd.ApiModify[model.User](nil, nil,
		"username", "name", "email", "roles", "disabled"))

	app.GET("/:id/delete", curd.ParseParamId, curd.ApiDelete[model.User](nil, nil))

	app.GET("/:id/password", curd.ParseParamId, userPassword)

	app.GET("/:id/enable", curd.ParseParamId, curd.ApiDisable[model.User](false, nil, nil))

	app.GET("/:id/disable", curd.ParseParamId, curd.ApiDisable[model.User](true, nil, nil))

}

func userMe(ctx *gin.Context) {
	id := ctx.GetInt64("user")
	var user model.User
	has, err := db.Engine.ID(id).Get(&user)
	if err != nil {
		curd.Error(ctx, err)
		return
	}
	if !has {
		curd.Fail(ctx, "用户不存在")
		return
	}
	curd.OK(ctx, &user)
}

func userPassword(ctx *gin.Context) {
	var p model.Password
	p.Id = ctx.GetInt64("id")
	pwd := ctx.PostForm("password")

	//p.Password = md5hash(pwd)
	p.Password = pwd //前端已经加密过

	_, err := db.Engine.Cols("password").Update(&p)
	if err != nil {
		curd.Error(ctx, err)
		return
	}

	curd.OK(ctx, nil)
}
