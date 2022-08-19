package api

import (
	"github.com/gin-gonic/gin"
	"github.com/timshannon/bolthold"
	"iot-master/internal/config"
	"iot-master/internal/db"
	"iot-master/model"
	lib2 "iot-master/pkg/lib"
)

var tokens = lib2.ExpireCache{Timeout: 60 * 60 * 2} //2h 可以改成配置

type authObj struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func auth(ctx *gin.Context) {
	username := ctx.Query("username")
	password := ctx.Query("password")

	var user model.User
	err := db.Store().FindOne(&user, bolthold.Where("Username").Eq(username))

	if err == bolthold.ErrNotFound {
		replyFail(ctx, "找不到用户")
		return
	} else if err != nil {
		replyError(ctx, err)
		return
	}

	if user.Disabled {
		replyFail(ctx, "用户已禁用")
		return
	}

	var obj model.Password
	err = db.Store().Get(user.Id, &obj)
	if err == bolthold.ErrNotFound {
		//初始化密码
		dp := config.Config.DefaultPassword
		if dp == "" {
			dp = "123456"
		}
		obj.Password = md5hash(dp)
	} else if err != nil {
		replyError(ctx, err)
		return
	}

	if obj.Password != password {
		replyFail(ctx, "密码错误")
		return
	}

	//生成Token
	token := lib2.RandomString(12)

	//保存用户
	tokens.Store(token, &user)

	replyOk(ctx, token)
}
