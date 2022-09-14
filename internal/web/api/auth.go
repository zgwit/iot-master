package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/internal/config"
	"github.com/zgwit/iot-master/internal/db"
	"github.com/zgwit/iot-master/model"
	"github.com/zgwit/iot-master/pkg/lib"
)

var tokens = lib.ExpireCache{Timeout: 60 * 60 * 2} //2h 可以改成配置

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
		obj.Password = md5hash(dp)
	}

	if obj.Password != password {
		replyFail(ctx, "密码错误")
		return
	}

	//生成Token
	token := lib.RandomString(12)

	//保存用户
	tokens.Store(token, &user)

	replyOk(ctx, token)
}
