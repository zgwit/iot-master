package api

import (
	"github.com/gin-gonic/gin"
	curd "github.com/zgwit/iot-master/v4/pkg/web/curd"
	"github.com/zgwit/iot-master/v4/pkg/web/export"
	"github.com/zgwit/iot-master/v4/product"
	"github.com/zgwit/iot-master/v4/types"
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
// @Success 200 {object} curd.ReplyList[types.Product] 返回产品信息
// @Router /product/search [post]
func noopProductSearch() {}

// @Summary 查询产品
// @Schemes
// @Description 查询产品
// @Tags product
// @Param search query curd.ParamList true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyList[types.Product] 返回产品信息
// @Router /product/list [get]
func noopProductList() {}

// @Summary 创建产品
// @Schemes
// @Description 创建产品
// @Tags product
// @Param search body types.Product true "产品信息"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[types.Product] 返回产品信息
// @Router /product/create [post]
func noopProductCreate() {}

// @Summary 修改产品
// @Schemes
// @Description 修改产品
// @Tags product
// @Param id path int true "产品ID"
// @Param product body types.Product true "产品信息"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[types.Product] 返回产品信息
// @Router /product/{id} [post]
func noopProductUpdate() {}

// @Summary 获取产品
// @Schemes
// @Description 获取产品
// @Tags product
// @Param id path int true "产品ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[types.Product] 返回产品信息
// @Router /product/{id} [get]
func noopProductGet() {}

// @Summary 删除产品
// @Schemes
// @Description 删除产品
// @Tags product
// @Param id path int true "产品ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[types.Product] 返回产品信息
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

// @Summary 获取产品详情
// @Schemes
// @Description 获取产品详情
// @Tags product
// @Param id path int true "产品ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[product.Manifest] 返回产品信息
// @Router /product/{id}/manifest [get]
func noopProductManifestGet() {}

// @Summary 修改产品详情
// @Schemes
// @Description 修改产品详情
// @Tags product
// @Param id path int true "产品ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[product.Manifest] 返回产品信息
// @Router /product/{id}/stop [post]
func noopProductManifestPost() {}

func productRouter(app *gin.RouterGroup) {

	app.POST("/count", curd.ApiCount[types.Product]())
	app.POST("/search", curd.ApiSearch[types.Product]())
	app.GET("/list", curd.ApiList[types.Product]())
	app.POST("/create", curd.ApiCreateHook[types.Product](curd.GenerateRandomId[types.Product](8), nil))
	app.GET("/:id", curd.ParseParamStringId, curd.ApiGet[types.Product]())
	app.POST("/:id", curd.ParseParamStringId, curd.ApiUpdateHook[types.Product](nil, nil,
		"id", "name", "version", "desc", "properties", "functions", "events", "parameters", "validators", "aggregators"))
	app.GET("/:id/delete", curd.ParseParamStringId, curd.ApiDeleteHook[types.Product](nil, nil))
	app.GET("/export", export.ApiExport("product", "产品"))
	app.POST("/import", export.ApiImport("product"))

	app.GET(":id/manifest", curd.ParseParamStringId, func(ctx *gin.Context) {
		p := product.Get(ctx.GetString("id"))
		if p == nil {
			curd.Fail(ctx, "产品未加载")
			return
		}
		curd.OK(ctx, p.Manifest)
	})

	app.POST(":id/manifest", curd.ParseParamStringId, func(ctx *gin.Context) {
		var m product.Manifest
		err := ctx.BindJSON(&m)
		if err != nil {
			curd.Error(ctx, err)
			return
		}

		err = product.Store(ctx.GetString("id"), &m)
		if err != nil {
			curd.Error(ctx, err)
			return
		}
		curd.OK(ctx, nil)
	})
}
