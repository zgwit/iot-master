package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/v3/model"
	"github.com/zgwit/iot-master/v3/pkg/curd"
	"github.com/zgwit/iot-master/v3/pkg/db"
)

func userRouter(app *gin.RouterGroup) {

	app.GET("/me", userMe)

	app.POST("/search", curd.ApiSearch[model.User]())

	app.GET("/list", curd.ApiList[model.User]())

	app.POST("/create", curd.ParseParamId, curd.ApiCreate[model.User](nil, nil))

	app.GET("/:id", curd.ParseParamId, curd.ApiGet[model.User]())

	app.POST("/:id", curd.ParseParamId, curd.ApiModify[model.User](nil, nil,
		"username", "name", "email", "disabled"))

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
