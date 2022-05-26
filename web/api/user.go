package api

import (
	"github.com/gin-gonic/gin"
)


func userMe(ctx *gin.Context) {
	replyOk(ctx, ctx.MustGet("user"))
}

func userPassword(ctx *gin.Context) {



	replyOk(ctx, nil)
}
