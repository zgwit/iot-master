package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/v2/internal/config"
	"github.com/zgwit/iot-master/v2/internal/db"
	"github.com/zgwit/iot-master/v2/model"
)

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
		obj.Password = md5hash(dp)
	}

	if obj.Password != password {
		replyFail(ctx, "密码错误")
		return
	}

	//生成Token
	token, err := jwtGenerate(user.Id)
	if err != nil {
		return
	}

	replyOk(ctx, gin.H{
		token: token,
	})
}
