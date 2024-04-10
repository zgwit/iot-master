package settings

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/zgwit/iot-master/v4/api"
	"github.com/zgwit/iot-master/v4/config"
	"github.com/zgwit/iot-master/v4/web/curd"
)

// @Summary 查询配置
// @Schemes
// @Description 查询配置
// @Tags setting
// @Param module path string true "模块，web database log ..."
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[map[string]any] 返回配置
// @Router /setting/:module [get]
func settingGet(ctx *gin.Context) {
	module := ctx.Param("module")
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
	module := ctx.Param("module")

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

// @Summary 查询配置表单
// @Schemes
// @Description 查询配置表单
// @Tags setting
// @Param module path string true "模块，web database log ..."
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[Module] 返回配置表单
// @Router /setting/:module/form [get]
func settingForm(ctx *gin.Context) {
	module := ctx.Param("module")
	m := modules.Load(module)
	if m == nil {
		curd.Fail(ctx, "模块不存在")
		return
	}
	curd.OK(ctx, m)
}

// @Summary 查询所有配置
// @Schemes
// @Description 查询所有配置
// @Tags setting
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[[]Module] 返回配置表单
// @Router /setting/modules [get]
func settingModules(ctx *gin.Context) {
	var ms []*Module
	modules.Range(func(_ string, item *Module) bool {
		ms = append(ms, item)
		return true
	})
	curd.OK(ctx, ms)
}

func init() {
	api.Register("POST", "setting/:module", settingSet)
	api.Register("GET", "setting/:module", settingGet)
	api.Register("GET", "setting/:module/form", settingForm)
	api.Register("GET", "setting/modules", settingModules)
}
