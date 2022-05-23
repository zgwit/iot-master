package api

import (
	"archive/zip"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/zgwit/iot-master/database"
	"github.com/zgwit/iot-master/model"
	"io"
	"os"
)

func hmiRoutes(app *gin.RouterGroup) {
	app.POST("list", hmiList)
	app.POST("create", hmiCreate)

	app.Use(parseParamStringId)

	app.GET(":id", hmiDetail)
	app.POST(":id", hmiUpdate)
	app.GET(":id/delete", hmiDelete)

	app.POST(":id/upload", hmiUpload)
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

func hmiUpload(ctx *gin.Context) {
	id := ctx.GetString("id")

	zf, err := os.OpenFile(id+".zip", os.O_CREATE, os.ModePerm)
	if err != nil {
		replyError(ctx, err)
	}
	defer zf.Close()

	zfw := zip.NewWriter(zf)
	defer zfw.Close()

	file, header, err := ctx.Request.FormFile("file")
	if err != nil {
		replyError(ctx, err)
		return
	}
	defer file.Close()
	writer, err := zfw.Create(header.Filename)
	if err != nil {
		replyError(ctx, err)
		return
	}

	_, err = io.Copy(writer, file)
	if err != nil {
		replyError(ctx, err)
		return
	}

	err = zfw.Flush()
	if err != nil {
		replyError(ctx, err)
		return
	}

	replyOk(ctx, nil)
}
