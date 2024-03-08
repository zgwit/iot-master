package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/v4/pkg/db"
	"github.com/zgwit/iot-master/v4/types"
	"github.com/zgwit/iot-master/v4/web/curd"
	"xorm.io/xorm/schemas"
)

// @Summary 产品版本列表
// @Schemes
// @Description 产品版本列表
// @Tags product-version
// @Param id path int true "项目ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[[]types.ProductVersion] 返回产品版本信息
// @Router /product/{id}/version [get]
func productVersionList(ctx *gin.Context) {
}

type projectVersion struct {
	Name string `json:"name"`
}

// @Summary 创建产品版本
// @Schemes
// @Description 创建产品版本
// @Tags product-version
// @Param id path int true "项目ID"
// @Param version path int true "设备ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[int]
// @Router /product/{id}/version/create [post]
func productVersionCreate(ctx *gin.Context) {
	var pd types.ProductVersion
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
// @Param id path int true "项目ID"
// @Param version path int true "设备ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[int]
// @Router /product/{id}/version/{version}/delete [get]
func productVersionDelete(ctx *gin.Context) {
	_, err := db.Engine.ID(schemas.PK{ctx.Param("id"), ctx.Param("version")}).Delete(new(types.ProductVersion))
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
// @Param id path int true "项目ID"
// @Param version path int true "设备ID"
// @Param product-version body types.ProductVersion true "产品版本信息"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[int]
// @Router /product/{id}/version/{version} [post]
func productVersionUpdate(ctx *gin.Context) {
	var pd types.ProductVersion
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

func productVersionRouter(app *gin.RouterGroup) {
	app.GET("/list", curd.ParseParamStringId, curd.ApiListWithId[types.ProductVersion]("product_id"))
	app.POST("/:version/create", productVersionCreate)
	app.GET("/:version/delete", productVersionDelete)
	app.POST("/:version", productVersionUpdate)
}
