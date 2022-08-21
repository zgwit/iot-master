package api

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/timshannon/bolthold"
	"github.com/zgwit/iot-master/internal/config"
	"github.com/zgwit/iot-master/internal/db"
	"github.com/zgwit/iot-master/model"
	"time"
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
	err := db.Store().FindOne(&user, bolthold.Where("Username").Eq(obj.Username))

	if err == bolthold.ErrNotFound {
		//管理员自动创建
		if obj.Username == "admin" {
			user.Username = obj.Username
			user.Nickname = "管理员"
			user.Created = time.Now()
			err = db.Store().Insert(bolthold.NextSequence(), &user)
			if err != nil {
				replyError(ctx, err)
				return
			}
		} else {
			replyFail(ctx, "找不到用户")
			return
		}
	} else if err != nil {
		replyError(ctx, err)
		return
	}

	if user.Disabled {
		replyFail(ctx, "用户已禁用")
		return
	}

	var pwd model.Password
	err = db.Store().Get(user.Id, &pwd)
	if err == bolthold.ErrNotFound {
		//初始化密码
		dp := config.Config.DefaultPassword
		if dp == "" {
			dp = "123456"
		}
		pwd.Password = md5hash(dp)
	} else if err != nil {
		replyError(ctx, err)
		return
	}

	if pwd.Password != obj.Password {
		replyFail(ctx, "密码错误")
		return
	}

	//存入session
	session.Set("user", user.Id)
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

	//user := u.(uint64)
	session.Clear()
	_ = session.Save()
	replyOk(ctx, nil)
}

type passwordObj struct {
	Old string `json:"old"`
	New string `json:"new"`
}

func password(ctx *gin.Context) {

	var obj passwordObj
	if err := ctx.ShouldBindJSON(&obj); err != nil {
		replyError(ctx, err)
		return
	}

	var pwd model.Password
	err := db.Store().Get(ctx.GetUint64("user"), &pwd)
	if err != nil {
		replyError(ctx, err)
		return
	}
	if obj.Old != pwd.Password {
		replyFail(ctx, "密码错误")
		return
	}

	pwd.Password = obj.New //前端已经加密过
	err = db.Store().Update(pwd.Id, &pwd)
	if err != nil {
		replyError(ctx, err)
		return
	}

	replyOk(ctx, nil)
}
