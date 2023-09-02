package api

import (
	"github.com/gin-gonic/gin"
	curd2 "github.com/zgwit/iot-master/v4/curd"
	"github.com/zgwit/iot-master/v4/db"
	"github.com/zgwit/iot-master/v4/model"
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
// @Param id path string true "用户ID"
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
// @Param id path string true "用户ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[model.User] 返回用户信息
// @Router /user/{id} [get]
func noopUserGet() {}

// @Summary 删除用户
// @Schemes
// @Description 删除用户
// @Tags user
// @Param id path string true "用户ID"
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
// @Param id path string true "用户ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[model.User] 返回用户信息
// @Router /user/{id}/enable [get]
func noopUserEnable() {}

// @Summary 禁用用户
// @Schemes
// @Description 禁用用户
// @Tags user
// @Param id path string true "用户ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[model.User] 返回用户信息
// @Router /user/{id}/disable [get]
func noopUserDisable() {}

func userRouter(app *gin.RouterGroup) {

	app.GET("/me", userMe)

	app.POST("/search", curd2.ApiSearch[model.User]())

	app.GET("/list", curd2.ApiList[model.User]())

	app.POST("/create", curd2.ParseParamStringId, curd2.ApiCreateHook[model.User](curd2.GenerateRandomId[model.User](6), nil))

	app.GET("/:id", curd2.ParseParamStringId, curd2.ApiGet[model.User]())

	app.POST("/:id", curd2.ParseParamStringId, curd2.ApiUpdateHook[model.User](nil, nil,
		"username", "name", "cellphone", "email", "roles", "disabled"))

	app.GET("/:id/delete", curd2.ParseParamStringId, curd2.ApiDeleteHook[model.User](nil, nil))

	app.GET("/:id/password", curd2.ParseParamStringId, userPassword)

	app.GET("/:id/enable", curd2.ParseParamStringId, curd2.ApiDisableHook[model.User](false, nil, nil))

	app.GET("/:id/disable", curd2.ParseParamStringId, curd2.ApiDisableHook[model.User](true, nil, nil))

	app.GET("/export", curd2.ApiExport("user", "用户"))

	app.POST("/import", curd2.ApiImport("user"))

}

func userMe(ctx *gin.Context) {
	id := ctx.GetString("user")
	var user model.User
	has, err := db.Engine.ID(id).Get(&user)
	if err != nil {
		curd2.Error(ctx, err)
		return
	}
	if !has {
		curd2.Fail(ctx, "用户不存在")
		return
	}
	curd2.OK(ctx, &user)
}

func userPassword(ctx *gin.Context) {
	var p model.Password
	p.Id = ctx.GetString("id")
	pwd := ctx.PostForm("password")

	//p.Password = md5hash(pwd)
	p.Password = pwd //前端已经加密过

	_, err := db.Engine.Cols("password").Update(&p)
	if err != nil {
		curd2.Error(ctx, err)
		return
	}

	curd2.OK(ctx, nil)
}
