package master

import (
	"embed"
	_ "github.com/zgwit/iot-master/v3/docs"
	"github.com/zgwit/iot-master/v3/internal"
	"github.com/zgwit/iot-master/v3/internal/api"
	"github.com/zgwit/iot-master/v3/internal/app"
	"github.com/zgwit/iot-master/v3/internal/broker"
	"github.com/zgwit/iot-master/v3/pkg/banner"
	"github.com/zgwit/iot-master/v3/pkg/build"
	"github.com/zgwit/iot-master/v3/pkg/web"
	"net/http"
)

//go:embed all:www
var wwwFiles embed.FS

func Startup() error {
	banner.Print()
	build.Println()

	//加载主程序
	err := internal.Open()
	if err != nil {
		return err
	}
	defer internal.Close()

	//Web服务
	engine := web.CreateEngine()

	//注册前端接口
	api.RegisterRoutes(engine.Group("/api"))

	//注册接口文档
	web.RegisterSwaggerDocs(&engine.RouterGroup)

	//使用$前缀区分插件
	engine.Any("/app/:app/*path", app.ProxyApp)

	//监听Websocket
	engine.GET("/mqtt", broker.GinHandler)

	//前端静态文件
	engine.RegisterFS(http.FS(wwwFiles), "www", "index.html")

	//监听HTTP
	engine.Serve()

	return nil
}

func Shutdown() error {

	internal.Close()

	//只关闭Web就行了，其他通过defer关闭

	return nil
}
