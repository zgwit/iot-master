package product

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/zgwit/iot-master/v4/api"
	"github.com/zgwit/iot-master/v4/pkg/db"
	"github.com/zgwit/iot-master/v4/web/curd"
	"io"
	"os"
	"path/filepath"
	"xorm.io/xorm/schemas"
)

func init() {
	api.Register("GET", "/product/:id/version/list", curd.ParseParamStringId, curd.ApiListById[ProductVersion]("product_id"))
	api.Register("POST", "/product/:id/version/create", productVersionCreate)
	api.Register("GET", "/product/:id/version/:version/delete", productVersionDelete)
	api.Register("POST", "/product/:id/version/:version", productVersionUpdate)
	api.Register("GET", "/product/:id/version/:version/config/:config", productVersionConfigGet)
	api.Register("POST", "/product/:id/version/:version/config/:config", productVersionConfigSet)
}

// @Summary 产品版本列表
// @Schemes
// @Description 产品版本列表
// @Tags product-version
// @Param id path int true "产品ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[[]ProductVersion] 返回产品版本信息
// @Router /product/{id}/version/list [get]
func productVersionList(ctx *gin.Context) {
}

type projectVersion struct {
	Name string `json:"name"`
}

// @Summary 创建产品版本
// @Schemes
// @Description 创建产品版本
// @Tags product-version
// @Param id path int true "产品ID"
// @Param version path int true "版本ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[int]
// @Router /product/{id}/version/create [post]
func productVersionCreate(ctx *gin.Context) {
	var pd ProductVersion
	err := ctx.ShouldBindJSON(&pd)
	if err != nil {
		curd.Error(ctx, err)
		return
	}

	pd.ProductId = ctx.Param("id")
	_, err = db.Engine.InsertOne(&pd)

	if err != nil {
		curd.Error(ctx, err)
		return
	}
	curd.OK(ctx, nil)
}

// @Summary 删除产品版本
// @Schemes
// @Description 删除产品版本
// @Tags product-version
// @Param id path int true "产品ID"
// @Param version path int true "版本ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[int]
// @Router /product/{id}/version/{version}/delete [get]
func productVersionDelete(ctx *gin.Context) {
	_, err := db.Engine.ID(schemas.PK{ctx.Param("id"), ctx.Param("version")}).Delete(new(ProductVersion))
	if err != nil {
		curd.Error(ctx, err)
		return
	}
	curd.OK(ctx, nil)
}

// @Summary 修改产品版本
// @Schemes
// @Description 修改产品版本
// @Tags product-version
// @Param id path int true "产品ID"
// @Param version path int true "版本ID"
// @Param product-version body ProductVersion true "产品版本信息"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[int]
// @Router /product/{id}/version/{version} [post]
func productVersionUpdate(ctx *gin.Context) {
	var pd ProductVersion
	err := ctx.ShouldBindJSON(&pd)
	if err != nil {
		curd.Error(ctx, err)
		return
	}
	_, err = db.Engine.ID(schemas.PK{ctx.Param("id"), ctx.Param("version")}).Cols("name").Update(&pd)
	if err != nil {
		curd.Error(ctx, err)
		return
	}
	curd.OK(ctx, nil)
}

// @Summary 获得产品配置
// @Schemes
// @Description 获得产品配置
// @Tags product-version
// @Param id path int true "产品ID"
// @Param version path int true "版本ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[any]
// @Router /product/{id}/version/{version}/config/{config} [get]
func productVersionConfigGet(ctx *gin.Context) {
	fn := filepath.Join(viper.GetString("data"), "product", ctx.Param("id"), ctx.Param("version"), ctx.Param("config")+".json")
	buf, err := os.ReadFile(fn)
	if err != nil {
		curd.Error(ctx, err)
		return
	}

	var data any
	err = json.Unmarshal(buf, &data)
	if err != nil {
		curd.Error(ctx, err)
		return
	}

	curd.OK(ctx, data)
}

// @Summary 修改产品配置
// @Schemes
// @Description 修改产品配置
// @Tags product-version
// @Param id path int true "产品ID"
// @Param version path int true "版本ID"
// @Param config body any true "产品版本信息"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[int]
// @Router /product/{id}/version/{version}/config/{config} [post]
func productVersionConfigSet(ctx *gin.Context) {
	dir := filepath.Join(viper.GetString("data"), "product", ctx.Param("id"), ctx.Param("version"))
	_ = os.MkdirAll(dir, os.ModePerm)

	fn := filepath.Join(dir, ctx.Param("config")+".json")
	file, err := os.Create(fn)
	if err != nil {
		curd.Error(ctx, err)
		return
	}
	defer file.Close()

	//写入文件
	_, err = io.Copy(file, ctx.Request.Body)
	if err != nil {
		curd.Error(ctx, err)
		return
	}

	curd.OK(ctx, nil)
}
