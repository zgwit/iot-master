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

	app.Use(parseParamStringId)

	app.GET(":id", hmiDetail)
	app.POST(":id", hmiUpdate)
	app.GET(":id/delete", hmiDelete)

	app.GET(":id/export")

	//组态的附件
	app.GET(":id/attachment/list", hmiAttachments)
	app.POST(":id/attachment/upload", hmiAttachmentUpload)
	app.POST(":id/attachment/rename", hmiAttachmentRename)
	app.GET(":id/attachment/:name", hmiAttachment)
	app.GET(":id/attachment/:name/delete", hmiAttachmentDelete)
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
	err := database.Master.One("Id", ctx.GetString("id"), &hmi)
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
	hmi.Id = ctx.GetString("id")

	err = database.Master.Update(&hmi)
	if err != nil {
		replyError(ctx, err)
		return
	}

	replyOk(ctx, hmi)
}

func hmiDelete(ctx *gin.Context) {
	hmi := model.HMI{Id: ctx.GetString("id")}
	err := database.Master.DeleteStruct(&hmi)
	if err != nil {
		replyError(ctx, err)
		return
	}

	replyOk(ctx, hmi)
}

func hmiAttachments(ctx *gin.Context) {
	id := ctx.GetString("id")
	dir := filepath.Join(config.Config.Data, "hmi", id)
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		replyError(ctx, err)
		return
	}

	items := make([]map[string]interface{}, 0)
	item := make(map[string]interface{})
	for _, f := range files {
		item["name"] = f.Name()
		item["time"] = f.ModTime()
		item["size"] = f.Size()
	}
	replyOk(ctx, items)
}

func hmiAttachmentUpload(ctx *gin.Context) {
	id := ctx.GetString("id")

	//创建目录
	dir := filepath.Join(config.Config.Data, "hmi", id)
	_ = os.MkdirAll(dir, os.ModePerm)

	//zf, err := os.OpenFile(id+".zip", os.O_CREATE, os.ModePerm)
	file, header, err := ctx.Request.FormFile("file")
	if err != nil {
		replyError(ctx, err)
		return
	}
	defer file.Close()


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
	Old string `json:"old"`
	New string `json:"new"`
}

func hmiAttachmentRename(ctx *gin.Context) {
	var rename RenameBody
	err := ctx.ShouldBindJSON(&rename)
	if err != nil {
		replyError(ctx, err)
		return
	}

	id := ctx.GetString("id")
	dir := filepath.Join(config.Config.Data, "hmi", id)

	oldPath := filepath.Join(dir, rename.Old)
	newPath := filepath.Join(dir, rename.New)

	err = os.Rename(oldPath, newPath)
	if err != nil {
		replyError(ctx, err)
		return
	}

	replyOk(ctx, nil)
}

func hmiAttachment(ctx *gin.Context) {
	id := ctx.GetString("id")
	dir := filepath.Join(config.Config.Data, "hmi", id)
	name := ctx.Param("name")
	filename := filepath.Join(dir, name)
	http.ServeFile(ctx.Writer, ctx.Request, filename)
}

func hmiAttachmentDelete(ctx *gin.Context) {
	id := ctx.GetString("id")
	dir := filepath.Join(config.Config.Data, "hmi", id)
	name := ctx.Param("name")
	filename := filepath.Join(dir, name)
	err := os.Remove(filename)
	if err != nil {
		replyError(ctx, err)
		return
	}
	replyOk(ctx, nil)
}
