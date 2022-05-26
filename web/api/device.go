package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/influx"
	"github.com/zgwit/iot-master/master"
	"github.com/zgwit/iot-master/model"
	"github.com/zgwit/iot-master/tsdb"
	"golang.org/x/net/websocket"
	"regexp"
	"strconv"
	"time"
)

func deviceList(ctx *gin.Context) {
	devs := make([]*model.DeviceEx, 0) //len(devices)

	var body paramSearchEx
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		replyError(ctx, err)
		return
	}

	query := body.toQuery()

	query.Join("LEFT", "element", "device.element_id=element.id")
	query.Join("LEFT", "link", "device.link_id=link.id")

	cnt, err := query.FindAndCount(devs)
	if err != nil {
		replyError(ctx, err)
		return
	}
	replyList(ctx, devs, cnt)

	//dev.Element = element.Name
	//dev.DeviceContent = element.DeviceContent
	//dev.Link = link.Name
	//dev.Link = link.SN
	//dev.Link = link.Remote
}

func afterDeviceCreate(data interface{}) error {
	device := data.(*model.Device)
	//启动
	_, err := master.LoadDevice(device.Id)
	return err
}

func deviceDetail(ctx *gin.Context) {
	var device model.Device

	//补充信息
	dev := model.DeviceEx{Device: device}
	d := master.GetDevice(dev.Id)
	if d != nil {
		dev.Running = d.Running()
	}

	replyOk(ctx, dev)
}

func afterDeviceUpdate(data interface{}) error {
	device := data.(*model.Device)
	//重新启动
	_ = master.RemoveDevice(device.Id)
	_, err := master.LoadDevice(device.Id)
	return err
}

func afterDeviceDelete(data interface{}) error {
	device := data.(*model.Device)
	//关闭
	return master.RemoveDevice(device.Id)
}

func deviceStart(ctx *gin.Context) {
	device := master.GetDevice(ctx.GetInt64("id"))
	if device == nil {
		replyFail(ctx, "not found")
		return
	}
	err := device.Start()
	if err != nil {
		replyError(ctx, err)
		return
	}

	replyOk(ctx, nil)
}

func deviceStop(ctx *gin.Context) {
	device := master.GetDevice(ctx.GetInt64("id"))
	if device == nil {
		replyFail(ctx, "not found")
		return
	}
	err := device.Stop()
	if err != nil {
		replyError(ctx, err)
		return
	}

	replyOk(ctx, nil)
}

func afterDeviceEnable(data interface{}) error {
	device := data.(*model.Device)
	_, err := master.LoadDevice(device.Id)
	return err
}

func afterDeviceDisable(data interface{}) error {
	device := data.(*model.Device)
	dev := master.GetDevice(device.Id)
	if dev == nil {
		return nil
	}
	return dev.Stop()
}

func deviceContext(ctx *gin.Context) {
	device := master.GetDevice(ctx.GetInt64("id"))
	if device == nil {
		replyFail(ctx, "找不到设备")
		return
	}
	replyOk(ctx, device.Context)
}

func deviceRefresh(ctx *gin.Context) {
	device := master.GetDevice(ctx.GetInt64("id"))
	if device == nil {
		replyFail(ctx, "找不到设备")
		return
	}
	err := device.Refresh()
	if err != nil {
		replyError(ctx, err)
		return
	}
	replyOk(ctx, nil)
}

func deviceRefreshPoint(ctx *gin.Context) {
	device := master.GetDevice(ctx.GetInt64("id"))
	if device == nil {
		replyFail(ctx, "找不到设备")
		return
	}
	val, err := device.RefreshPoint(ctx.Param("name"))
	if err != nil {
		replyError(ctx, err)
		return
	}
	replyOk(ctx, val)
}

type executeBody struct {
	Command   string    `json:"command"`
	Arguments []float64 `json:"arguments"`
}

func deviceExecute(ctx *gin.Context) {
	var body executeBody
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		replyError(ctx, err)
		return
	}

	device := master.GetDevice(ctx.GetInt64("id"))
	if device == nil {
		replyFail(ctx, "找不到设备")
		return
	}
	err = device.Execute(body.Command, body.Arguments)
	if err != nil {
		replyError(ctx, err)
		return
	}
	replyOk(ctx, nil)
}

func deviceWatch(ctx *gin.Context) {
	device := master.GetDevice(ctx.GetInt64("id"))
	if device == nil {
		replyFail(ctx, "找不到设备")
		return
	}
	websocket.Handler(func(ws *websocket.Conn) {
		watchAllEvents(ws, device)
	}).ServeHTTP(ctx.Writer, ctx.Request)
}

var timeReg *regexp.Regexp

func init() {
	timeReg = regexp.MustCompile(`^(-?\d+)(h|m|s)$`)
}

func parseTime(tm string) (int64, error) {
	ss := timeReg.FindStringSubmatch(tm)
	if ss == nil || len(ss) != 3 {
		return 0, errors.New("错误时间")
	}
	val, _ := strconv.ParseInt(ss[1], 10, 64)
	switch ss[2] {
	case "d":
		val *= 24 * 60 * 60 * 1000
	case "h":
		val *= 60 * 60 * 1000
	case "m":
		val *= 60 * 1000
	case "s":
		val *= 1000
	}
	return val, nil
}

func deviceValueHistory(ctx *gin.Context) {
	id := ctx.Param("id")
	key := ctx.Param("name")
	start := ctx.DefaultQuery("start", "-5h")
	end := ctx.DefaultQuery("end", "0h")
	window := ctx.DefaultQuery("window", "10m")

	//优先查询InfluxDB
	if influx.Opened() {
		values, err := influx.Query(map[string]string{"id": id}, key, start, end, window)
		if err != nil {
			replyError(ctx, err)
			return
		}
		replyOk(ctx, values)
		return
	}

	//查询内部数据库
	if tsdb.Opened() {
		//相对时间转化为时间戳
		s, err := parseTime(start)
		if err != nil {
			replyError(ctx, err)
			return
		}
		s += time.Now().UnixMilli()

		e, err := parseTime(end)
		if err != nil {
			replyError(ctx, err)
			return
		}
		e += time.Now().UnixMilli()

		w, err := parseTime(window)
		if err != nil {
			replyError(ctx, err)
			return
		}
		values, err := tsdb.Query(id, key, s, e, w)
		if err != nil {
			replyError(ctx, err)
			return
		}
		replyOk(ctx, values)
		return
	}

	replyFail(ctx, "没有开启历史数据库")
}
