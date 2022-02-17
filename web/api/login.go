package api

import (
	"github.com/asdine/storm/v3"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/database"
	"github.com/zgwit/iot-master/model"
	"net/http"
)

type loginObj struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Remember bool   `json:"remember"`
}

func login(ctx *gin.Context) {
	session := sessions.Default(ctx)

	var obj loginObj
	if err := ctx.ShouldBindJSON(&obj); err != nil {
		replyError(ctx, err)
		return
	}

	var user model.User
	err := database.Master.Find("username",obj.Username, &user)
	if err != nil {
		if err == storm.ErrNotFound {
			replyFail(ctx, "找不到用户")
			return
		}
		replyError(ctx, err)
		return
	}


	if user.Disabled {
		replyFail(ctx, "用户已禁用")
		return
	}

	var password model.Password
	err = database.Master.One("ID", user.ID, &password)
	//初始化密码
	if err == storm.ErrNotFound {
		password.ID = user.ID
		password.Password = "123456" //TODO 加密，可配置，或随机
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

	//存入session
	session.Set("user", user)
	_ = session.Save()

	replyOk(ctx, user)
}

func logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	_ = session.Save()
	c.JSON(http.StatusOK, gin.H{})
}