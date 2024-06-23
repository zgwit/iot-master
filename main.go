package master

import (
	"github.com/god-jason/bucket/api"
	"github.com/god-jason/bucket/boot"
	"github.com/god-jason/bucket/web"
	"github.com/spf13/viper"
	"path/filepath"
)

// @title 物联大师接口文档
// @version 4.0 版本
// @description API文档
// @BasePath /api/
// @InstanceName master
// @query.collection.format multi
func main() {}

func Startup() error {

	err := boot.Startup()
	if err != nil {
		return err
	}

	//注册前端接口
	api.RegisterRoutes(web.Engine.Group("/api"))

	//监听Websocket
	//web.Engine.GET("/mqtt", broker.GinBridge)

	//附件
	//web.Engine.Static("/static", "static")
	web.Engine.Static("/attach", filepath.Join(viper.GetString("data"), "attach"))

	//前端 移入子工程 github.com/iot-master-contrib/webui
	//web.Static.Put("", http.FS(wwwFiles), "www", "index.html")

	return nil
}

func Shutdown() error {
	return boot.Shutdown()
}
