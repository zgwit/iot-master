package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/db"
	"github.com/zgwit/iot-master/model"
)


func eventClear(ctx *gin.Context) {
	cnt, err := db.Engine.Delete(&model.Event{})
	if err != nil {
		replyError(ctx, err)
		return
	}

	replyOk(ctx, cnt)
}
