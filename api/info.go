package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/curd"
	"github.com/zgwit/iot-master/pkg/build"
	"runtime"
)

func info(ctx *gin.Context) {
	curd.OK(ctx, gin.H{
		"version": build.Version,
		"build":   build.Build,
		"git":     build.GitHash,
		"gin":     gin.Version,
		"runtime": runtime.Version(),
	})
}
