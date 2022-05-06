package api

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/database"
	"github.com/zgwit/iot-master/model"
	"github.com/zgwit/storm/v3"
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
		replyError(ctx, err)
		return
	}

	var user model.User
	err := database.Master.One("Username", obj.Username, &user)
	if err == storm.ErrNotFound {
		//管理员自动创建
		if obj.Username == "admin" {
			user.Username = obj.Username
			user.Nickname = "管理员"
			err = database.Master.Save(&user)
		} else {
			replyFail(ctx, "找不到用户")
			return
		}

	}
	if err != nil {
		replyError(ctx, err)
		return
	}

	if user.Disabled {
		replyFail(ctx, "用户已禁用")
		return
	}

	var password model.Password
	err = database.Master.One("Id", user.Id, &password)
	//初始化密码
	if err == storm.ErrNotFound {
		password.Id = user.Id
		password.Password = md5hash("123456")
		err = database.Master.Save(&password)
	}
	if err != nil {
		replyError(ctx, err)
		return
	}

	if password.Password != obj.Password {
		replyFail(ctx, "密钥错误")
		return
	}

	_ = database.History.Save(model.Event{UserId: user.Id, Event: "登录"})

	//存入session
	session.Set("user", &user)
	_ = session.Save()

	replyOk(ctx, user)
}

func logout(ctx *gin.Context) {
	session := sessions.Default(ctx)
	u := session.Get("user")
	if u == nil {
		replyFail(ctx, "未登录")
		return
	}

	user := u.(*model.User)
	_ = database.History.Save(model.Event{UserId: user.Id, Event: "登录"})

	session.Clear()
	_ = session.Save()
	replyOk(ctx, nil)
}
