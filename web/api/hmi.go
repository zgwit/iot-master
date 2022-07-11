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

func hmiLoad(ctx *gin.Context) {
	filename := filepath.Join(config.Config.Data, "hmi", ctx.Param("id"), "hmi.json")
	http.ServeFile(ctx.Writer, ctx.Request, filename)
}

func hmiSave(ctx *gin.Context) {
	dir := filepath.Join(config.Config.Data, "hmi", ctx.Param("id"))
	filename := filepath.Join(dir, "hmi.json")
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

func hmiExport(ctx *gin.Context) {
	id := ctx.Param("id")
	var hmi model.Hmi
	has, err := db.Engine.ID(id).Get(&hmi)
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

	obj := model.Manifest[model.Hmi]{
		Type:  "hmi",
		Node:  config.Config.Node,
		Time:  time.Now(),
		Model: &hmi,
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

	dir := filepath.Join(config.Config.Data, "hmi", id)
	err = zipper.CompressDir(dir)
	if err != nil {
		replyError(ctx, err)
		return
	}
	//replyOk(ctx, nil)
}
