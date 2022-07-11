package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"golang.org/x/mod/semver"
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
	//err := lib.ZipIntoWriter(dir, ctx.Writer)

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
	//if err != nil {
	//	replyError(ctx, err)
	//	return
	//}
	//replyOk(ctx, nil)
}

func componentImport(ctx *gin.Context) {
	//解析Body
	file, header, err := ctx.Request.FormFile("file")
	if err != nil {
		replyError(ctx, err)
		return
	}
	defer file.Close()

	tempDir := filepath.Join(os.TempDir(), lib.RandomString(40))
	err = lib.Unzip(file, header.Size, tempDir)
	if err != nil {
		replyError(ctx, err)
		return
	}

	filename := filepath.Join(tempDir, "manifest.json")
	data, err := os.ReadFile(filename)
	if err != nil {
		replyFail(ctx, "找不到入口文件")
		return
	}

	var manifest model.Manifest[model.Component]
	err = json.Unmarshal(data, &manifest)
	if err != nil {
		replyError(ctx, err)
		return
	}

	if manifest.Type != "component" {
		_ = os.RemoveAll(tempDir)
		replyFail(ctx, "不是组件文件")
		return
	}
	_ = os.Remove(filename)

	component := manifest.Model
	dir := filepath.Join(config.Config.Data, "component", component.Id)

	//删除已经存在的目录
	var old model.Component
	has, err := db.Engine.ID(component.Id).Get(&old)
	if has && err == nil {
		//版本比较，新版本取代旧版本
		if semver.Compare(component.Version, old.Version) > 0 {
			_, _ = db.Engine.ID(component.Id).Delete(&old)
		} else {
			replyFail(ctx, "已经存在，如果需要更新，请删除")
			return
		}
	}

	_, err = db.Engine.InsertOne(component)
	if err != nil {
		_ = os.RemoveAll(tempDir)
		replyError(ctx, err)
		return
	}

	_ = os.RemoveAll(dir)
	err = os.Rename(tempDir, dir)
	if err != nil {
		_ = os.RemoveAll(tempDir)
		replyError(ctx, err)
		return
	}

	replyOk(ctx, component)
}
