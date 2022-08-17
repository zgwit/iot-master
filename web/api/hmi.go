package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"golang.org/x/mod/semver"
	"io"
	"io/fs"
	"iot-master/db"
	"iot-master/internal/config"
	"iot-master/model"
	lib2 "iot-master/pkg/lib"
	"iot-master/pkg/zip"
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
	//err := lib.ZipIntoWriter(dir, ctx.Writer)

	zipper := zip.NewZipper(ctx.Writer)
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

	info := lib2.NewFileInfo("manifest.json", int64(len(buf)), fs.ModePerm, time.Now(), false)

	err = zipper.CompressFileInfoAndContent("manifest.json", info, buf)
	if err != nil {
		replyError(ctx, err)
		return
	}

	dir := filepath.Join(config.Config.Data, "hmi", id)
	err = zipper.CompressDir(dir)
	//if err != nil {
	//	replyError(ctx, err)
	//	return
	//}
	//replyOk(ctx, nil)
}

func hmiImport(ctx *gin.Context) {
	//解析Body
	file, header, err := ctx.Request.FormFile("file")
	if err != nil {
		replyError(ctx, err)
		return
	}
	defer file.Close()

	tempDir := filepath.Join(os.TempDir(), lib2.RandomString(40))
	err = zip.Unzip(file, header.Size, tempDir)
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

	var manifest model.Manifest[model.Hmi]
	err = json.Unmarshal(data, &manifest)
	if err != nil {
		replyError(ctx, err)
		return
	}

	if manifest.Type != "hmi" {
		_ = os.RemoveAll(tempDir)
		replyFail(ctx, "不是组态文件")
		return
	}
	_ = os.Remove(filename)

	hmi := manifest.Model
	dir := filepath.Join(config.Config.Data, "hmi", hmi.Id)

	//删除已经存在的目录
	var old model.Hmi
	has, err := db.Engine.ID(hmi.Id).Get(&old)
	if has && err == nil {
		//版本比较，新版本取代旧版本
		if semver.Compare(hmi.Version, old.Version) > 0 {
			_, _ = db.Engine.ID(hmi.Id).Delete(&old)
		} else {
			replyFail(ctx, "已经存在，如果需要更新，请删除")
			return
		}
	}

	_, err = db.Engine.InsertOne(hmi)
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

	replyOk(ctx, hmi)
}
