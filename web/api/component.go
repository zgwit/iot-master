package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"io/fs"
	"iot-master/config"
	"iot-master/db"
	"iot-master/lib"
	"iot-master/model"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func componentLoad(ctx *gin.Context) {
	filename := filepath.Join(config.Config.Data, "component", ctx.Param("id"), "component.json")
	http.ServeFile(ctx.Writer, ctx.Request, filename)
}

func componentSave(ctx *gin.Context) {
	dir := filepath.Join(config.Config.Data, "component", ctx.Param("id"))
	filename := filepath.Join(dir, "component.json")
	_ = os.MkdirAll(dir, os.ModePerm) //应对目录不存在的问题
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

func componentExport(ctx *gin.Context) {
	id := ctx.Param("id")
	var component model.Component
	has, err := db.Engine.ID(id).Get(&component)
	if err != nil {
		replyError(ctx, err)
		return
	} else if !has {
		replyFail(ctx, "记录不存在")
		return
	}

	ctx.Header("Content-Type", `application/x-zip-compressed`)
	ctx.Header("Content-Disposition", `attachment; filename="`+id+`.zip"`)
	//err := lib.ZipDir(dir, ctx.Writer)

	zipper := lib.NewZipper(ctx.Writer)
	defer zipper.Close()

	obj := model.Manifest[model.Component]{
		Type:  "component",
		Node:  config.Config.Node,
		Time:  time.Now(),
		Model: &component,
	}

	buf, err := json.Marshal(&obj)

	if err != nil {
		replyError(ctx, err)
		return
	}

	info := lib.NewFileInfo("manifest.json", int64(len(buf)), fs.ModePerm, time.Now(), false)

	err = zipper.CompressFileInfoAndContent("manifest.json", info, buf)
	if err != nil {
		replyError(ctx, err)
		return
	}

	dir := filepath.Join(config.Config.Data, "component", id)
	err = zipper.CompressDir(dir)
	if err != nil {
		replyError(ctx, err)
		return
	}
	//replyOk(ctx, nil)
}
