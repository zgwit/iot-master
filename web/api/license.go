package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zgwit/go-license"
	"github.com/zgwit/iot-master/active"
)

func licenseDetail(ctx *gin.Context) {
	replyOk(ctx, active.Licence())
}

type licenseObj struct {
	License string `json:"license"`
}

func licenseUpdate(ctx *gin.Context) {
	var obj licenseObj
	err := ctx.BindJSON(&obj)
	if err != nil {
		replyError(ctx, err)
		return
	}

	var lic license.Licence
	err = lic.Decode(obj.License)
	if err != nil {
		replyError(ctx, err)
		return
	}

	err = active.Validate(&lic)
	if err != nil {
		replyError(ctx, err)
		return
	}

	err = active.Save(&lic)
	if err != nil {
		replyError(ctx, err)
		return
	}

	replyOk(ctx, nil)
}
