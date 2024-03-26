package tunnel

import (
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/v4/api"
	"github.com/zgwit/iot-master/v4/pkg/db"
	"github.com/zgwit/iot-master/v4/web/curd"
)

func init() {

	api.Register("GET", "/device/:id/protocol/options", func(ctx *gin.Context) {
		id := ctx.Param("id")
		
		var cols []string
		for _, o := range p.DeviceOptions {
			cols = append(cols, o.Key)
		}

		var opts map[string]any
		_, err = db.Engine.ID(id).Cols(cols...).Get(&opts)
		if err != nil {
			curd.Error(ctx, err)
			return
		}

		curd.OK(ctx, opts)
	})

	api.Register("POST", "/device/:id/protocol/options", func(ctx *gin.Context) {
		id := ctx.Param("id")

		var opts map[string]any
		err := ctx.BindJSON(&opts)
		if err != nil {
			curd.Error(ctx, err)
			return
		}

		var cols []string
		for k, _ := range opts {
			cols = append(cols, k)
		}

		_, err = db.Engine.ID(id).Cols(cols...).Update(&opts)

		curd.OK(ctx, nil)
	})

}
