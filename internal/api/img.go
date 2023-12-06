package api

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/v4/pkg/web/curd"
	"os"
	"regexp"
	"time"
)

// @Summary 上传图片
// @Schemes
// @Description 可以上传多个图片
// @Tags project
// @Param img body file true "图片"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[string] 返回图片信息
// @Router /img/upload [post]
func noopImgCreate() {}

func imgRouter(app *gin.RouterGroup) {
	app.POST("/create", imgSave)
}
func imgSave(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		curd.Error(c, errors.New("图片接收失败"))
		return
	}
	imgArr := form.File["file"]
	path, err := os.Getwd()
	if err != nil {
		curd.Error(c, errors.New("根路径获取失败"))
		return
	}
	t := time.Now()
	logs := make(map[string]string)
	for i, v := range imgArr {
		fileName, err := reFileName(v.Filename, i+1)
		if err != nil {
			curd.Error(c, err)
			return
		}
		dst := fmt.Sprintf("%s/static/img/%d/%d/%v", path, t.Year(), int(t.Month()), fileName)
		logs[v.Filename] = dst

		if err = c.SaveUploadedFile(v, dst); err != nil {
			curd.Error(c, errors.New("图片存储失败"))
			return
		}
	}
	curd.OK(c, logs)
}
func reFileName(f string, i int) (string, error) {
	format := `\.\w+$`
	re, err := regexp.Compile(format)
	if err != nil {
		return f, errors.New("图片格式正则没有匹配到")
	}
	index := re.FindStringIndex(f)
	if index == nil {
		return f, errors.New("没有找到图片格式的正则下标")
	}
	t := time.Now()
	fileName := fmt.Sprintf("%d%d%d%d%d", t.Day(), t.Hour(), t.Minute(), t.Second(), i)

	f = fileName + f[index[0]:] //加了文件后缀
	return f, nil
}
