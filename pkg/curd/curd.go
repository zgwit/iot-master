package curd

import (
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/v3/pkg/db"
	"github.com/zgwit/iot-master/v3/pkg/log"
	"reflect"
)

type Hook func(value interface{}) error

func ApiSearch[T any](fields ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var body ParamSearch
		err := ctx.ShouldBindJSON(&body)
		if err != nil {
			Error(ctx, err)
			return
		}

		query := body.ToQuery()

		//查询字段
		fs := ctx.QueryArray("field")
		if len(fs) > 0 {
			query.Cols(fs...)
		} else if len(fields) > 0 {
			query.Cols(fields...)
		}

		var datum []T
		cnt, err := query.FindAndCount(&datum)
		if err != nil {
			Error(ctx, err)
			return
		}

		//OK(ctx, cs)
		List(ctx, datum, cnt)
	}
}

func ApiListWithId[T any](field string, fields ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var body ParamList
		err := ctx.ShouldBindQuery(&body)
		if err != nil {
			Error(ctx, err)
			return
		}

		query := body.ToQuery()

		//查询字段
		fs := ctx.QueryArray("field")
		if len(fs) > 0 {
			query.Cols(fs...)
		} else if len(fields) > 0 {
			query.Cols(fields...)
		}

		//添加条件
		id := ctx.MustGet("id")
		query.Where(field+"=?", id)

		var datum []T
		cnt, err := query.FindAndCount(&datum)
		if err != nil {
			Error(ctx, err)
			return
		}

		//OK(ctx, cs)
		List(ctx, datum, cnt)
	}
}

func ApiList[T any](fields ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var body ParamList
		err := ctx.ShouldBindQuery(&body)
		if err != nil {
			Error(ctx, err)
			return
		}

		query := body.ToQuery()

		//查询字段
		fs := ctx.QueryArray("field")
		if len(fs) > 0 {
			query.Cols(fs...)
		} else if len(fields) > 0 {
			query.Cols(fields...)
		}

		var datum []T
		cnt, err := query.FindAndCount(&datum)
		if err != nil {
			Error(ctx, err)
			return
		}

		//OK(ctx, cs)
		List(ctx, datum, cnt)
	}
}

func ApiCreate[T any](before, after Hook) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var data T
		err := ctx.ShouldBindJSON(&data)
		if err != nil {
			Error(ctx, err)
			return
		}

		if before != nil {
			if err := before(&data); err != nil {
				Error(ctx, err)
				return
			}
		}

		_, err = db.Engine.InsertOne(&data)
		if err != nil {
			Error(ctx, err)
			return
		}

		if after != nil {
			go func() {
				err := after(&data)
				if err != nil {
					log.Error(err)
				}
			}()
		}

		OK(ctx, &data)
	}
}

func ApiModify[T any](before, after Hook, fields ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var data T
		err := ctx.ShouldBindJSON(&data)
		//写入ID
		id := ctx.MustGet("id")
		//value.Elem().FieldByName("Id").Set(reflect.ValueOf(id))

		if err != nil {
			Error(ctx, err)
			return
		}

		if before != nil {
			if err := before(&data); err != nil {
				Error(ctx, err)
				return
			}
		}

		_, err = db.Engine.ID(id).Cols(fields...).Update(&data)
		if err != nil {
			Error(ctx, err)
			return
		}

		if after != nil {
			go func() {
				err := after(&data)
				if err != nil {
					log.Error(err)
				}
			}()
		}

		OK(ctx, &data)
	}
}

func ApiDelete[T any](before, after Hook) gin.HandlerFunc {
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
			go func() {
				err := after(id)
				if err != nil {
					log.Error(err)
				}
			}()
		}

		OK(ctx, nil)
	}
}

func ApiGet[T any](fields ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.MustGet("id")

		query := db.Engine.ID(id)
		//查询字段
		fs := ctx.QueryArray("field")
		if len(fs) > 0 {
			query.Cols(fs...)
		} else if len(fields) > 0 {
			query.Cols(fields...)
		}

		var data T
		has, err := query.Get(&data)
		if err != nil {
			Error(ctx, err)
			return
		} else if !has {
			Fail(ctx, "记录不存在")
			return
		}
		OK(ctx, &data)
	}
}

func ApiDisable[T any](disable bool, before, after Hook) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//id := ctx.GetInt64("id")
		id := ctx.MustGet("id")
		if before != nil {
			if err := before(id); err != nil {
				Error(ctx, err)
				return
			}
		}

		//value := reflect.New(mod)
		//value.Elem().FieldByName("Disabled").SetBool(disable)
		//data := value.Interface()
		var data T
		reflect.ValueOf(data).FieldByName("Disabled").SetBool(disable)

		_, err := db.Engine.ID(id).Cols("disabled").Update(&data)
		if err != nil {
			Error(ctx, err)
			return
		}

		if after != nil {
			go func() {
				err := after(id)
				if err != nil {
					log.Error(err)
				}
			}()
		}

		OK(ctx, nil)
	}
}
