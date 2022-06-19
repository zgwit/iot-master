package web

import (
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/camera"
	"path/filepath"
	"strconv"
)

func registerStreamRoutes(app *gin.RouterGroup) {
	app.GET("/:id/*file", func(ctx *gin.Context) {
		id := ctx.Param("id")
		idd, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			return
		}
		file := ctx.Param("file")
		c := camera.GetCamera(idd)
		if c == nil {
			return
		}
		ext := filepath.Ext(file)
		if ext == ".m3u8" {
			ctx.Data(200, "application/x-mpegURL", c.Playlist())
			return
		}

		ctx.Data(200, "data/MP2T", c.Segment(file))
	})

	//time.AfterFunc(5*time.Second, func() {
	//	c := camera.Camera{
	//		Camera: model.Camera{
	//			MediaUri: "rtsp://admin:Jason927@192.168.1.105:554/onvif1",
	//		},
	//	}
	//
	//	err := c.Open()
	//	if err != nil {
	//		fmt.Println(err)
	//		return
	//	}
	//
	//	//time.AfterFunc(60*time.Second, func() {
	//	//	_ = c.Close()
	//	//})
	//})
}
