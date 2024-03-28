package product

import (
	"github.com/zgwit/iot-master/v4/api"
	"github.com/zgwit/iot-master/v4/web/curd"
)

func init() {
	api.Register("POST", "/product/count", curd.ApiCount[Product]())
	api.Register("POST", "/product/search", curd.ApiSearch[Product]())
	api.Register("GET", "/product/list", curd.ApiList[Product]())
	api.Register("POST", "/product/create", curd.ApiCreateHook[Product](curd.GenerateID[Product](), nil))
	api.Register("GET", "/product/:id", curd.ParseParamStringId, curd.ApiGet[Product]())
	api.Register("POST", "/product/:id", curd.ParseParamStringId, curd.ApiUpdateHook[Product](nil, nil,
		"id", "name", "url", "icon", "description", "keywords", "protocol", "disabled"))
	api.Register("GET", "/product/:id/delete", curd.ParseParamStringId, curd.ApiDeleteHook[Product](nil, nil))
}

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
// @Success 200 {object} curd.ReplyList[Product] 返回产品信息
// @Router /product/search [post]
func noopProductSearch() {}

// @Summary 查询产品
// @Schemes
// @Description 查询产品
// @Tags product
// @Param search query curd.ParamList true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyList[Product] 返回产品信息
// @Router /product/list [get]
func noopProductList() {}

// @Summary 创建产品
// @Schemes
// @Description 创建产品
// @Tags product
// @Param search body Product true "产品信息"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[Product] 返回产品信息
// @Router /product/create [post]
func noopProductCreate() {}

// @Summary 修改产品
// @Schemes
// @Description 修改产品
// @Tags product
// @Param id path int true "产品ID"
// @Param product body Product true "产品信息"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[Product] 返回产品信息
// @Router /product/{id} [post]
func noopProductUpdate() {}

// @Summary 获取产品
// @Schemes
// @Description 获取产品
// @Tags product
// @Param id path int true "产品ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[Product] 返回产品信息
// @Router /product/{id} [get]
func noopProductGet() {}

// @Summary 删除产品
// @Schemes
// @Description 删除产品
// @Tags product
// @Param id path int true "产品ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[Product] 返回产品信息
// @Router /product/{id}/delete [get]
func noopProductDelete() {}
