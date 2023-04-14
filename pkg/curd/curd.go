package curd

import (
	"archive/zip"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/v3/pkg/db"
	"github.com/zgwit/iot-master/v3/pkg/log"
	"io/ioutil"
	"reflect"
)

type Hook func(value interface{}) error

func ApiCount[T any]() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var body ParamSearch
		err := ctx.ShouldBindJSON(&body)
		if err != nil {
			Error(ctx, err)
			return
		}

		query := body.ToQuery()

		var d T
		cnt, err := query.Count(d)
		if err != nil {
			Error(ctx, err)
			return
		}

		OK(ctx, cnt)
	}
}

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

func ApiCreate[T any](before, after func(m *T) error) gin.HandlerFunc {
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

func ApiModify[T any](before, after func(m *T) error, fields ...string) gin.HandlerFunc {
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

func ApiDelete[T any](before, after func(id any) error) gin.HandlerFunc {
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

func ApiDisable[T any](disable bool, before, after func(id any) error) gin.HandlerFunc {
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
		value := reflect.ValueOf(data).Elem()
		field := value.FieldByName("Disabled")
		field.SetBool(disable)

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

func ApiExport[T any](filename string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//id := ctx.MustGet("id")
		ids := ctx.QueryArray("id")

		var data T
		var datum []map[string]any

		query := db.Engine.Table(data)
		if ids != nil && len(ids) > 0 {
			query = query.In("id", ids)
		}
		err := query.Find(&datum)
		if err != nil {
			Error(ctx, err)
			return
		}

		//下载头
		ctx.Header("Content-Type", "application/octet-stream")
		ctx.Header("Content-Disposition", "attachment; filename="+filename+".zip") // 用来指定下载下来的文件名
		ctx.Header("Content-Transfer-Encoding", "binary")

		writer := zip.NewWriter(ctx.Writer)

		for _, data := range datum {
			id := data["id"]
			fn := fmt.Sprintf("%v.json", id)
			f, err := writer.Create(fn)
			if err != nil {
				return
			}

			buf, _ := json.Marshal(data)
			_, err = f.Write(buf)
			if err != nil {
				return
			}
		}

		err = writer.Close()
		if err != nil {
			return
		}
	}
}

func ApiImport[T any]() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		formFile, err := ctx.FormFile("file")
		if err != nil {
			Error(ctx, err)
			return
		}

		file, err := formFile.Open()
		if err != nil {
			Error(ctx, err)
			return
		}
		defer file.Close()

		reader, err := zip.NewReader(file, formFile.Size)
		if err != nil {
			Error(ctx, err)
			return
		}

		//数据解析
		var datum []map[string]any
		for _, file := range reader.File {
			if file.FileInfo().IsDir() {
				continue
			}

			reader, err := file.Open()
			buf, err := ioutil.ReadAll(reader)
			if err != nil {
				Error(ctx, err)
				return
			}

			var data map[string]any
			err = json.Unmarshal(buf, &data)
			if err != nil {
				Error(ctx, err)
				return
			}

			datum = append(datum, data)
		}

		//插入数据
		var data T
		n, err := db.Engine.Table(data).InsertMulti(datum)
		if err != nil {
			Error(ctx, err)
			return
		}

		OK(ctx, n)
	}
}
