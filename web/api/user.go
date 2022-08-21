package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/internal/db"
	"github.com/zgwit/iot-master/model"
)

func userMe(ctx *gin.Context) {
	id := ctx.GetUint64("user")
	var user model.User
	err := db.Store().Get(id, &user)
	if err != nil {
		replyError(ctx, err)
		return
	}
	replyOk(ctx, &user)
}

func userPassword(ctx *gin.Context) {
	var p model.Password
	p.Id = ctx.GetUint64("id")
	pwd := ctx.PostForm("password")

	//p.Password = md5hash(pwd)
	p.Password = pwd //前端已经加密过

	err := db.Store().Update(p.Id, &pwd)
	if err != nil {
		replyError(ctx, err)
		return
	}

	replyOk(ctx, nil)
}
