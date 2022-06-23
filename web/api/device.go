package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/db"
	"github.com/zgwit/iot-master/history"
	"github.com/zgwit/iot-master/master"
	"github.com/zgwit/iot-master/model"
	"golang.org/x/net/websocket"
)

func deviceList(ctx *gin.Context) {
	var devs []*model.DeviceEx

	var body paramSearchEx
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		replyError(ctx, err)
		return
	}

	query := body.toQuery()

	query.Table("device")
	query.Select("device.*, " + //TODO 只返回需要的字段
		" 0 as running, product.name as product, tunnel.name as tunnel")
	query.Join("LEFT", "product", "device.product_id=product.id")
	query.Join("LEFT", "tunnel", "device.tunnel_id=tunnel.id")

	//err = query.Find(devs)
	cnt, err := query.FindAndCount(&devs)
	if err != nil {
		replyError(ctx, err)
		return
	}

	//补充running状态
	for _, dev := range devs {
		d := master.GetDevice(dev.Id)
		if d != nil {
			dev.Running = d.Running()
		}
	}

	replyList(ctx, devs, cnt)
}

func afterDeviceCreate(data interface{}) error {
	device := data.(*model.Device)
	//启动
	dev, err := master.LoadDevice(device.Id)
	if err == nil {
		err = dev.Start()
	}
	return err
}

func deviceDetail(ctx *gin.Context) {
	var device model.DeviceEx
	has, err := db.Engine.ID(ctx.GetInt64("id")).Get(&device.Device)
	if err != nil {
		replyError(ctx, err)
		return
	}
	if !has {
		replyFail(ctx, "设备不存在")
		return
	}

	if device.ProductId != "" {
		var template model.Product
		has, err := db.Engine.ID(device.ProductId).Get(&template)
		if has && err == nil {
			device.Product = template.Name
			device.DeviceContent = template.DeviceContent
		}
	}

	d := master.GetDevice(device.Id)
	if d != nil {
		device.Running = d.Running()
	}

	replyOk(ctx, device)
}

func afterDeviceUpdate(data interface{}) error {
	device := data.(*model.Device)
	//重新启动
	_ = master.RemoveDevice(device.Id)
	dev, err := master.LoadDevice(device.Id)
	if err == nil {
		err = dev.Start()
	}
	return err
}

func afterDeviceDelete(id interface{}) error {
	return master.RemoveDevice(id.(int64))
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

func afterDeviceEnable(id interface{}) error {
	_ = master.RemoveDevice(id.(int64))
	dev, err := master.LoadDevice(id.(int64))
	if err != nil {
		err = dev.Start()
	}
	return err
}

func afterDeviceDisable(id interface{}) error {
	return master.RemoveDevice(id.(int64))
}

func deviceContext(ctx *gin.Context) {
	device := master.GetDevice(ctx.GetInt64("id"))
	if device == nil {
		replyFail(ctx, "找不到设备")
		return
	}
	replyOk(ctx, device.Context)
}

func deviceContextUpdate(ctx *gin.Context) {
	var values map[string]interface{}
	err := ctx.ShouldBindJSON(values)
	if err != nil {
		replyError(ctx, err)
		return
	}

	device := master.GetDevice(ctx.GetInt64("id"))
	if device == nil {
		replyFail(ctx, "找不到设备")
		return
	}

	for k, v := range values {
		err := device.Set(k, v)
		if err != nil {
			replyError(ctx, err)
			return
		}
	}

	replyOk(ctx, nil)
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
	replyOk(ctx, device.Context)
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
	Command   string        `json:"command"`
	Arguments []interface{} `json:"arguments"`
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

func deviceValueHistory(ctx *gin.Context) {
	key := ctx.Param("name")
	start := ctx.DefaultQuery("start", "-5h")
	end := ctx.DefaultQuery("end", "0h")
	window := ctx.DefaultQuery("window", "10m")

	if history.Storage == nil {
		replyFail(ctx, "未开启历史存储")
		return
	}

	values, err := history.Storage.Query(ctx.GetInt64("id"), key, start, end, window)
	if err != nil {
		replyError(ctx, err)
		return
	}
	replyOk(ctx, values)
	return
}
