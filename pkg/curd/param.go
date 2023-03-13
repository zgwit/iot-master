package curd

import (
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/v3/pkg/db"
	"reflect"
	"xorm.io/builder"
	"xorm.io/xorm"
)

type ParamSearch struct {
	Skip     int                    `form:"skip" json:"skip"`
	Limit    int                    `form:"limit" json:"limit"`
	Sort     map[string]int         `form:"sort" json:"sort"`
	Filters  map[string]interface{} `form:"filter" json:"filter"`
	Keywords map[string]string      `form:"keyword" json:"keyword"`
}

func (body *ParamSearch) ToQuery() *xorm.Session {
	if body.Limit < 1 {
		body.Limit = 20
	}
	op := db.Engine.Limit(body.Limit, body.Skip)

	for k, v := range body.Filters {
		if reflect.TypeOf(v).Kind() == reflect.Slice {
			ll := len(v.([]interface{}))
			if ll > 0 {
				if ll == 1 {
					op.And(k+"=?", v.([]interface{})[0])
				} else {
					op.In(k, v)
				}
			}
		} else {
			if v != nil {
				op.And(k+"=?", v)
			}
		}
	}

	//builder.Or(builder.Like{})
	if len(body.Keywords) > 0 {
		likes := make([]builder.Cond, 0)
		for k, v := range body.Keywords {
			if v != "" {
				//op.And(k+" like ?", "%"+v+"%")
				likes = append(likes, &builder.Like{k, v})
			}
		}
		if len(likes) > 0 {
			op.And(builder.Or(likes...))
		}
	}

	if len(body.Sort) > 0 {
		for k, v := range body.Sort {
			if v > 0 {
				op.Asc(k)
			} else {
				op.Desc(k)
			}
		}
	} else {
		op.Desc("id")
	}

	return op
}

type ParamId struct {
	Id int64 `uri:"id"`
}
type ParamStringId struct {
	Id string `uri:"id"`
}

func ParseParamId(ctx *gin.Context) {
	var pid ParamId
	err := ctx.ShouldBindUri(&pid)
	if err != nil {
		Error(ctx, err)
		ctx.Abort()
		return
	}
	ctx.Set("id", pid.Id)
	ctx.Next()
}

func ParseParamStringId(ctx *gin.Context) {
	var pid ParamStringId
	err := ctx.ShouldBindUri(&pid)
	if err != nil {
		Error(ctx, err)
		ctx.Abort()
		return
	}
	ctx.Set("id", pid.Id)
	ctx.Next()
}

type ParamList struct {
	Skip  int `form:"skip" json:"skip"`
	Limit int `form:"limit" json:"limit"`
}

func (body *ParamList) ToQuery() *xorm.Session {
	if body.Limit < 1 {
		body.Limit = 20
	}
	op := db.Engine.Limit(body.Limit, body.Skip)
	op.Desc("id")
	return op
}
