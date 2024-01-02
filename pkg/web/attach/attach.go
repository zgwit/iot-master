package attach

import (
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/v4/pkg/web/curd"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

type attachInfo struct {
	Name   string    `json:"name,omitempty"`
	Mime   string    `json:"mime,omitempty"`
	Time   time.Time `json:"time"`
	Size   int64     `json:"size,omitempty"`
	Folder bool      `json:"folder,omitempty"`
}

type RenameBody struct {
	Name string `json:"name,omitempty"`
}

type MoveBody struct {
	Path string `json:"path,omitempty"`
}

func ApiList(root string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//列出目录
		filename := filepath.Join(root, ctx.Param("name"))
		files, err := os.ReadDir(filename)
		if err != nil {
			curd.Error(ctx, err)
			return
		}

		var items []*attachInfo
		for _, stat := range files {
			info, _ := stat.Info()
			item := &attachInfo{
				Name:   info.Name(),
				Time:   info.ModTime(),
				Size:   info.Size(),
				Folder: info.IsDir(),
			}
			//识别类型
			if !stat.IsDir() {
				item.Mime = mime.TypeByExtension(filepath.Ext(stat.Name()))
			}
			items = append(items, item)
		}
		curd.OK(ctx, items)
	}
}

func ApiInfo(root string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//列出目录
		filename := filepath.Join(root, ctx.Param("name"))
		info, err := os.Stat(filename)
		if err != nil {
			curd.Error(ctx, err)
			return
		}
		item := &attachInfo{
			Name:   info.Name(),
			Time:   info.ModTime(),
			Size:   info.Size(),
			Folder: info.IsDir(),
		}

		curd.OK(ctx, item)
	}
}

func ApiUpload(root string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		dir := filepath.Join(root, ctx.Param("name"))
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
					curd.Error(ctx, err)
					return
				}
			}
		}

		curd.OK(ctx, nil)
	}
}

func ApiView(root string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		filename := filepath.Join(root, ctx.Param("name"))
		http.ServeFile(ctx.Writer, ctx.Request, filename)
	}
}

func ApiDownload(root string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		filename := filepath.Join(root, ctx.Param("name"))
		ctx.Header("Content-Disposition", "attachment; filename="+ctx.Param("name"))
		http.ServeFile(ctx.Writer, ctx.Request, filename)
	}
}

func ApiRename(root string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var rename RenameBody
		err := ctx.ShouldBindJSON(&rename)
		if err != nil {
			curd.Error(ctx, err)
			return
		}

		filename := filepath.Join(root, ctx.Param("name"))
		newPath := filepath.Join(filepath.Dir(filename), rename.Name)

		err = os.Rename(filename, newPath)
		if err != nil {
			curd.Error(ctx, err)
			return
		}
		curd.OK(ctx, nil)
	}
}

func ApiRemove(root string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		filename := filepath.Join(root, ctx.Param("name"))
		err := os.Remove(filename)
		if err != nil {
			curd.Error(ctx, err)
			return
		}
		curd.OK(ctx, nil)
	}
}

func ApiMove(root string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var move MoveBody
		err := ctx.ShouldBindJSON(&move)
		if err != nil {
			curd.Error(ctx, err)
			return
		}

		filename := filepath.Join(root, ctx.Param("name"))
		newPath := filepath.Join(root, move.Path, filepath.Base(filename))

		err = os.Rename(filename, newPath)
		if err != nil {
			curd.Error(ctx, err)
			return
		}
		curd.OK(ctx, nil)
	}
}

func ApiMakeDir(root string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		filename := filepath.Join(root, ctx.Param("name"))
		err := os.MkdirAll(filename, os.ModePerm)
		if err != nil {
			curd.Error(ctx, err)
			return
		}
		curd.OK(ctx, nil)
	}
}

func Routers(root string, app *gin.RouterGroup) {

	app.GET("/list/*name", ApiList(root))

	app.GET("/info/*name", ApiInfo(root))

	app.GET("/view/*name", ApiView(root))

	app.POST("/upload/*name", ApiUpload(root))

	app.GET("/download/*name", ApiDownload(root))

	app.POST("/rename/*name", ApiRename(root))

	app.GET("/remove/*name", ApiRemove(root))

	app.POST("/move/*name", ApiMove(root))

	app.GET("/mkdir/*name", ApiMakeDir(root))
}
