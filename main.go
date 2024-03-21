package master

import (
	"github.com/spf13/viper"
	"github.com/zgwit/iot-master/v4/api"
	"github.com/zgwit/iot-master/v4/broker"
	_ "github.com/zgwit/iot-master/v4/docs"
	"github.com/zgwit/iot-master/v4/internal"
	"github.com/zgwit/iot-master/v4/pkg/config"
	"github.com/zgwit/iot-master/v4/pkg/log"
	"github.com/zgwit/iot-master/v4/web"
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

	//加载配置文件
	err := config.Load()
	if err != nil {
		log.Error(err)
		_ = config.Store()
	}

	//加载主程序
	err = internal.Open()
	if err != nil {
		return err
	}

	web.Start()

	//注册前端接口
	api.RegisterRoutes(web.Engine.Group("/api"))

	//注册接口文档
	web.RegisterSwaggerDocs(&web.Engine.RouterGroup, "master")

	//监听Websocket
	web.Engine.GET("/mqtt", broker.GinBridge)

	//附件
	//web.Engine.Static("/static", "static")
	web.Engine.Static("/attach", filepath.Join(viper.GetString("data"), "attach"))

	//前端 移入子工程 github.com/iot-master-contrib/webui
	//web.Static.Put("", http.FS(wwwFiles), "www", "index.html")

	return nil
}

func Shutdown() error {

	internal.Close()

	//只关闭Web就行了，其他通过defer关闭
	return web.Shutdown()
}
