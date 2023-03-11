package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/v3/model"
	"github.com/zgwit/iot-master/v3/pkg/db"
)

func alarmRead(ctx *gin.Context) {
	alarm := model.Alarm{
		Read: true,
	}
	cnt, err := db.Engine.ID(ctx.GetInt64("id")).Cols("read").Update(alarm)
	if err != nil {
		replyError(ctx, err)
		return
	}
	replyOk(ctx, cnt)
}
