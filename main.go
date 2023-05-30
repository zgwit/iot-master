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

// @title 物联大师接口文档
// @version 3.2 版本
// @description API文档
// @BasePath /api/
// @InstanceName master
// @query.collection.format multi
func main() {}

func Startup(engine *web.Engine) error {
	banner.Print()
	build.Println()

	//加载主程序
	err := internal.Open()
	if err != nil {
		return err
	}
	//defer internal.Close()

	//注册前端接口
	api.RegisterRoutes(engine.Group("/api"))

	//注册接口文档
	web.RegisterSwaggerDocs(&engine.RouterGroup)

	//使用$前缀区分插件
	engine.Any("/app/:app/*path", app.ProxyApp)

	//监听Websocket
	engine.GET("/mqtt", broker.GinHandler)

	return nil
}

func Static(fs *web.FileSystem) {
	//前端静态文件
	fs.Put("", http.FS(wwwFiles), "www", "index.html")
}

func Shutdown() error {

	internal.Close()

	//只关闭Web就行了，其他通过defer关闭

	return nil
}
