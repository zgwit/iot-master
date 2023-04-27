package curd

import (
	"archive/zip"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/v3/pkg/db"
	"io/ioutil"
)

func ApiImport(table string) gin.HandlerFunc {
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
		n, err := db.Engine.Table(table).Insert(datum)
		if err != nil {
			Error(ctx, err)
			return
		}

		OK(ctx, n)
	}
}
