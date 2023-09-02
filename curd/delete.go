package curd

import (
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/v4/db"
)

func ApiDelete[T any]() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.MustGet("id")

		var data T
		_, err := db.Engine.ID(id).Delete(&data)
		if err != nil {
			Error(ctx, err)
			return
		}

		OK(ctx, nil)
	}
}

func ApiDeleteHook[T any](before, after func(id any) error) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.MustGet("id")
		if before != nil {
			if err := before(id); err != nil {
				Error(ctx, err)
				return
			}
		}

		var data T
		_, err := db.Engine.ID(id).Delete(&data)
		if err != nil {
			Error(ctx, err)
			return
		}

		if after != nil {
			if err := after(&data); err != nil {
				Error(ctx, err)
				return
			}
		}

		OK(ctx, nil)
	}
}
