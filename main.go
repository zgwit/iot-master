package master

import (
	"embed"
	"encoding/json"
	paho "github.com/eclipse/paho.mqtt.golang"
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/v3/api"
	"github.com/zgwit/iot-master/v3/app"
	"github.com/zgwit/iot-master/v3/broker"
	_ "github.com/zgwit/iot-master/v3/docs"
	"github.com/zgwit/iot-master/v3/internal"
	"github.com/zgwit/iot-master/v3/model"
	"github.com/zgwit/iot-master/v3/pkg/log"
	"github.com/zgwit/iot-master/v3/pkg/mqtt"
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

	//加载主程序
	err := internal.Open()
	if err != nil {
		return err
	}
	//defer internal.Close()

	//注册前端接口
	api.RegisterRoutes(engine.Group("/api"))

	//注册接口文档
	web.RegisterSwaggerDocs(&engine.RouterGroup, "master")

	//附件
	engine.Static("/attach", "attach")

	//监听Websocket
	engine.GET("/mqtt", broker.GinHandler)

	//监听插件
	mqtt.Client.Subscribe("master/register", 0, func(client paho.Client, message paho.Message) {
		var a model.App
		err := json.Unmarshal(message.Payload(), &a)
		if err != nil {
			log.Error(err)
			return
		}
		log.Info("app register ", a.Id, " ", a.Name, " ", a.Type, " ", a.Address)
		app.Applications.Store(a.Id, &a)

		//插件反向代理
		engine.Any("/app/"+a.Id+"/*path", func(ctx *gin.Context) {
			rp, err := web.CreateReverseProxy(a.Type, a.Address)
			if err != nil {
				_ = ctx.Error(err)
				return
			}
			rp.ServeHTTP(ctx.Writer, ctx.Request)
			ctx.Abort()
		})
	})

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
