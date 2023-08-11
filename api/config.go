package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/v4/pkg/curd"
	"github.com/zgwit/iot-master/v4/pkg/db"
	"github.com/zgwit/iot-master/v4/pkg/log"
	"github.com/zgwit/iot-master/v4/pkg/mqtt"
	"github.com/zgwit/iot-master/v4/pkg/oem"
	"github.com/zgwit/iot-master/v4/pkg/web"
)

// @Summary 查询WEB配置
// @Schemes
// @Description 查询WEB配置
// @Tags config
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[web.Options] 返回WEB配置
// @Router /config/web [get]
func configGetWeb(ctx *gin.Context) {
	curd.OK(ctx, web.GetOptions())
}

// @Summary 修改WEB配置
// @Schemes
// @Description 修改WEB配置
// @Tags config
// @Param cfg body web.Options true "WEB配置"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[int]
// @Router /config/web [post]
func configSetWeb(ctx *gin.Context) {
	var conf web.Options
	err := ctx.BindJSON(&conf)
	if err != nil {
		curd.Error(ctx, err)
		return
	}
	web.SetOptions(conf)
	err = web.Store()
	if err != nil {
		curd.Error(ctx, err)
		return
	}
	curd.OK(ctx, nil)
}

// @Summary 查询OEM配置
// @Schemes
// @Description 查询OEM配置
// @Tags config
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[oem.Options] 返回OEM配置
// @Router /config/oem [get]
func configGetOem(ctx *gin.Context) {
	curd.OK(ctx, oem.GetOptions())
}

// @Summary 修改OEM配置
// @Schemes
// @Description 修改OEM配置
// @Tags config
// @Param cfg body oem.Options true "OEM配置"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[int]
// @Router /config/oem [post]
func configSetOem(ctx *gin.Context) {
	var conf oem.Options
	err := ctx.BindJSON(&conf)
	if err != nil {
		curd.Error(ctx, err)
		return
	}
	oem.SetOptions(conf)
	err = oem.Store()
	if err != nil {
		curd.Error(ctx, err)
		return
	}
	curd.OK(ctx, nil)
}

// @Summary 查询日志配置
// @Schemes
// @Description 查询日志配置
// @Tags config
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[log.Options] 返回日志配置
// @Router /config/log [get]
func configGetLog(ctx *gin.Context) {
	curd.OK(ctx, log.GetOptions())
}

// @Summary 修改日志配置
// @Schemes
// @Description 修改日志配置
// @Tags config
// @Param cfg body log.Options true "日志配置"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[int]
// @Router /config/log [post]
func configSetLog(ctx *gin.Context) {
	var conf log.Options
	err := ctx.BindJSON(&conf)
	if err != nil {
		curd.Error(ctx, err)
		return
	}
	log.SetOptions(conf)
	err = log.Store()
	if err != nil {
		curd.Error(ctx, err)
		return
	}
	curd.OK(ctx, nil)
}

// @Summary 查询MQTT配置
// @Schemes
// @Description 查询MQTT配置
// @Tags config
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[mqtt.Options] 返回MQTT配置
// @Router /config/mqtt [get]
func configGetMqtt(ctx *gin.Context) {
	curd.OK(ctx, mqtt.GetOptions())
}

// @Summary 修改MQTT配置
// @Schemes
// @Description 修改MQTT配置
// @Tags config
// @Param cfg body mqtt.Options true "MQTT配置"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[int]
// @Router /config/mqtt [post]
func configSetMqtt(ctx *gin.Context) {
	var conf mqtt.Options
	err := ctx.BindJSON(&conf)
	if err != nil {
		curd.Error(ctx, err)
		return
	}
	mqtt.SetOptions(conf)
	err = mqtt.Store()
	if err != nil {
		curd.Error(ctx, err)
		return
	}
	curd.OK(ctx, nil)
}

// @Summary 查询数据库配置
// @Schemes
// @Description 查询数据库配置
// @Tags config
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[db.Options] 返回数据库配置
// @Router /config/db [get]
func configGetDatabase(ctx *gin.Context) {
	curd.OK(ctx, db.GetOptions())
}

// @Summary 修改数据库配置
// @Schemes
// @Description 修改数据库配置
// @Tags config
// @Param cfg body db.Options true "数据库配置"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[int]
// @Router /config/db [post]
func configSetDatabase(ctx *gin.Context) {
	var conf db.Options
	err := ctx.BindJSON(&conf)
	if err != nil {
		curd.Error(ctx, err)
		return
	}

	db.SetOptions(conf)
	err = db.Store()
	if err != nil {
		curd.Error(ctx, err)
		return
	}
	curd.OK(ctx, nil)
}

func configRouter(app *gin.RouterGroup) {

	app.POST("/web", configSetWeb)
	app.GET("/web", configGetWeb)

	app.POST("/oem", configSetOem)
	app.GET("/oem", configGetOem)

	app.POST("/log", configSetLog)
	app.GET("/log", configGetLog)

	app.POST("/mqtt", configSetMqtt)
	app.GET("/mqtt", configGetMqtt)

	app.POST("/database", configSetDatabase)
	app.GET("/database", configGetDatabase)

}
