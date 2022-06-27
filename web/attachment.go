package web

import (
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"iot-master/config"
	"net/http"
	"os"
	"path/filepath"
)

type RenameBody struct {
	Filename string `json:"filename"`
}

func replyOk(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{"data": data})
}

func replyFail(ctx *gin.Context, err string) {
	ctx.JSON(http.StatusOK, gin.H{"error": err})
}

func replyError(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
}

func registerAttachment(name string, app *gin.RouterGroup) {
	//组态的附件
	app.GET("/:id/*name", func(ctx *gin.Context) {
		filename := filepath.Join(config.Config.Data, "hmi", ctx.Param("id"), ctx.Param("name"))
		stat, err := os.Stat(filename)
		if err != nil {
			replyError(ctx, err)
			return
		}

		//列出目录
		if stat.IsDir() {
			files, err := ioutil.ReadDir(filename)
			if err != nil {
				replyError(ctx, err)
				return
			}

			items := make([]map[string]interface{}, 0)
			for _, f := range files {
				item := make(map[string]interface{})
				item["name"] = f.Name()
				item["time"] = f.ModTime()
				item["size"] = f.Size()
				item["folder"] = f.IsDir()
				items = append(items, item)
			}
			replyOk(ctx, items)
		} else {
			http.ServeFile(ctx.Writer, ctx.Request, filename)
		}
	})

	app.POST("/:id/*name", func(ctx *gin.Context) {
		dir := filepath.Join(config.Config.Data, name, ctx.Param("id"), ctx.Param("name"))
		_ = os.MkdirAll(dir, os.ModePerm) //创建目录

		//解析Body
		file, header, err := ctx.Request.FormFile("file")
		if err != nil {
			replyError(ctx, err)
			return
		}
		defer file.Close()

		//创建写入文件
		filename := filepath.Join(dir, header.Filename)
		writer, err := os.OpenFile(filename, os.O_CREATE, os.ModePerm)
		if err != nil {
			replyError(ctx, err)
			return
		}
		defer writer.Close()

		_, err = io.Copy(writer, file)
		if err != nil {
			replyError(ctx, err)
			return
		}

		replyOk(ctx, nil)
	})

	app.PATCH("/:id/*name", func(ctx *gin.Context) {
		var rename RenameBody
		err := ctx.ShouldBindJSON(&rename)
		if err != nil {
			replyError(ctx, err)
			return
		}

		filename := filepath.Join(config.Config.Data, name, ctx.Param("id"), ctx.Param("name"))
		newPath := filepath.Join(config.Config.Data, name, ctx.Param("id"), rename.Filename)

		err = os.Rename(filename, newPath)
		if err != nil {
			replyError(ctx, err)
			return
		}

		replyOk(ctx, nil)
	})

	app.DELETE("/:id/*name", func(ctx *gin.Context) {
		filename := filepath.Join(config.Config.Data, name, ctx.Param("id"), ctx.Param("name"))
		err := os.Remove(filename)
		if err != nil {
			replyError(ctx, err)
			return
		}
		replyOk(ctx, nil)
	})

}
