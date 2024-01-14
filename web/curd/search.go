package curd

import (
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/v4/db"
	"strings"
)

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

		var datum []*T
		cnt, err := query.FindAndCount(&datum)
		if err != nil {
			Error(ctx, err)
			return
		}

		//OK(ctx, cs)
		List(ctx, datum, cnt)
	}
}

func ApiSearchHook[T any](after func(datum []*T) error, fields ...string) gin.HandlerFunc {
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

		var datum []*T
		cnt, err := query.FindAndCount(&datum)
		if err != nil {
			Error(ctx, err)
			return
		}

		if after != nil {
			if err := after(datum); err != nil {
				Error(ctx, err)
				return
			}
		}

		//OK(ctx, cs)
		List(ctx, datum, cnt)
	}
}

func ApiSearchMapHook[T any](after func(datum []map[string]any) error, fields ...string) gin.HandlerFunc {
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

		var data T
		var datum []map[string]any
		cnt, err := query.Table(data).FindAndCount(&datum)
		if err != nil {
			Error(ctx, err)
			return
		}

		//后续处理
		if after != nil {
			err := after(datum)
			if err != nil {
				Error(ctx, err)
				return
			}
		}

		//OK(ctx, cs)
		List(ctx, datum, cnt)
	}
}

type Join struct {
	Table        string
	LocaleField  string
	ForeignField string
	Field        string
	As           string
}

func ApiSearchWith[T any](join []*Join, fields ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var body ParamSearch
		err := ctx.ShouldBindJSON(&body)
		if err != nil {
			Error(ctx, err)
			return
		}

		table := db.Engine.TableName(new(T))
		query := body.ToJoinQuery(table)

		var s []string
		//查询字段
		fs := ctx.QueryArray("field")
		if len(fs) > 0 {
			for _, f := range fs {
				s = append(s, table+"."+f)
			}
		} else if len(fields) > 0 {
			for _, f := range fields {
				s = append(s, table+"."+f)
			}
		} else {
			s = append(s, table+".*")
		}

		var datum []*T
		//var datum []map[string]any
		//session := query.Table(table)

		//补充字段
		for _, j := range join {
			s = append(s, j.Table+"."+db.Engine.Quote(j.Field)+" as "+db.Engine.Quote(j.As))
		}
		query.Select(strings.Join(s, ","))

		//连接查询
		for _, j := range join {
			query.Join("LEFT OUTER", j.Table,
				j.Table+"."+db.Engine.Quote(j.ForeignField)+"="+
					table+"."+db.Engine.Quote(j.LocaleField))
		}

		cnt, err := query.FindAndCount(&datum)
		if err != nil {
			Error(ctx, err)
			return
		}

		//OK(ctx, cs)
		List(ctx, datum, cnt)
	}
}

func ApiSearchWithHook[T any](join []*Join, after func(datum []*T) error, fields ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var body ParamSearch
		err := ctx.ShouldBindJSON(&body)
		if err != nil {
			Error(ctx, err)
			return
		}

		table := db.Engine.TableName(new(T))
		query := body.ToJoinQuery(table)

		var s []string
		//查询字段
		fs := ctx.QueryArray("field")
		if len(fs) > 0 {
			for _, f := range fs {
				s = append(s, table+"."+f)

			}
		} else if len(fields) > 0 {
			for _, f := range fields {
				s = append(s, table+"."+f)
			}
		} else {
			s = append(s, table+".*")
		}

		var datum []*T
		//var datum []map[string]any
		//session := query.Table(table)

		//补充字段
		for _, j := range join {
			s = append(s, j.Table+"."+db.Engine.Quote(j.Field)+" as "+db.Engine.Quote(j.As))
		}
		query.Select(strings.Join(s, ","))

		//连接查询
		for _, j := range join {
			query.Join("LEFT OUTER", j.Table,
				j.Table+"."+db.Engine.Quote(j.ForeignField)+"="+
					table+"."+db.Engine.Quote(j.LocaleField))
		}

		cnt, err := query.FindAndCount(&datum)
		if err != nil {
			Error(ctx, err)
			return
		}

		//后续处理
		if after != nil {
			if err := after(datum); err != nil {
				Error(ctx, err)
				return
			}
		}

		//OK(ctx, cs)
		List(ctx, datum, cnt)
	}
}
