package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/config"
	"github.com/zgwit/iot-master/db"
	"github.com/zgwit/iot-master/history"
)

type installDatabaseBody struct {
	Url string
}

func installDatabase(ctx *gin.Context) {
	var cfg config.Database
	err := ctx.BindJSON(&cfg)
	if err != nil {
		replyError(ctx, err)
		return
	}
	err = db.Open(&cfg)
	if err != nil {
		replyError(ctx, err)
		return
	}
	replyOk(ctx, nil)
}

func installHistory(ctx *gin.Context) {
	var cfg config.History
	err := ctx.BindJSON(&cfg)
	if err != nil {
		replyError(ctx, err)
		return
	}
	err = history.Open(&cfg)
	if err != nil {
		replyError(ctx, err)
		return
	}
	replyOk(ctx, nil)

}

func installSystem(ctx *gin.Context) {
	//var cfg config.Configure
	err := ctx.BindJSON(&config.Config)
	if err != nil {
		replyError(ctx, err)
		return
	}

	err = config.Store()
	if err != nil {
		replyError(ctx, err)
		return
	}

	replyOk(ctx, nil)
}
