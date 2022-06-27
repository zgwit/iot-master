package api

import (
	"github.com/gin-gonic/gin"
	"iot-master/db"
	"iot-master/model"
)

func eventClear(ctx *gin.Context) {
	cnt, err := db.Engine.Delete(&model.Event{})
	if err != nil {
		replyError(ctx, err)
		return
	}

	replyOk(ctx, cnt)
}
