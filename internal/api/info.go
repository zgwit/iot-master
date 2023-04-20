package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/v3/pkg/build"
	"github.com/zgwit/iot-master/v3/pkg/curd"
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
