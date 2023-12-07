package attach

import (
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/v4/pkg/web/curd"
	"mime"
	"net/http"
	"os"
	"path/filepath"
)

func ObjectApiList(root string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//列出目录
		filename := filepath.Join(root, ctx.Param("id"), ctx.Param("name"))
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

func ObjectApiInfo(root string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//列出目录
		filename := filepath.Join(root, ctx.Param("id"), ctx.Param("name"))
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

func ObjectApiUpload(root string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		dir := filepath.Join(root, ctx.Param("id"), ctx.Param("name"))
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
}

func ObjectApiView(root string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		filename := filepath.Join(root, ctx.Param("id"), ctx.Param("name"))
		http.ServeFile(ctx.Writer, ctx.Request, filename)
	}
}

func ObjectApiDownload(root string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		filename := filepath.Join(root, ctx.Param("id"), ctx.Param("name"))
		ctx.Header("Content-Disposition", "attachment; filename="+ctx.Param("name"))
		http.ServeFile(ctx.Writer, ctx.Request, filename)
	}
}

func ObjectApiRename(root string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var rename RenameBody
		err := ctx.ShouldBindJSON(&rename)
		if err != nil {
			curd.Error(ctx, err)
			return
		}

		filename := filepath.Join(root, ctx.Param("id"), ctx.Param("name"))
		newPath := filepath.Join(filepath.Dir(filename), rename.Name)

		err = os.Rename(filename, newPath)
		if err != nil {
			curd.Error(ctx, err)
			return
		}
		curd.OK(ctx, nil)
	}
}

func ObjectApiRemove(root string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		filename := filepath.Join(root, ctx.Param("id"), ctx.Param("name"))
		err := os.Remove(filename)
		if err != nil {
			curd.Error(ctx, err)
			return
		}
		curd.OK(ctx, nil)
	}
}

func ObjectApiMove(root string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var move MoveBody
		err := ctx.ShouldBindJSON(&move)
		if err != nil {
			curd.Error(ctx, err)
			return
		}

		filename := filepath.Join(root, ctx.Param("id"), ctx.Param("name"))
		newPath := filepath.Join(root, ctx.Param("id"), move.Path, filepath.Base(filename))

		err = os.Rename(filename, newPath)
		if err != nil {
			curd.Error(ctx, err)
			return
		}
		curd.OK(ctx, nil)
	}
}

func ObjectApiMakeDir(root string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		filename := filepath.Join(root, ctx.Param("id"), ctx.Param("name"))
		err := os.MkdirAll(filename, os.ModePerm)
		if err != nil {
			curd.Error(ctx, err)
			return
		}
		curd.OK(ctx, nil)
	}
}

func ObjectRouters(root string, app *gin.RouterGroup) {

	group := app.Group("/:id/attach")

	group.GET("/list/*name", ObjectApiList(root))

	group.GET("/info/*name", ObjectApiInfo(root))

	group.GET("/view/*name", ObjectApiView(root))

	group.POST("/upload/*name", ObjectApiUpload(root))

	group.GET("/download/*name", ObjectApiDownload(root))

	group.POST("/rename/*name", ObjectApiRename(root))

	group.GET("/remove/*name", ObjectApiRemove(root))

	group.POST("/move/*name", ObjectApiMove(root))

	group.GET("/mkdir/*name", ObjectApiMakeDir(root))
}
