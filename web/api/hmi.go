package api

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/zgwit/iot-master/config"
	"github.com/zgwit/iot-master/database"
	"github.com/zgwit/iot-master/model"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

func hmiRoutes(app *gin.RouterGroup) {
	app.POST("list", hmiList)
	app.POST("create", hmiCreate)

	app.GET(":id", hmiDetail)
	app.POST(":id", hmiUpdate)
	app.GET(":id/delete", hmiDelete)

	app.GET(":id/export")

	//组态的附件

	//attachment := app.Group(":id/attachment/")
	app.GET(":id/attachment/*name", hmiAttachmentRead)
	app.POST(":id/attachment/*name", hmiAttachmentUpload)
	app.PATCH(":id/attachment/*name", hmiAttachmentRename)
	app.DELETE(":id/attachment/*name", hmiAttachmentDelete)

}

func hmiList(ctx *gin.Context) {
	hs, cnt, err := normalSearch(ctx, database.Master, &model.HMI{})
	if err != nil {
		replyError(ctx, err)
		return
	}
	replyList(ctx, hs, cnt)
}

func hmiCreate(ctx *gin.Context) {
	var hmi model.HMI
	err := ctx.ShouldBindJSON(&hmi)
	if err != nil {
		replyError(ctx, err)
		return
	}

	//使用UUId作为Id
	if hmi.Id == "" {
		hmi.Id = uuid.NewString()
	}
	//保存
	err = database.Master.Save(&hmi)
	if err != nil {
		replyError(ctx, err)
		return
	}

	replyOk(ctx, hmi)
}

func hmiDetail(ctx *gin.Context) {
	var hmi model.HMI
	err := database.Master.One("Id", ctx.Param("id"), &hmi)
	if err != nil {
		replyError(ctx, err)
		return
	}
	replyOk(ctx, hmi)
}

func hmiUpdate(ctx *gin.Context) {
	var hmi model.HMI
	err := ctx.ShouldBindJSON(&hmi)
	if err != nil {
		replyError(ctx, err)
		return
	}
	hmi.Id = ctx.Param("id")

	err = database.Master.Update(&hmi)
	if err != nil {
		replyError(ctx, err)
		return
	}

	replyOk(ctx, hmi)
}

func hmiDelete(ctx *gin.Context) {
	hmi := model.HMI{Id: ctx.Param("id")}
	err := database.Master.DeleteStruct(&hmi)
	if err != nil {
		replyError(ctx, err)
		return
	}

	replyOk(ctx, hmi)
}

func hmiAttachmentRead(ctx *gin.Context) {
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
}

func hmiAttachmentUpload(ctx *gin.Context) {
	//dir := filepath.Join(config.Config.Data, "hmi", ctx.Param("id"))
	dir := filepath.Join(config.Config.Data, "hmi", ctx.Param("id"), ctx.Param("name"))
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
}

type RenameBody struct {
	Filename string `json:"filename"`
}

func hmiAttachmentRename(ctx *gin.Context) {
	var rename RenameBody
	err := ctx.ShouldBindJSON(&rename)
	if err != nil {
		replyError(ctx, err)
		return
	}

	filename := filepath.Join(config.Config.Data, "hmi", ctx.Param("id"), ctx.Param("name"))
	newPath := filepath.Join(config.Config.Data, "hmi", ctx.Param("id"), rename.Filename)

	err = os.Rename(filename, newPath)
	if err != nil {
		replyError(ctx, err)
		return
	}

	replyOk(ctx, nil)
}

func hmiAttachmentDelete(ctx *gin.Context) {
	filename := filepath.Join(config.Config.Data, "hmi", ctx.Param("id"), ctx.Param("name"))
	err := os.Remove(filename)
	if err != nil {
		replyError(ctx, err)
		return
	}
	replyOk(ctx, nil)
}
