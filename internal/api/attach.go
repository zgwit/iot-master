package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/v3/pkg/curd"
	"io/ioutil"
	"mime"
	"os"
	"path/filepath"
	"time"
)

const AttachRoot = "attach"

type attachInfo struct {
	Name   string    `json:"name,omitempty"`
	Mime   string    `json:"mime,omitempty"`
	Time   time.Time `json:"time"`
	Size   int64     `json:"size,omitempty"`
	Folder bool      `json:"folder,omitempty"`
}

type renameBody struct {
	Name string `json:"name,omitempty"`
}

type moveBody struct {
	Path string `json:"path,omitempty"`
}

// @Summary 查询附件
// @Schemes
// @Description 查询附件
// @Tags attach
// @Param name path string true "路径"
// @Produce json
// @Success 200 {object} curd.ReplyList[attachInfo]
// @Router /attach/list/{name} [get]
func attachList(ctx *gin.Context) {
	//列出目录
	filename := filepath.Join(AttachRoot, ctx.Param("name"))
	files, err := ioutil.ReadDir(filename)
	if err != nil {
		curd.Error(ctx, err)
		return
	}

	var items []*attachInfo
	for _, stat := range files {
		item := &attachInfo{
			Name:   stat.Name(),
			Time:   stat.ModTime(),
			Size:   stat.Size(),
			Folder: stat.IsDir(),
		}
		//识别类型
		if !stat.IsDir() {
			item.Mime = mime.TypeByExtension(filepath.Ext(stat.Name()))
		}
		items = append(items, item)
	}
	curd.OK(ctx, items)
}

// @Summary 上传附件
// @Schemes
// @Description 上传附件（支持多文件，不用特定为file）
// @Tags attach
// @Param file formData file true "附件"
// @Param name path string true "路径"
// @Accept mpfd
// @Produce json
// @Success 200 {object} curd.ReplyData[any]
// @Router /attach/upload/{name} [post]
func attachUpload(ctx *gin.Context) {
	dir := filepath.Join(AttachRoot, ctx.Param("name"))
	_ = os.MkdirAll(dir, os.ModePerm) //创建目录

	form, err := ctx.MultipartForm()
	if err != nil {
		curd.Error(ctx, err)
		return
	}

	//接收所有附件
	for _, files := range form.File {
		for _, header := range files {
			filename := filepath.Join(dir, header.Filename)
			err := ctx.SaveUploadedFile(header, filename)
			if err != nil {
				return
			}
		}
	}

	curd.OK(ctx, nil)
}

// @Summary 重命名附件
// @Schemes
// @Description 重命名附件
// @Tags attach
// @Param name path string true "路径"
// @Param b body renameBody true "新名称"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[any]
// @Router /attach/rename/{name} [post]
func attachRename(ctx *gin.Context) {
	var rename renameBody
	err := ctx.ShouldBindJSON(&rename)
	if err != nil {
		curd.Error(ctx, err)
		return
	}

	filename := filepath.Join(AttachRoot, ctx.Param("name"))
	newPath := filepath.Join(filepath.Dir(filename), rename.Name)

	err = os.Rename(filename, newPath)
	if err != nil {
		curd.Error(ctx, err)
		return
	}
	curd.OK(ctx, nil)
}

// @Summary 删除附件
// @Schemes
// @Description 删除附件
// @Tags attach
// @Param name path string true "路径"
// @Produce json
// @Success 200 {object} curd.ReplyData[any]
// @Router /attach/remove/{name} [get]
func attachRemove(ctx *gin.Context) {
	filename := filepath.Join(AttachRoot, ctx.Param("name"))
	err := os.Remove(filename)
	if err != nil {
		curd.Error(ctx, err)
		return
	}
	curd.OK(ctx, nil)
}

// @Summary 移动附件
// @Schemes
// @Description 移动附件
// @Tags attach
// @Param name path string true "路径"
// @Param b body moveBody true "新路径"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[any]
// @Router /attach/move/{name} [post]
func attachMove(ctx *gin.Context) {
	var move moveBody
	err := ctx.ShouldBindJSON(&move)
	if err != nil {
		curd.Error(ctx, err)
		return
	}

	filename := filepath.Join(AttachRoot, ctx.Param("name"))
	newPath := filepath.Join(AttachRoot, move.Path, filepath.Base(filename))

	err = os.Rename(filename, newPath)
	if err != nil {
		curd.Error(ctx, err)
		return
	}
	curd.OK(ctx, nil)
}

// @Summary 创建目录
// @Schemes
// @Description 创建目录
// @Tags attach
// @Param name path string true "路径"
// @Produce json
// @Success 200 {object} curd.ReplyData[any]
// @Router /attach/mkdir/{name} [get]
func attachMkDir(ctx *gin.Context) {
	filename := filepath.Join(AttachRoot, ctx.Param("name"))
	err := os.MkdirAll(filename, os.ModePerm)
	if err != nil {
		curd.Error(ctx, err)
		return
	}
	curd.OK(ctx, nil)
}

func attachRouter(app *gin.RouterGroup) {

	app.GET("/list/*name", attachList)
	//app.GET("/info/*name", attachInfo)

	app.POST("/upload/*name", attachUpload)

	app.POST("/rename/*name", attachRename)

	app.GET("/remove/*name", attachRemove)

	app.GET("/move/*name", attachMove)

	app.GET("/mkdir/*name", attachMkDir)
}
