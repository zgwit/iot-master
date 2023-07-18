package api

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/v4/model"
	"github.com/zgwit/iot-master/v4/pkg/curd"
	"github.com/zgwit/iot-master/v4/pkg/db"
)

type loginObj struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Remember bool   `json:"remember"`
}

func md5hash(text string) string {
	h := md5.New()
	h.Write([]byte(text))
	sum := h.Sum(nil)
	return hex.EncodeToString(sum)
}

func login(ctx *gin.Context) {
	session := sessions.Default(ctx)

	var obj loginObj
	if err := ctx.ShouldBindJSON(&obj); err != nil {
		curd.Error(ctx, err)
		return
	}

	var user model.User
	has, err := db.Engine.Where("username=?", obj.Username).Get(&user)
	if err != nil {
		curd.Error(ctx, err)
		return
	}

	if !has {
		//管理员自动创建
		if obj.Username == "admin" {
			user.Id = "admin"
			user.Username = obj.Username
			user.Name = "管理员"
			_, err = db.Engine.InsertOne(&user)
			if err != nil {
				curd.Error(ctx, err)
				return
			}
		} else {
			curd.Fail(ctx, "找不到用户")
			return
		}
	}

	if user.Disabled {
		curd.Fail(ctx, "用户已禁用")
		return
	}

	var password model.Password
	has, err = db.Engine.ID(user.Id).Get(&password)
	if err != nil {
		curd.Error(ctx, err)
		return
	}

	//初始化密码
	if !has {
		dp := "123456"

		password.Id = user.Id
		password.Password = md5hash(dp)
		_, err = db.Engine.InsertOne(&password)
		if err != nil {
			curd.Error(ctx, err)
			return
		}
	}

	if password.Password != obj.Password {
		curd.Fail(ctx, "密码错误")
		return
	}

	//_, _ = db.Engine.InsertOne(&model.UserEvent{UserId: user.Id, ModEvent: model.ModEvent{Type: "登录"}})

	//存入session
	session.Set("user", user.Id)
	_ = session.Save()

	curd.OK(ctx, user)
}

func logout(ctx *gin.Context) {
	session := sessions.Default(ctx)
	u := session.Get("user")
	if u == nil {
		curd.Fail(ctx, "未登录")
		return
	}

	//user := u.(int64)
	//_, _ = db.Engine.InsertOne(&model.UserEvent{UserId: user, ModEvent: model.ModEvent{Type: "退出"}})

	session.Clear()
	_ = session.Save()
	curd.OK(ctx, nil)
}

type passwordObj struct {
	Old string `json:"old"`
	New string `json:"new"`
}

func password(ctx *gin.Context) {

	var obj passwordObj
	if err := ctx.ShouldBindJSON(&obj); err != nil {
		curd.Error(ctx, err)
		return
	}

	var pwd model.Password
	has, err := db.Engine.ID(ctx.GetString("user")).Get(&pwd)
	if err != nil {
		curd.Error(ctx, err)
		return
	}
	if !has {
		curd.Fail(ctx, "用户不存在")
		return
	}
	if obj.Old != pwd.Password {
		curd.Fail(ctx, "密码错误")
		return
	}

	pwd.Password = obj.New //前端已经加密过
	_, err = db.Engine.ID(ctx.GetString("user")).Cols("password").Update(&pwd)
	if err != nil {
		curd.Error(ctx, err)
		return
	}

	curd.OK(ctx, nil)
}
