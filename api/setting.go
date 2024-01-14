package api

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/zgwit/iot-master/v4/config"
	"github.com/zgwit/iot-master/v4/pkg/web/curd"
)

// @Summary 查询配置
// @Schemes
// @Description 查询配置
// @Tags setting
// @Param module path string true "模块，web database log ..."
// @Accept json
// @Produce json
// @Success 200 {object} map[string]any 返回配置
// @Router /setting/:module [get]
func settingGet(ctx *gin.Context) {
	module := ctx.GetString("module")
	curd.OK(ctx, viper.GetStringMap(module))
}

// @Summary 修改配置
// @Schemes
// @Description 修改配置
// @Tags setting
// @Param module path string true "模块，web database log ..."
// @Param cfg body map[string]any true "配置内容"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[int]
// @Router /setting/:module [post]
func settingSet(ctx *gin.Context) {
	module := ctx.GetString("module")

	var conf map[string]any
	err := ctx.ShouldBindJSON(&conf)
	if err != nil {
		curd.Error(ctx, err)
		return
	}
	for k, v := range conf {
		viper.Set(module+"."+k, v)
	}

	err = config.Store()
	if err != nil {
		curd.Error(ctx, err)
		return
	}
	curd.OK(ctx, nil)
}

func settingRouter(app *gin.RouterGroup) {
	app.POST("/:module", settingSet)
	app.GET("/:module", settingGet)

}
