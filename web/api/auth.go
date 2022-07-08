package api

import (
	"github.com/gin-gonic/gin"
	"iot-master/config"
	"iot-master/db"
	"iot-master/model"
)

type authObj struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func auth(ctx *gin.Context) {
	username := ctx.Query("username")
	password := ctx.Query("password")

	var user model.User
	has, err := db.Engine.Where("username=?", username).Get(&user)
	if err != nil {
		replyError(ctx, err)
		return
	}

	if !has {
		replyFail(ctx, "找不到用户")
		return
	}

	if user.Disabled {
		replyFail(ctx, "用户已禁用")
		return
	}

	var obj model.Password
	has, err = db.Engine.ID(user.Id).Get(&obj)
	if err != nil {
		replyError(ctx, err)
		return
	}

	//初始化密码
	if !has {
		dp := config.Config.DefaultPassword
		if dp == "" {
			dp = "123456"
		}
		obj.Password = dp
	}

	if obj.Password != password {
		replyFail(ctx, "密码错误")
		return
	}

	//TODO 生成Token
	replyOk(ctx, user.Id)
}
