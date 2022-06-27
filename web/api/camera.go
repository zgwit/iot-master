package api

import (
	"github.com/gin-gonic/gin"
	"iot-master/camera"
	"iot-master/db"
	"iot-master/model"
)

func cameraList(ctx *gin.Context) {
	var cameras []*model.CameraEx

	var body paramSearchEx
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		replyError(ctx, err)
		return
	}

	query := body.toQuery()
	query.Select("camera.*, " + //TODO 只返回需要的字段
		" 0 as running")
	cnt, err := query.FindAndCount(&cameras)
	if err != nil {
		replyError(ctx, err)
		return
	}
	for _, c := range cameras {
		d := camera.GetCamera(c.Id)
		if d != nil {
			c.Running = d.Running()
		}
	}

	replyList(ctx, cameras, cnt)
}

func afterCameraCreate(data interface{}) error {
	c := data.(*model.Camera)
	if !c.Disabled {
		return camera.LoadCamera(c.Id)
	}
	return nil
}

func cameraDetail(ctx *gin.Context) {
	var c model.CameraEx
	has, err := db.Engine.ID(ctx.GetInt64("id")).Get(&c.Camera)
	if err != nil {
		replyError(ctx, err)
		return
	}
	if !has {
		replyFail(ctx, "记录不存在")
		return
	}
	d := camera.GetCamera(c.Id)
	if d != nil {
		c.Running = d.Running()
	}
	replyOk(ctx, c)
}

func afterCameraUpdate(data interface{}) error {
	c := data.(*model.Camera)
	_ = camera.RemoveCamera(c.Id)
	return camera.LoadCamera(c.Id)
}

func afterCameraDelete(id interface{}) error {
	return camera.RemoveCamera(id.(int64))
}

func cameraStart(ctx *gin.Context) {
	c := camera.GetCamera(ctx.GetInt64("id"))
	if c == nil {
		replyFail(ctx, "not found")
		return
	}
	err := c.Open()
	if err != nil {
		replyError(ctx, err)
		return
	}

	replyOk(ctx, nil)
}

func cameraStop(ctx *gin.Context) {
	c := camera.GetCamera(ctx.GetInt64("id"))
	if c == nil {
		replyFail(ctx, "not found")
		return
	}
	err := c.Close()
	if err != nil {
		replyError(ctx, err)
		return
	}

	replyOk(ctx, nil)
}

func afterCameraEnable(id interface{}) error {
	_ = camera.RemoveCamera(id.(int64))
	return camera.LoadCamera(id.(int64))
}

func afterCameraDisable(id interface{}) error {
	return camera.RemoveCamera(id.(int64))
}
