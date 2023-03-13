package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/v3/args"
	"github.com/zgwit/iot-master/v3/pkg/curd"
	"runtime"
)

func info(ctx *gin.Context) {
	curd.OK(ctx, gin.H{
		"version": args.Version,
		"build":   args.BuildTime,
		"git":     args.GitHash,
		"gin":     gin.Version,
		"runtime": runtime.Version(),
	})
}
