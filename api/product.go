package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/v4/product"
	"github.com/zgwit/iot-master/v4/web/curd"
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

func productRouter(app *gin.RouterGroup) {

	app.POST("/count", curd.ApiCount[product.Product]())
	app.POST("/search", curd.ApiSearch[product.Product]())
	app.GET("/list", curd.ApiList[product.Product]())
	app.POST("/create", curd.ApiCreateHook[product.Product](curd.GenerateID[product.Product](), nil))
	app.GET("/:id", curd.ParseParamStringId, curd.ApiGet[product.Product]())
	app.POST("/:id", curd.ParseParamStringId, curd.ApiUpdateHook[product.Product](nil, nil,
		"id", "name", "url", "icon", "description", "keywords", "disabled"))
	app.GET("/:id/delete", curd.ParseParamStringId, curd.ApiDeleteHook[product.Product](nil, nil))

}
