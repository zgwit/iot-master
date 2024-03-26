package protocol

import (
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/v4/api"
	"github.com/zgwit/iot-master/v4/web/curd"
)

func init() {

	api.Register("GET", "/protocol/list", func(ctx *gin.Context) {
		var ps []*Protocol
		for _, p := range protocols {
			ps = append(ps, p)
		}
		curd.OK(ctx, ps)
	})

	api.Register("GET", "/protocol/:name/mappers", func(ctx *gin.Context) {
		name := ctx.Param("name")
		if p, ok := protocols[name]; ok {
			curd.OK(ctx, p.ProductMappers)
		} else {
			curd.Fail(ctx, "协议找不到")
		}
	})

	api.Register("GET", "/protocol/:name/pollers", func(ctx *gin.Context) {
		name := ctx.Param("name")
		if p, ok := protocols[name]; ok {
			curd.OK(ctx, p.ProductPollers)
		} else {
			curd.Fail(ctx, "协议找不到")
		}
	})

	api.Register("GET", "/protocol/:name/options", func(ctx *gin.Context) {
		name := ctx.Param("name")
		if p, ok := protocols[name]; ok {
			curd.OK(ctx, p.TunnelOptions)
		} else {
			curd.Fail(ctx, "协议找不到")
		}
	})
}

// @Summary 协议列表
// @Schemes
// @Description 协议列表
// @Tags protocol
// @Produce json
// @Success 200 {object} curd.ReplyData[Protocol] 返回协议列表
// @Router /protocol/list [get]
func noopProtocolList() {}

// @Summary 协议参数
// @Schemes
// @Description 协议参数
// @Tags protocol
// @Produce json
// @Success 200 {object} curd.ReplyData[[]types.FormItem] 返回协议参数
// @Router /protocol/options [get]
func noopProtocolOptions() {}

// @Summary 协议轮询器
// @Schemes
// @Description 协议轮询器
// @Tags protocol
// @Produce json
// @Success 200 {object} curd.ReplyData[[]types.FormItem] 返回协议轮询器
// @Router /protocol/pollers [get]
func noopProtocolPollers() {}

// @Summary 协议映射
// @Schemes
// @Description 协议映射
// @Tags protocol
// @Produce json
// @Success 200 {object} curd.ReplyData[[]types.FormItem] 返回协议映射
// @Router /protocol/mappers [get]
func noopProtocolMappers() {}
