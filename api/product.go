package api

import (
	"github.com/gin-gonic/gin"
	curd2 "github.com/zgwit/iot-master/v4/curd"
	"github.com/zgwit/iot-master/v4/model"
)

// @Summary 查询产品数量
// @Schemes
// @Description 查询产品数量
// @Tags product
// @Param search body curd.ParamSearch true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[int64] 返回产品数量
// @Router /product/count [post]
func noopProductCount() {}

// @Summary 查询产品
// @Schemes
// @Description 查询产品
// @Tags product
// @Param search body curd.ParamSearch true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyList[model.Product] 返回产品信息
// @Router /product/search [post]
func noopProductSearch() {}

// @Summary 查询产品
// @Schemes
// @Description 查询产品
// @Tags product
// @Param search query curd.ParamList true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyList[model.Product] 返回产品信息
// @Router /product/list [get]
func noopProductList() {}

// @Summary 创建产品
// @Schemes
// @Description 创建产品
// @Tags product
// @Param search body model.Product true "产品信息"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[model.Product] 返回产品信息
// @Router /product/create [post]
func noopProductCreate() {}

// @Summary 修改产品
// @Schemes
// @Description 修改产品
// @Tags product
// @Param id path int true "产品ID"
// @Param product body model.Product true "产品信息"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[model.Product] 返回产品信息
// @Router /product/{id} [post]
func noopProductUpdate() {}

// @Summary 获取产品
// @Schemes
// @Description 获取产品
// @Tags product
// @Param id path int true "产品ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[model.Product] 返回产品信息
// @Router /product/{id} [get]
func noopProductGet() {}

// @Summary 删除产品
// @Schemes
// @Description 删除产品
// @Tags product
// @Param id path int true "产品ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[model.Product] 返回产品信息
// @Router /product/{id}/delete [get]
func noopProductDelete() {}

// @Summary 导出产品
// @Schemes
// @Description 导出产品
// @Tags product
// @Accept json
// @Produce octet-stream
// @Router /product/export [get]
func noopProductExport() {}

// @Summary 导入产品
// @Schemes
// @Description 导入产品
// @Tags product
// @Param file formData file true "压缩包"
// @Accept mpfd
// @Produce json
// @Success 200 {object} curd.ReplyData[int64] 返回产品数量
// @Router /product/import [post]
func noopProductImport() {}

func productRouter(app *gin.RouterGroup) {

	app.POST("/count", curd2.ApiCount[model.Product]())
	app.POST("/search", curd2.ApiSearch[model.Product]())
	app.GET("/list", curd2.ApiList[model.Product]())
	app.POST("/create", curd2.ApiCreateHook[model.Product](curd2.GenerateRandomId[model.Product](8), nil))
	app.GET("/:id", curd2.ParseParamStringId, curd2.ApiGet[model.Product]())
	app.POST("/:id", curd2.ParseParamStringId, curd2.ApiUpdateHook[model.Product](nil, nil,
		"id", "name", "version", "desc", "properties", "functions", "events", "parameters", "validators", "aggregators"))
	app.GET("/:id/delete", curd2.ParseParamStringId, curd2.ApiDeleteHook[model.Product](nil, nil))
	app.GET("/export", curd2.ApiExport("product", "产品"))
	app.POST("/import", curd2.ApiImport("product"))
}
