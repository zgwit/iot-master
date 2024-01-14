package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/v4/db"
	"github.com/zgwit/iot-master/v4/types"
	"github.com/zgwit/iot-master/v4/web"
	"github.com/zgwit/iot-master/v4/web/curd"
)

func auth(ctx *gin.Context) {
	username := ctx.Query("username")
	password := ctx.Query("password")

	var user types.User
	has, err := db.Engine.Where("id=?", username).Get(&user)
	if err != nil {
		curd.Error(ctx, err)
		return
	}

	if !has {
		curd.Fail(ctx, "找不到用户")
		return
	}

	if user.Disabled {
		curd.Fail(ctx, "用户已禁用")
		return
	}

	var obj types.Password
	has, err = db.Engine.ID(user.Id).Get(&obj)
	if err != nil {
		curd.Error(ctx, err)
		return
	}

	//初始化密码
	if !has {
		dp := "123456"
		obj.Password = md5hash(dp)
	}

	if obj.Password != password {
		curd.Fail(ctx, "密码错误")
		return
	}

	//生成Token
	token, err := web.JwtGenerate(user.Id)
	if err != nil {
		return
	}

	curd.OK(ctx, gin.H{
		token: token,
	})
}
