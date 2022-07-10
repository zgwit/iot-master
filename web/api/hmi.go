package api

import (
	"github.com/gin-gonic/gin"
	"io"
	"iot-master/config"
	"iot-master/lib"
	"net/http"
	"os"
	"path/filepath"
)

func hmiLoad(ctx *gin.Context) {
	filename := filepath.Join(config.Config.Data, "hmi", ctx.Param("id"), "hmi.json")
	http.ServeFile(ctx.Writer, ctx.Request, filename)
}

func hmiSave(ctx *gin.Context) {
	filename := filepath.Join(config.Config.Data, "hmi", ctx.Param("id"), "hmi.json")
	file, err := os.OpenFile(filename, os.O_CREATE, os.ModePerm)
	if err != nil {
		replyError(ctx, err)
		return
	}
	defer file.Close()

	_, err = io.Copy(file, ctx.Request.Body)
	if err != nil {
		replyError(ctx, err)
		return
	}

	replyOk(ctx, nil)
}

func hmiExport(ctx *gin.Context) {
	dir := filepath.Join(config.Config.Data, "hmi", ctx.Param("id"))
	ctx.Header("Content-Type", `application/x-zip-compressed`)
	ctx.Header("Content-Disposition", `attachment; filename="`+ctx.Param("id")+`.zip"`)
	err := lib.ZipDir(dir, ctx.Writer)

	if err != nil {
		replyError(ctx, err)
		return
	}
	//replyOk(ctx, nil)
}
