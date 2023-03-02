package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/v3/internal"
)

func deviceProperties(ctx *gin.Context) {
	device := internal.Devices.Load(ctx.GetString("id"))
	if device == nil {
		replyFail(ctx, "设备未上线")
		return
	}
	replyOk(ctx, device.Properties)
}
