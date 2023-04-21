package curd

import (
	"github.com/gin-gonic/gin"
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

func ApiSearchHook[T any](after func(datum []T) error, fields ...string) gin.HandlerFunc {
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

func ApiSearchWith[T any](table string, join []Join, fields ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var body ParamSearch
		err := ctx.ShouldBindJSON(&body)
		if err != nil {
			Error(ctx, err)
			return
		}

		query := body.ToQuery()

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

		//var data T
		var datum []map[string]any
		session := query.Table(table)

		//补充字段
		for _, j := range join {
			s = append(s, j.Table+"."+j.Field+" as "+j.As)
		}
		session.Select(strings.Join(s, ","))

		//连接查询
		for _, j := range join {
			session.Join("LEFT OUTER", j.Table, j.Table+"."+j.ForeignField+"="+table+"."+j.LocaleField)
		}

		cnt, err := session.FindAndCount(&datum)
		if err != nil {
			Error(ctx, err)
			return
		}

		//OK(ctx, cs)
		List(ctx, datum, cnt)
	}
}
