package api

import (
	"github.com/gin-gonic/gin"
	"github.com/timshannon/bolthold"
	"reflect"
	"regexp"
)

type paramSearchEx struct {
	Skip     int                    `form:"skip" json:"skip"`
	Limit    int                    `form:"limit" json:"limit"`
	Sort     map[string]int         `form:"sort" json:"sort"`
	Filters  map[string]interface{} `form:"filter" json:"filter"`
	Keywords map[string]string      `form:"keyword" json:"keyword"`
}

func (body *paramSearchEx) toQuery() *bolthold.Query {
	if body.Limit < 1 {
		body.Limit = 20
	}

	op := &bolthold.Query{}
	op.Skip(body.Skip).Limit(body.Limit)

	for k, v := range body.Filters {
		if reflect.TypeOf(v).Kind() == reflect.Slice {
			ll := len(v.([]interface{}))
			if ll > 0 {
				if ll == 1 {
					op.And(k).Eq(v.([]interface{})[0])
				} else {
					op.And(k).In(bolthold.Slice(v))
				}
			}
		} else {
			if v != nil {
				op.And(k).Eq(v)
			}
		}
	}

	for k, v := range body.Keywords {
		if v != "" {
			//op.And(k+" like", "%"+v+"%")
			reg := regexp.MustCompile(v)
			op.And(k).RegExp(reg)
		}
	}

	if len(body.Sort) > 0 {
		for k, v := range body.Sort {
			op.SortBy(k)
			if v < 0 {
				op.Reverse()
			}
		}
	} else {
		//默认ID逆序
		op.SortBy("Id").Reverse()
	}

	return op
}

type paramId struct {
	Id int64 `uri:"id"`
}
type paramStringId struct {
	Id string `uri:"id"`
}

func parseParamId(ctx *gin.Context) {
	var pid paramId
	err := ctx.ShouldBindUri(&pid)
	if err != nil {
		replyError(ctx, err)
		ctx.Abort()
		return
	}
	ctx.Set("id", pid.Id)
	ctx.Next()
}

func parseParamStringId(ctx *gin.Context) {
	var pid paramStringId
	err := ctx.ShouldBindUri(&pid)
	if err != nil {
		replyError(ctx, err)
		ctx.Abort()
		return
	}
	ctx.Set("id", pid.Id)
	ctx.Next()
}
