package camera

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"iot-master/db"
	"iot-master/log"
	"iot-master/model"
	"math"
	"path/filepath"
	"strconv"
	"sync"
)

var allCameras sync.Map

//LoadCameras 加载设备
func LoadCameras() error {
	var cameras []*model.Camera
	err := db.Engine.Limit(math.MaxInt).Find(&cameras)
	if err != nil {
		return err
	}
	for _, d := range cameras {
		if d.Disabled {
			continue
		}

		dev := NewCamera(d)
		allCameras.Store(d.Id, dev)
		err = dev.Open()
		if err != nil {
			log.Error(err)
		}
	}
	return nil
}

//LoadCamera 加载设备
func LoadCamera(id int64) error {
	camera := &model.Camera{}
	has, err := db.Engine.ID(id).Get(camera)
	if err != nil {
		return err
	}
	if !has {
		return fmt.Errorf("找不到设备 %d", id)
	}

	dev := NewCamera(camera)
	allCameras.Store(id, dev)
	err = dev.Open()
	if err != nil {
		return err
	}
	return nil
}

//GetCamera 获取设备
func GetCamera(id int64) *Camera {
	d, ok := allCameras.Load(id)
	if ok {
		return d.(*Camera)
	}
	return nil
}

//RemoveCamera 删除设备
func RemoveCamera(id int64) error {
	d, ok := allCameras.LoadAndDelete(id)
	if ok {
		dev := d.(*Camera)
		return dev.Close()
	}
	return nil //error
}

func Start() error {
	err := LoadCameras()
	if err != nil {
		return err
	}

	//创建API服务，接收数据流
	gin.SetMode(gin.ReleaseMode)
	app := gin.New()

	//接收数据上传
	app.PUT("/:id/*file", func(ctx *gin.Context) {
		id := ctx.Param("id")
		file := ctx.Param("file")

		fmt.Println("upload", id, file)

		idd, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			return
		}
		c := GetCamera(idd)
		if c == nil {
			return
		}

		buf, _ := io.ReadAll(ctx.Request.Body)
		if filepath.Ext(file) == ".m3u8" {
			c.playlist = buf
		} else {
			c.segments.Enqueue(&Segment{name: file, data: buf})
			if c.segments.Size() > 6 {
				c.segments.Dequeue()
			}
		}

		ctx.String(200, "")
	})

	go func() {
		//TODO 端口改为可配置
		err := app.Run(":143")
		if err != nil {
			log.Error(err)
		}
	}()

	return nil
}

func Stop() {

}
