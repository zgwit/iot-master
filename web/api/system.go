package api

import "github.com/gin-gonic/gin"

func systemRoutes(app *gin.RouterGroup) {
	app.GET("version")
	app.GET("cpu")
	app.GET("memory")
	app.GET("disk")
	app.GET("cron")
	app.GET("protocol")

}