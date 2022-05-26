package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/db"
	"reflect"
)

type hook func(value interface{}) error

func createSliceFromType(mod reflect.Type) interface{} {
	//datas := reflect.MakeSlice(reflect.SliceOf(mod), 0, 10).Interface()

	//解决不可寻址的问题，参考modern-go/reflect2 safe_slice.go
	val := reflect.MakeSlice(reflect.SliceOf(mod), 0, 1)
	ptr := reflect.New(val.Type())
	ptr.Elem().Set(val)
	return ptr.Interface()
}

func curdApiList(mod reflect.Type) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		datum := createSliceFromType(mod)

		var body paramSearchEx
		err := ctx.ShouldBindJSON(&body)
		if err != nil {
			replyError(ctx, err)
			return
		}

		query := body.toQuery()

		cnt, err := query.FindAndCount(datum)
		if err != nil {
			replyError(ctx, err)
			return
		}

		//replyOk(ctx, cs)
		replyList(ctx, datum, cnt)
	}
}

func curdApiCreate(mod reflect.Type, before, after hook) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		data := reflect.New(mod).Interface()
		err := ctx.ShouldBindJSON(&data)
		if err != nil {
			replyError(ctx, err)
			return
		}

		if before != nil {
			if err := before(data); err != nil {
				replyError(ctx, err)
				return
			}
		}

		_, err = db.Engine.Insert(data)
		if err != nil {
			replyError(ctx, err)
			return
		}

		if after != nil {
			err = after(data)
			if err != nil {
				replyError(ctx, err)
				return
			}
		}

		replyOk(ctx, data)
	}
}

func curdApiModify(mod reflect.Type, updateFields []string, before, after hook) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var pid paramId
		err := ctx.ShouldBindUri(&pid)
		if err != nil {
			replyError(ctx, err)
			return
		}

		data := reflect.New(mod).Interface()
		err = ctx.ShouldBindJSON(&data)
		if err != nil {
			replyError(ctx, err)
			return
		}

		if before != nil {
			if err := before(pid.Id); err != nil {
				replyError(ctx, err)
				return
			}
		}

		_, err = db.Engine.ID(pid.Id).Cols(updateFields...).Update(data)
		if err != nil {
			replyError(ctx, err)
			return
		}

		if after != nil {
			err = after(data)
			if err != nil {
				replyError(ctx, err)
				return
			}
		}

		replyOk(ctx, data)
	}
}

func curdApiDelete(mod reflect.Type, before, after hook) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var pid paramId
		err := ctx.ShouldBindUri(&pid)
		if err != nil {
			replyError(ctx, err)
			return
		}

		if before != nil {
			if err := before(pid.Id); err != nil {
				replyError(ctx, err)
				return
			}
		}

		data := reflect.New(mod).Interface()
		_, err = db.Engine.ID(pid.Id).Delete(data)
		if err != nil {
			replyError(ctx, err)
			return
		}

		if after != nil {
			err = after(data)
			if err != nil {
				replyError(ctx, err)
				return
			}
		}

		replyOk(ctx, nil)
	}
}

func curdApiGet(mod reflect.Type) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var pid paramId
		err := ctx.ShouldBindUri(&pid)
		if err != nil {
			replyError(ctx, err)
			return
		}

		data := reflect.New(mod).Interface()
		has, err := db.Engine.ID(pid.Id).Get(data)
		if err != nil {
			replyError(ctx, err)
			return
		} else if !has {
			replyFail(ctx, "记录不存在")
			return
		}
		replyOk(ctx, data)
	}
}

func curdApiDisable(mod reflect.Type, disable bool, before, after hook) gin.HandlerFunc  {
	return func(ctx *gin.Context) {
		var pid paramId
		err := ctx.ShouldBindUri(&pid)
		if err != nil {
			replyError(ctx, err)
			return
		}

		value := reflect.New(mod)
		value.FieldByName("disabled").SetBool(disable)
		data := value.Interface()
		_, err = db.Engine.ID(pid.Id).Cols("disabled").Update(data)
		if err != nil {
			replyError(ctx, err)
			return
		}

		if after != nil {
			err = after(data)
			if err != nil {
				replyError(ctx, err)
				return
			}
		}

		replyOk(ctx, nil)
	}
}
