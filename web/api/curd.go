package api

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/timshannon/bolthold"
	"github.com/zgwit/iot-master/internal/db"
	"github.com/zgwit/iot-master/pkg/log"
	"reflect"
)

type hook func(value interface{}) error

func generateUUID(data interface{}) error {
	value := reflect.ValueOf(data).Elem()
	field := value.FieldByName("Id")
	//使用UUId作为Id
	//field.IsZero() 如果为空串时，生成UUID
	if field.Len() == 0 {
		field.SetString(uuid.NewString())
	}
	return nil
}

func createCurdApiList[T any]() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var body paramSearchEx
		err := ctx.ShouldBindJSON(&body)
		if err != nil {
			replyError(ctx, err)
			return
		}

		query := body.toQuery()

		var datum []T
		cnt, err := db.Store().Count(datum, query)
		if err != nil {
			replyError(ctx, err)
			return
		}

		err = db.Store().Find(&datum, query)
		if err != nil {
			replyError(ctx, err)
			return
		}

		//replyOk(ctx, cs)
		replyList(ctx, datum, cnt)
	}
}

func createCurdApiCreate[T any](before, after hook) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var data T
		err := ctx.ShouldBindJSON(&data)
		if err != nil {
			replyError(ctx, err)
			return
		}

		if before != nil {
			if err := before(&data); err != nil {
				replyError(ctx, err)
				return
			}
		}

		err = db.Store().Insert(bolthold.NextSequence(), &data)
		if err != nil {
			replyError(ctx, err)
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

		replyOk(ctx, &data)
	}
}

func createCurdApiModify[T any](before, after hook) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var data T
		err := ctx.ShouldBindJSON(&data)
		//写入ID
		id := ctx.MustGet("id")
		//value.Elem().FieldByName("Id").Set(reflect.ValueOf(id))

		if err != nil {
			replyError(ctx, err)
			return
		}

		if before != nil {
			if err := before(&data); err != nil {
				replyError(ctx, err)
				return
			}
		}

		err = db.Store().Update(id, &data)
		if err != nil {
			replyError(ctx, err)
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

		replyOk(ctx, &data)
	}
}

func createCurdApiDelete[T any](before, after hook) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.MustGet("id")
		if before != nil {
			if err := before(id); err != nil {
				replyError(ctx, err)
				return
			}
		}

		var data T
		err := db.Store().Delete(id, data)
		if err != nil {
			replyError(ctx, err)
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

		replyOk(ctx, nil)
	}
}

func createCurdApiGet[T any]() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.MustGet("id")
		var data T
		err := db.Store().Get(id, &data)
		if err != nil {
			replyError(ctx, err)
			return
		}
		replyOk(ctx, &data)
	}
}

func createCurdApiDisable[T any](disable bool, before, after hook) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//id := ctx.GetUint64("id")
		id := ctx.MustGet("id")
		if before != nil {
			if err := before(id); err != nil {
				replyError(ctx, err)
				return
			}
		}

		//value := reflect.New(mod)
		//value.Elem().FieldByName("Disabled").SetBool(disable)
		//data := value.Interface()
		var data T

		err := db.Store().Get(id, &data)
		if err != nil {
			replyError(ctx, err)
			return
		}

		reflect.ValueOf(data).FieldByName("Disabled").SetBool(disable)

		err = db.Store().Update(id, &data)
		if err != nil {
			replyError(ctx, err)
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

		replyOk(ctx, nil)
	}
}
