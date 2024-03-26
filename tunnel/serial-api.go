package tunnel

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/v4/api"
	"github.com/zgwit/iot-master/v4/pkg/db"
	"github.com/zgwit/iot-master/v4/pkg/log"
	"github.com/zgwit/iot-master/v4/web/curd"
	"go.bug.st/serial"
)

func init() {

	api.Register("POST", "/serial/count", curd.ApiCount[Serial]())

	api.Register("POST", "/serial/search", curd.ApiSearchHook[Serial](func(serials []*Serial) error {
		for k, ser := range serials {
			c := GetSerial(ser.Id)
			if c != nil {
				serials[k].Running = c.Running
			}
		}
		return nil
	}))

	api.Register("GET", "/serial/list", curd.ApiList[Serial]())

	api.Register("POST", "/serial/create", curd.ApiCreateHook[Serial](curd.GenerateID[Serial](), func(value *Serial) error {
		return LoadSerial(value)
	}))

	api.Register("GET", "/serial/:id", curd.ParseParamStringId, curd.ApiGetHook[Serial](func(ser *Serial) error {
		c := GetSerial(ser.Id)
		if c != nil {
			ser.Running = c.Running
		}
		return nil
	}))

	api.Register("POST", "/serial/:id", curd.ParseParamStringId, curd.ApiUpdateHook[Serial](nil, func(value *Serial) error {
		c := GetSerial(value.Id)
		err := c.Close()
		if err != nil {
			log.Error(err)
		}
		return LoadSerial(value)
	},
		"id", "name", "description", "heartbeat", "poller_period", "poller_interval", "protocol_name", "protocol_options",
		"port_name", "baud_rate", "data_bits", "stop_bits", "parity_mode", "retry_timeout", "retry_maximum", "disabled"))

	api.Register("GET", "/serial/:id/delete", curd.ParseParamStringId, curd.ApiDeleteHook[Serial](nil, func(value interface{}) error {
		id := value.(string)
		c := GetSerial(id)
		return c.Close()
	}))

	api.Register("GET", "/serial/:id/disable", curd.ParseParamStringId, curd.ApiDisableHook[Serial](true, nil, func(value interface{}) error {
		id := value.(string)
		c := GetSerial(id)
		return c.Close()
	}))

	api.Register("GET", "/serial/:id/enable", curd.ParseParamStringId, curd.ApiDisableHook[Serial](false, nil, func(value interface{}) error {
		id := value.(string)
		var m Serial
		has, err := db.Engine.ID(id).Get(&m)
		if err != nil {
			return err
		}
		if !has {
			return fmt.Errorf("找不到 %s", id)
		}
		return LoadSerial(&m)
	}))

	api.Register("GET", "/serial/:id/start", curd.ParseParamStringId, func(ctx *gin.Context) {
		id := ctx.GetString("id")
		c := GetSerial(id)
		if c == nil {
			curd.Fail(ctx, "找不到连接")
			return
		}
		err := c.Open()
		if err != nil {
			curd.Error(ctx, err)
			return
		}
		curd.OK(ctx, nil)
	})

	api.Register("GET", "/serial/:id/stop", curd.ParseParamStringId, func(ctx *gin.Context) {
		id := ctx.GetString("id")
		c := GetSerial(id)
		if c == nil {
			curd.Fail(ctx, "找不到连接")
			return
		}
		err := c.Close()
		if err != nil {
			curd.Error(ctx, err)
			return
		}
		curd.OK(ctx, nil)
	})

	api.Register("GET", "/serial/ports", func(ctx *gin.Context) {
		list, err := serial.GetPortsList()
		if err != nil {
			curd.Error(ctx, err)
			return
		}
		curd.OK(ctx, list)
	})

}

// @Summary 查询串口数量
// @Schemes
// @Description 查询串口数量
// @Tags serial
// @Param search body curd.ParamSearch true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[int64] 返回串口数量
// @Router /serial/count [post]
func noopSerialCount() {}

// @Summary 查询串口
// @Schemes
// @Description 查询串口
// @Tags serial
// @Param search body curd.ParamSearch true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyList[Serial] 返回串口信息
// @Router /serial/search [post]
func noopSerialSearch() {}

// @Summary 查询串口
// @Schemes
// @Description 查询串口
// @Tags serial
// @Param search query curd.ParamList true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyList[Serial] 返回串口信息
// @Router /serial/list [get]
func noopSerialList() {}

// @Summary 创建串口
// @Schemes
// @Description 创建串口
// @Tags serial
// @Param search body Serial true "串口信息"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[Serial] 返回串口信息
// @Router /serial/create [post]
func noopSerialCreate() {}

// @Summary 修改串口
// @Schemes
// @Description 修改串口
// @Tags serial
// @Param id path int true "串口ID"
// @Param serial body Serial true "串口信息"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[Serial] 返回串口信息
// @Router /serial/{id} [post]
func noopSerialUpdate() {}

// @Summary 获取串口
// @Schemes
// @Description 获取串口
// @Tags serial
// @Param id path int true "串口ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[Serial] 返回串口信息
// @Router /serial/{id} [get]
func noopSerialGet() {}

// @Summary 删除串口
// @Schemes
// @Description 删除串口
// @Tags serial
// @Param id path int true "串口ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[Serial] 返回串口信息
// @Router /serial/{id}/delete [get]
func noopSerialDelete() {}

// @Summary 启用串口
// @Schemes
// @Description 启用串口
// @Tags serial
// @Param id path int true "串口ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[Serial] 返回串口信息
// @Router /serial/{id}/enable [get]
func noopSerialEnable() {}

// @Summary 禁用串口
// @Schemes
// @Description 禁用串口
// @Tags serial
// @Param id path int true "串口ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[Serial] 返回串口信息
// @Router /serial/{id}/disable [get]
func noopSerialDisable() {}

// @Summary 启动串口
// @Schemes
// @Description 启动串口
// @Tags serial
// @Param id path int true "串口ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[Serial] 返回串口信息
// @Router /serial/{id}/start [get]
func noopSerialStart() {}

// @Summary 停止串口
// @Schemes
// @Description 停止串口
// @Tags serial
// @Param id path int true "串口ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[Serial] 返回串口信息
// @Router /serial/{id}/stop [get]
func noopSerialStop() {}

// @Summary 串口列表
// @Schemes
// @Description 串口列表
// @Tags serial
// @Produce json
// @Success 200 {object} curd.ReplyData[string] 返回串口列表
// @Router /serial/ports [get]
func noopSerialPorts() {}
