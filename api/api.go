package api

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/v3/config"
	"github.com/zgwit/iot-master/v3/internal"
	"github.com/zgwit/iot-master/v3/model"
	"github.com/zgwit/iot-master/v3/pkg/curd"
	"net/http"
)

func catchError(ctx *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			//runtime.Stack()
			//debug.Stack()
			switch err.(type) {
			case error:
				curd.Error(ctx, err.(error))
			case string:
				curd.Fail(ctx, err.(string))
			default:
				ctx.JSON(http.StatusOK, gin.H{"error": err})
			}
		}
	}()
	ctx.Next()

	//TODO 内容如果为空，返回404

}

func mustLogin(ctx *gin.Context) {
	token := ctx.Request.URL.Query().Get("token")
	if token == "" {
		token = ctx.Request.Header.Get("Authorization")
		if token != "" {
			//此处需要去掉 Bearer
			token = token[7:]
		}
	}

	if token != "" {
		claims, err := jwtVerify(token)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			ctx.Abort()
		}
		ctx.Set("user", claims.Id) //与session统一
		ctx.Next()
		return
	}

	//检查Session
	session := sessions.Default(ctx)
	if user := session.Get("user"); user != nil {
		ctx.Set("user", user)
		ctx.Next()
	} else {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		ctx.Abort()
	}
}

