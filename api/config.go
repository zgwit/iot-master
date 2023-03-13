package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/v3/config"
	"github.com/zgwit/iot-master/v3/pkg/curd"
)

func loadConfig(ctx *gin.Context) {
	curd.OK(ctx, &config.Config)
}

func saveConfig(ctx *gin.Context) {
	var conf config.Configure
	err := ctx.BindJSON(&conf)
	if err != nil {
		curd.Error(ctx, err)
		return
	}
	config.Config = conf
	err = config.Store()
	if err != nil {
		curd.Error(ctx, err)
		return
	}

	curd.OK(ctx, nil)
}
