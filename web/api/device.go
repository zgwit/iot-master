package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/database"
	"github.com/zgwit/iot-master/master"
	"github.com/zgwit/iot-master/model"
)

func deviceRoutes(app *gin.RouterGroup) {
	app.POST("list", deviceList)
	app.POST("create", deviceCreate)

	app.Use(parseParamId)
	app.POST(":id/update", deviceUpdate)
	app.GET(":id/delete", deviceDelete)
	app.GET(":id/start", deviceStart)
	app.GET(":id/stop", deviceStop)
	app.GET(":id/enable", deviceEnable)
	app.GET(":id/disable", deviceDisable)

}

func deviceList(ctx *gin.Context) {
	var devices []model.Device
	cnt, err := normalSearch(ctx, database.Master, &devices)
	if err != nil {
		replyError(ctx, err)
		return
	}
	replyList(ctx, devices, cnt)
}

func deviceCreate(ctx *gin.Context) {
	var device model.Device
	err := ctx.ShouldBindJSON(&device)
	if err != nil {
		replyError(ctx, err)
		return
	}

	err = database.Master.Save(&device)
	if err != nil {
		replyError(ctx, err)
		return
	}

	//TODO 启动

	replyOk(ctx, device)
}

func deviceUpdate(ctx *gin.Context) {
	var pid paramID
	err := ctx.ShouldBindUri(&pid)
	if err != nil {
		replyError(ctx, err)
		return
	}

	var device model.Device
	err = ctx.ShouldBindJSON(&device)
	if err != nil {
		replyError(ctx, err)
		return
	}
	device.ID = pid.ID

	err = database.Master.Update(&device)
	if err != nil {
		replyError(ctx, err)
		return
	}

	//TODO 重新启动

	replyOk(ctx, device)
}

func deviceDelete(ctx *gin.Context) {
	var pid paramID
	err := ctx.ShouldBindUri(&pid)
	if err != nil {
		replyError(ctx, err)
		return
	}
	device := model.Device{ID: pid.ID}
	err = database.Master.DeleteStruct(&device)
	if err != nil {
		replyError(ctx, err)
		return
	}

	//TODO 重新启动

	replyOk(ctx, device)
}

func deviceStart(ctx *gin.Context) {
	device := master.GetDevice(ctx.GetInt("id"))
	if device == nil {
		replyFail(ctx, "not found")
		return
	}
	err := device.Start()
	if err != nil {
		replyError(ctx, err)
		return
	}

	replyOk(ctx, nil)
}

func deviceStop(ctx *gin.Context) {
	device := master.GetDevice(ctx.GetInt("id"))
	if device == nil {
		replyFail(ctx, "not found")
		return
	}
	err := device.Stop()
	if err != nil {
		replyError(ctx, err)
		return
	}

	replyOk(ctx, nil)
}


func deviceEnable(ctx *gin.Context) {
	err := database.Master.UpdateField(&model.Device{ID: ctx.GetInt("id")}, "Disabled", false)
	if err != nil {
		replyError(ctx, err)
		return
	}
	//TODO 启动
	replyOk(ctx, nil)
}

func deviceDisable(ctx *gin.Context) {
	err := database.Master.UpdateField(&model.Device{ID: ctx.GetInt("id")}, "Disabled", true)
	if err != nil {
		replyError(ctx, err)
		return
	}
	//TODO 关闭
	replyOk(ctx, nil)
}