func RegisterRoutes(app *gin.RouterGroup) {
	//错误恢复，并返回至前端
	app.Use(catchError)

	app.GET("/oem", func(ctx *gin.Context) {
		curd.OK(ctx, &config.Config.Oem)
	})

	app.GET("/info", info)

	app.GET("/auth", auth)
	app.POST("/login", login)

	//检查 session，必须登录
	app.Use(mustLogin)

	app.GET("/logout", logout)
	app.POST("/password", password)

	app.GET("/apps", func(ctx *gin.Context) {
		apps := make([]*model.App, 0)
		internal.Applications.Range(func(name string, app *model.App) bool {
			if !app.Hidden {
				apps = append(apps, app)
			}
			return true
		})
		curd.OK(ctx, apps)
	})

	//修改配置
	app.GET("/config", loadConfig)
	app.POST("/config", saveConfig)

	//用户接口
	app.GET("/user/me", userMe)
	app.POST("/user/search", curd.ApiSearch[model.User]())
	app.GET("/user/list", curd.ApiList[model.User]())
	app.POST("/user/create", curd.ParseParamId, curd.ApiCreate[model.User](nil, nil))
	app.GET("/user/:id", curd.ParseParamId, curd.ApiGet[model.User]())
	app.POST("/user/:id", curd.ParseParamId, curd.ApiModify[model.User](nil, nil,
		"username", "name", "email", "disabled"))
	app.GET("/user/:id/delete", curd.ParseParamId, curd.ApiDelete[model.User](nil, nil))
	app.GET("/user/:id/password", curd.ParseParamId, userPassword)
	app.GET("/user/:id/enable", curd.ParseParamId, curd.ApiDisable[model.User](false, nil, nil))
	app.GET("/user/:id/disable", curd.ParseParamId, curd.ApiDisable[model.User](true, nil, nil))

	//产品接口
	app.POST("/product/search", curd.ApiSearch[model.Product]())
	app.GET("/product/list", curd.ApiList[model.Product]())
	app.POST("/product/create", curd.ApiCreate[model.Product](curd.GenerateRandomKey(8), nil))
	app.GET("/product/:id", curd.ParseParamStringId, curd.ApiGet[model.Product]())
	app.POST("/product/:id", curd.ParseParamStringId, curd.ApiModify[model.Product](nil, nil,
		"id", "name", "version", "desc", "properties", "functions", "events", "parameters", "constraints"))
	app.GET("/product/:id/delete", curd.ParseParamStringId, curd.ApiDelete[model.Product](nil, nil))

	//设备接口
	app.POST("/device/search", curd.ApiSearch[model.Device]())
	app.GET("/device/list", curd.ApiList[model.Device]())
	app.POST("/device/create", curd.ApiCreate[model.Device](curd.GenerateRandomKey(12), nil))
	app.GET("/device/:id", curd.ParseParamStringId, curd.ApiGet[model.Device]())
	app.POST("/device/:id", curd.ParseParamStringId, curd.ApiModify[model.Device](nil, nil,
		"id", "gateway_id", "product_id", "group_id", "type", "name", "desc", "username", "password", "parameters", "disabled"))
	app.GET("/device/:id/delete", curd.ParseParamStringId, curd.ApiDelete[model.Device](nil, nil))

	app.GET("/device/:id/values", curd.ParseParamStringId, deviceValues)
	app.POST("/device/:id/parameters", curd.ParseParamStringId, deviceParameters)

	//设备分组接口
	app.POST("/group/search", curd.ApiSearch[model.Group]())
	app.GET("/group/list", curd.ApiList[model.Group]())
	app.POST("/group/create", curd.ApiCreate[model.Group](nil, nil))
	app.GET("/group/:id", curd.ParseParamId, curd.ApiGet[model.Group]())
	app.POST("/group/:id", curd.ParseParamId, curd.ApiModify[model.Group](nil, nil,
		"name", "desc"))
	app.GET("/group/:id/delete", curd.ParseParamId, curd.ApiDelete[model.Group](nil, nil))

	//报警日志
	app.POST("/alarm/search", curd.ApiSearch[model.Alarm]())
	app.GET("/alarm/list", curd.ApiList[model.Alarm]())
	app.GET("/alarm/:id", curd.ParseParamId, curd.ApiGet[model.Alarm]())
	app.GET("/alarm/:id/delete", curd.ParseParamId, curd.ApiDelete[model.Alarm](nil, nil))
	app.GET("/alarm/:id/read", curd.ParseParamId, alarmRead)

	//服务器接口
	app.POST("/server/search", curd.ApiSearch[model.Server]())
	app.GET("/server/list", curd.ApiList[model.Server]())
	app.POST("/server/create", curd.ApiCreate[model.Server](nil, nil))
	app.GET("/server/:id", curd.ParseParamId, curd.ApiGet[model.Server]())
	app.POST("/server/:id", curd.ParseParamId, curd.ApiModify[model.Server](nil, nil,
		"name", "type", "port", "desc", "disabled"))
	app.GET("/server/:id/delete", curd.ParseParamId, curd.ApiDelete[model.Server](nil, nil))

	//应用接口
	app.POST("/app/search", curd.ApiSearch[model.App]())
	app.GET("/app/list", curd.ApiList[model.App]())
	app.POST("/app/create", curd.ApiCreate[model.App](curd.GenerateUuidKey, nil))
	app.GET("/app/:id", curd.ParseParamStringId, curd.ApiGet[model.App]())
	app.POST("/app/:id", curd.ParseParamStringId, curd.ApiModify[model.App](nil,
		nil, "id", "name", "type", "address", "desc", "disabled"))
	app.GET("/app/:id/delete", curd.ParseParamStringId, curd.ApiDelete[model.App](nil, nil))

	//插件接口
	app.POST("/plugin/search", curd.ApiSearch[model.Plugin]())
	app.GET("/plugin/list", curd.ApiList[model.Plugin]())
	app.POST("/plugin/create", curd.ApiCreate[model.Plugin](curd.GenerateUuidKey, nil))
	app.GET("/plugin/:id", curd.ParseParamStringId, curd.ApiGet[model.Plugin]())
	app.POST("/plugin/:id", curd.ParseParamStringId, curd.ApiModify[model.Plugin](nil, nil,
		"id", "name", "version", "command", "dependencies"))
	app.GET("/plugin/:id/delete", curd.ParseParamStringId, curd.ApiDelete[model.Plugin](nil, nil))

	//系统接口
	app.GET("/system/cpu-info", cpuInfo)
	app.GET("/system/cpu", cpuStats)
	app.GET("/system/memory", memStats)
	app.GET("/system/disk", diskStats)
	app.GET("/system/machine", machineInfo)

	//TODO 报接口错误（以下代码不生效，路由好像不是树形处理）
	app.Use(func(ctx *gin.Context) {
		curd.Fail(ctx, "Not found")
		ctx.Abort()
	})
}
