package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/v3/model"
	"github.com/zgwit/iot-master/v3/pkg/curd"
	"github.com/zgwit/iot-master/v3/pkg/lib"
	"reflect"
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

	app.POST("/count", curd.ApiCount[model.Product]())
	app.POST("/search", curd.ApiSearch[model.Product]())
	app.GET("/list", curd.ApiList[model.Product]())
	app.POST("/create", curd.ApiCreateHook[model.Product](func(m *model.Product) error {
		value := reflect.ValueOf(m).Elem()
		field := value.FieldByName("Id")
		if field.Len() == 0 {
			field.SetString(lib.RandomString(8))
		}
		return isExist("ID已存在", &model.Product{Id: m.Id})
	}, nil))
	app.GET("/:id", curd.ParseParamStringId, curd.ApiGet[model.Product]())
	app.POST("/:id", curd.ParseParamStringId, curd.ApiUpdateHook[model.Product](func(m *model.Product) error {
		return isExist("ID已存在", &model.Product{Id: m.Id})
	}, nil,
		"id", "name", "version", "desc", "properties", "functions", "events", "parameters", "validators", "aggregators"))
	app.GET("/:id/delete", curd.ParseParamStringId, curd.ApiDeleteHook[model.Product](nil, nil))
	app.GET("/export", curd.ApiExport("product", "产品"))
	app.POST("/import", curd.ApiImport("product"))
}
