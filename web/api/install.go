package api

import (
	"github.com/gin-gonic/gin"
	"iot-master/config"
	"iot-master/db"
	"iot-master/history"
)

type installBaseObj struct {
	Node string `json:"node"`
	Data string `json:"data"`
	Port string `json:"port"`
}

func installBase(ctx *gin.Context) {
	var cfg installBaseObj
	err := ctx.BindJSON(&cfg)
	if err != nil {
		replyError(ctx, err)
		return
	}

	config.Config.Node = cfg.Node
	config.Config.Data = cfg.Data
	config.Config.Web.Addr = cfg.Port

	replyOk(ctx, nil)
}

func installDatabase(ctx *gin.Context) {
	//var cfg config.Database
	cfg := &config.Config.Database
	err := ctx.BindJSON(cfg)
	if err != nil {
		replyError(ctx, err)
		return
	}
	err = db.Open(cfg)
	if err != nil {
		replyError(ctx, err)
		return
	}

	err = db.Sync()
	if err != nil {
		replyError(ctx, err)
		return
	}

	replyOk(ctx, nil)
}

func installHistory(ctx *gin.Context) {
	//var cfg config.History
	cfg := &config.Config.History
	err := ctx.BindJSON(cfg)
	if err != nil {
		replyError(ctx, err)
		return
	}
	err = history.Open(cfg)
	if err != nil {
		replyError(ctx, err)
		return
	}
	replyOk(ctx, nil)

}

func installSystem(ctx *gin.Context) {
	//var cfg config.Configure
	//err := ctx.BindJSON(&config.Config)
	//if err != nil {
	//	replyError(ctx, err)
	//	return
	//}

	err := config.Store()
	if err != nil {
		replyError(ctx, err)
		return
	}

	replyOk(ctx, nil)
}
