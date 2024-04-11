package master

import (
	"github.com/spf13/viper"
	"github.com/zgwit/iot-master/v4/api"
	"github.com/zgwit/iot-master/v4/boot"
	"github.com/zgwit/iot-master/v4/broker"     //内置MQTT服务器
	_ "github.com/zgwit/iot-master/v4/device"   //设备
	_ "github.com/zgwit/iot-master/v4/docs"     //文档
	_ "github.com/zgwit/iot-master/v4/modbus"   //默认集成Modbus协议
	_ "github.com/zgwit/iot-master/v4/product"  //产品
	_ "github.com/zgwit/iot-master/v4/project"  //项目
	_ "github.com/zgwit/iot-master/v4/protocol" //协议
	_ "github.com/zgwit/iot-master/v4/space"    //空间
	_ "github.com/zgwit/iot-master/v4/tunnel"   //通道
	_ "github.com/zgwit/iot-master/v4/user"     //通道
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

	err := boot.Startup()
	if err != nil {
		return err
	}

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
	return boot.Shutdown()
}
