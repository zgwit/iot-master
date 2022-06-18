package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/camera"
	"github.com/zgwit/iot-master/model"
	"io"
	"path/filepath"
	"sync"
	"time"
)

var videos sync.Map

func registerCameraRoutes(app *gin.RouterGroup) {

	app.PUT("/*file", func(ctx *gin.Context) {
		//解析Body
		buf, _ := io.ReadAll(ctx.Request.Body)

		fmt.Println(ctx.Param("file"), len(buf))
		videos.Store(ctx.Param("file"), buf)

		ctx.JSON(200, "")
	})

	app.GET("*file", func(ctx *gin.Context) {
		file := ctx.Param("file")
		if val, ok := videos.Load(file); ok {
			video := val.([]byte)
			ext := filepath.Ext(file)
			if ext == ".m3u8" {
				ctx.Data(200, "application/x-mpegURL", video)
			} else if ext == ".ts" {
				ctx.Data(200, "video/MP2T", video)
			}
		}
	})

	time.AfterFunc(5*time.Second, func() {
		c := camera.Camera{
			Camera: model.Camera{
				MediaUri: "rtsp://admin:Jason927@192.168.1.105:554/onvif1",
			},
		}

		err := c.Open()
		if err != nil {
			fmt.Println(err)
		}
	})
}
