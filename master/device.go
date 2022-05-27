package master

import (
	"errors"
	"fmt"
	"github.com/antonmedv/expr"
	"github.com/zgwit/iot-master/db"
	"github.com/zgwit/iot-master/events"
	"github.com/zgwit/iot-master/influx"
	"github.com/zgwit/iot-master/model"
	"github.com/zgwit/iot-master/protocol"
	"github.com/zgwit/iot-master/tsdb"
	"strconv"
	"time"
)

//Device 设备
type Device struct {
	model.Device
	events.EventEmitter

	Context map[string]interface{}

	pollers []*Poller
	alarms  []*Alarm

	//命令索引
	commandIndex map[string]*model.Command

	mapper *Mapper

	running bool
}

func NewDevice(m *model.Device) (*Device, error) {
	dev := &Device{
		Device:       *m,
		Context:      make(map[string]interface{}),
		commandIndex: make(map[string]*model.Command, 0),
		pollers:      make([]*Poller, 0),
		alarms:       make([]*Alarm, 0),
	}
	var err error

	//加载模板
	if dev.ElementId != "" {
		var element model.Element
		has, err := db.Engine.ID(dev.ElementId).Get(&element)
		if err != nil {
			return nil, err
		}
		if !has {
			return nil, fmt.Errorf("找不到模板 %s", dev.ElementId)
		}
		dev.DeviceContent = element.DeviceContent
	}

	//索引命令
	for _, cmd := range m.Commands {
		dev.commandIndex[cmd.Name] = cmd
	}

	err = dev.initCalculators()
	if err != nil {
		return nil, err
	}

	err = dev.initAlarms()
	if err != nil {
		return nil, err
	}

	return dev, nil
}

func (dev *Device) BindAdapter(adapter protocol.Adapter) error {
	var err error
	if dev.mapper != nil {
		if dev.mapper.adapter == adapter {
			//已经绑定，相同连接，则不用再处理
			return nil
		}
		// else dev.mapper.Off("data")
	}

	dev.mapper, err = newMapper(dev.Station, dev.Points, adapter)
	if err != nil {
		return err
	}

	//metric := strconv.Itoa(dev.Id)
	metric := strconv.FormatInt(dev.Id, 10)

	//处理数据变化结果
	dev.mapper.On("data", func(data map[string]interface{}) {
		//更新上下文
		for k, v := range data {
			dev.Context[k] = v
		}
		//数据变化后，更新计算
		for _, calculator := range dev.Calculators {
			val, err := calculator.Evaluate(dev.Context)
			if err != nil {
				dev.Emit("error", err)
			} else {
				data[calculator.As] = val //也上报和保存
				dev.Context[calculator.As] = val
			}
		}

		//处理策略
		for _, alarm := range dev.alarms {
			err := alarm.Execute(dev.Context)
			if err != nil {
				dev.Emit("error", err)
			}
		}

		//向上广播
		dev.Emit("data", data)

		//保存到时序数据库
		//是否有必要起协程 或者 使用单一进程进行写入
		go func() {
			if tsdb.Opened() {
				_ = tsdb.Write(metric, data)
			}

			if influx.Opened() {
				_ = influx.Write(map[string]string{"id": metric}, data)
			}
		}()
	})

	//关闭之前的轮询
	for _, p := range dev.pollers {
		p.Stop()
	}
	if dev.Pollers == nil {
		return nil
	}
	for _, v := range dev.Pollers {
		addr, _ := adapter.Address(v.Address)
		dev.pollers = append(dev.pollers, &Poller{Poller: *v, Addr: addr, mapper: dev.mapper})
	}

	return err
}

func (dev *Device) initAlarms() error {
	if dev.Alarms == nil {
		return nil
	}
	for _, v := range dev.Alarms {
		a := &Alarm{Alarm: *v}
		a.On("alarm", func(alarm *model.AlarmContent) {
			da := &model.DeviceAlarm{DeviceId: dev.Id, AlarmContent: *alarm}

			//入库
			_, _ = db.Engine.InsertOne(da)
			dev.createEvent("告警：" + alarm.Message)

			//上报
			dev.Emit("alarm", da)
		})
		dev.alarms = append(dev.alarms, a)
	}
	return nil
}

func (dev *Device) initCalculators() error {
	//初始化计算器
	for _, calculator := range dev.Calculators {
		err := calculator.Init()
		if err != nil {
			return err
		}
	}
	return nil
}

func (dev *Device) createEvent(event string) {
	_, _ = db.Engine.InsertOne(model.Event{Target: "device", TargetId: dev.Id, Event: event})
}

//Start 设备启动
func (dev *Device) Start() error {
	dev.createEvent("启动")

	//找到链接，导入协议
	tunnel := GetTunnel(dev.TunnelId)
	if tunnel == nil {
		return errors.New("找不到链接")
	}
	if tunnel.adapter == nil {
		return errors.New("未加载协议")
	}

	//绑定链接
	err := dev.BindAdapter(tunnel.adapter)
	if err != nil {
		return err
	}

	//采集器
	for _, poller := range dev.pollers {
		err := poller.Start()
		if err != nil {
			return err
		}
	}

	dev.running = true

	return nil
}

//Stop 结束设备
func (dev *Device) Stop() error {
	dev.running = false

	dev.createEvent("关闭")

	for _, poller := range dev.pollers {
		poller.Stop()
	}
	return nil
}

func (dev *Device) Running() bool {
	return dev.running
}

func (dev *Device) Refresh() error {
	for _, poller := range dev.pollers {
		poller.Execute()
	}
	return nil
}

func (dev *Device) RefreshPoint(name string) (float64, error) {
	return dev.mapper.Get(name)
}

//Execute 执行命令
func (dev *Device) Execute(command string, argv []float64) error {
	dev.createEvent("执行：" + command)

	cmd, ok := dev.commandIndex[command]
	if !ok {
		return fmt.Errorf("找不到命令：%s", command)
	}

	//参数加入环境变量
	//优先级：参数 > 表达式 > 初始值
	env := make(map[string]interface{})
	for i, v := range argv {
		env["$"+strconv.Itoa(i)] = v
	}
	for k, v := range dev.Context {
		env[k] = v
	}

	//执行
	for _, directive := range cmd.Directives {
		val := directive.Value
		if directive.Expression != "" {
			v, err := expr.Eval(directive.Expression, env)
			if err != nil {
				return err
			} else {
				val = v.(float64)
			}
		}
		//延迟执行
		if directive.Delay > 0 {
			time.AfterFunc(time.Duration(directive.Delay)*time.Millisecond, func() {
				err := dev.mapper.Set(directive.Point, val)
				dev.Emit("error", err)
			})
		} else {
			err := dev.mapper.Set(directive.Point, val)
			//dev.events.Publish("error", err)
			return err
		}
	}

	return nil
}
