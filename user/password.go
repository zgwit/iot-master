package user

import (
	"github.com/gin-gonic/gin"
	"github.com/god-jason/bucket/api"
	"go.mongodb.org/mongo-driver/bson"
)

type passwordObj struct {
	Old string `json:"old"`
	New string `json:"new"`
}

func password(ctx *gin.Context) {

	var obj passwordObj
	if err := ctx.ShouldBind(&obj); err != nil {
		api.Error(ctx, err)
		return
	}

	userId := ctx.GetString("user")

	var pwd Password
	has, err := _passwordTable.Get(userId, &pwd)
	if err != nil {
		api.Error(ctx, err)
		return
	}

	if !has {
		_, err = _passwordTable.Insert(map[string]any{
			"_id":      userId,
			"password": obj.New,
		})
		if err != nil {
			api.Error(ctx, err)
			return
		}
	} else {
		if obj.Old != pwd.Password {
			api.Fail(ctx, "密码错误")
			return
		}

		//pwd.Password = obj.New //前端已经加密过
		err = _passwordTable.Update(userId, bson.M{"password": obj.New})
		if err != nil {
			api.Error(ctx, err)
			return
		}
	}

	api.OK(ctx, nil)
}
