package api

import (
	"github.com/gin-gonic/gin"
	"io"
	"iot-master/config"
	"net/http"
	"os"
	"path/filepath"
)

func componentLoad(ctx *gin.Context) {
	filename := filepath.Join(config.Config.Data, "component", ctx.Param("id"), "component.json")
	http.ServeFile(ctx.Writer, ctx.Request, filename)
}

func componentSave(ctx *gin.Context) {
	filename := filepath.Join(config.Config.Data, "component", ctx.Param("id"), "component.json")
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
