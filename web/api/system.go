package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/protocols"
)

func systemRoutes(app *gin.RouterGroup) {
	app.GET("version")
	app.GET("cpu")
	app.GET("memory")
	app.GET("disk")
	app.GET("cron")
	app.GET("protocols", protocolList)

}

func protocolList(ctx *gin.Context) {
	ps := protocols.Protocols()
	replyOk(ctx, ps)
}