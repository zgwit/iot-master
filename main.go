package master

import (
	"embed"
	"github.com/zgwit/iot-master/v4/api"
	"github.com/zgwit/iot-master/v4/broker"
	"github.com/zgwit/iot-master/v4/config"
	"github.com/zgwit/iot-master/v4/core"
	_ "github.com/zgwit/iot-master/v4/docs"
	"github.com/zgwit/iot-master/v4/log"
	web2 "github.com/zgwit/iot-master/v4/web"
	"net/http"
)

//go:embed all:www
var wwwFiles embed.FS

// @title 物联大师接口文档
// @version 4.0 版本
// @description API文档
// @BasePath /api/
// @InstanceName master
// @query.collection.format multi
func main() {}

func Startup(engine *web2.Engine) error {

	//加载配置文件
	err := config.Load()
	if err != nil {
		log.Error(err)
		_ = config.Store()
	}

	//加载主程序
	err = core.Open()
	if err != nil {
		return err
	}

	//defer internal.Close()
	engine.Static("/static", "static")

	//注册前端接口
	api.RegisterRoutes(engine.Group("/api"))

	//注册接口文档
	web2.RegisterSwaggerDocs(&engine.RouterGroup, "master")

	//附件
	engine.Static("/attach", "attach")

	//监听Websocket
	engine.GET("/mqtt", broker.GinBridge)

	//监听插件
	//mqtt.Subscribe[types.App]("master/register", func(topic string, a *types.App) {
	//	log.Info("app register ", a.id, " ", a.name, " ", a.Type, " ", a.Address)
	//	plugin.Applications.Store(a.id, a)
	//
	//	//插件反向代理
	//	engine.Any("/app/"+a.id+"/*path", func(ctx *gin.Context) {
	//		rp, err := web.CreateReverseProxy(a.Type, a.Address)
	//		if err != nil {
	//			_ = ctx.Error(err)
	//			return
	//		}
	//		rp.ServeHTTP(ctx.Writer, ctx.Request)
	//		ctx.Abort()
	//	})
	//})

	return nil
}

func Static(fs *web2.FileSystem) {
	//前端静态文件
	fs.Put("", http.FS(wwwFiles), "www", "index.html")
}

func Shutdown() error {

	core.Close()

	//只关闭Web就行了，其他通过defer关闭

	return nil
}
