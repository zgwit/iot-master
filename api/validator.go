package api

import (
	"github.com/gin-gonic/gin"
	curd2 "github.com/zgwit/iot-master/v4/curd"
	"github.com/zgwit/iot-master/v4/model"
)

// @Summary 查询检查数量
// @Schemes
// @Description 查询检查数量
// @Tags validator
// @Param search body curd.ParamSearch true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[int64] 返回检查数量
// @Router /validator/count [post]
func noopValidatorCount() {}

// @Summary 查询检查
// @Schemes
// @Description 查询检查
// @Tags validator
// @Param search body curd.ParamSearch true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyList[model.Validator] 返回检查信息
// @Router /validator/search [post]
func noopValidatorSearch() {}

// @Summary 查询检查
// @Schemes
// @Description 查询检查
// @Tags validator
// @Param search query curd.ParamList true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyList[model.Validator] 返回检查信息
// @Router /validator/list [get]
func noopValidatorList() {}

// @Summary 创建检查
// @Schemes
// @Description 创建检查
// @Tags validator
// @Param search body model.Validator true "检查信息"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[model.Validator] 返回检查信息
// @Router /validator/create [post]
func noopValidatorCreate() {}

// @Summary 修改检查
// @Schemes
// @Description 修改检查
// @Tags validator
// @Param id path int true "检查ID"
// @Param validator body model.Validator true "检查信息"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[model.Validator] 返回检查信息
// @Router /validator/{id} [post]
func noopValidatorUpdate() {}

// @Summary 获取检查
// @Schemes
// @Description 获取检查
// @Tags validator
// @Param id path int true "检查ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[model.Validator] 返回检查信息
// @Router /validator/{id} [get]
func noopValidatorGet() {}

// @Summary 删除检查
// @Schemes
// @Description 删除检查
// @Tags validator
// @Param id path int true "检查ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[model.Validator] 返回检查信息
// @Router /validator/{id}/delete [get]
func noopValidatorDelete() {}

// @Summary 启用检查
// @Schemes
// @Description 启用检查
// @Tags validator
// @Param id path int true "检查ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[model.Validator] 返回检查信息
// @Router /validator/{id}/enable [get]
func noopValidatorEnable() {}

// @Summary 禁用检查
// @Schemes
// @Description 禁用检查
// @Tags validator
// @Param id path int true "检查ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[model.Validator] 返回检查信息
// @Router /validator/{id}/disable [get]
func noopValidatorDisable() {}

// @Summary 导出检查
// @Schemes
// @Description 导出检查
// @Tags validator
// @Accept json
// @Produce octet-stream
// @Router /validator/export [get]
func noopValidatorExport() {}

// @Summary 导入检查
// @Schemes
// @Description 导入检查
// @Tags validator
// @Param file formData file true "压缩包"
// @Accept mpfd
// @Produce json
// @Success 200 {object} curd.ReplyData[int64] 返回检查数量
// @Router /validator/import [post]
func noopValidatorImport() {}

func validatorRouter(app *gin.RouterGroup) {

	app.POST("/count", curd2.ApiCount[model.Validator]())
	app.POST("/search", curd2.ApiSearch[model.Validator]())
	app.GET("/list", curd2.ApiList[model.Validator]())
	app.POST("/create", curd2.ApiCreateHook[model.Validator](curd2.GenerateRandomId[model.Validator](8), nil))
	app.GET("/:id", curd2.ParseParamStringId, curd2.ApiGet[model.Validator]())
	app.POST("/:id", curd2.ParseParamStringId, curd2.ApiUpdateHook[model.Validator](nil, nil,
		"id", "product_id", "expression", "type", "title",
		"template", "level", "delay", "again", "total", "disabled"))
	app.GET("/:id/delete", curd2.ParseParamStringId, curd2.ApiDeleteHook[model.Validator](nil, nil))
	app.GET("/export", curd2.ApiExport("validator", "验证器"))
	app.POST("/import", curd2.ApiImport("validator"))

	app.GET(":id/disable", curd2.ParseParamStringId, curd2.ApiDisableHook[model.Validator](true, nil, nil))
	app.GET(":id/enable", curd2.ParseParamStringId, curd2.ApiDisableHook[model.Validator](false, nil, nil))
}
