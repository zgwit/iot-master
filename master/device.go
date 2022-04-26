package master

import (
	"errors"
	"fmt"
	"github.com/antonmedv/expr"
	"github.com/zgwit/iot-master/calc"
	"github.com/zgwit/iot-master/database"
	"github.com/zgwit/iot-master/events"
	"github.com/zgwit/iot-master/model"
	"github.com/zgwit/iot-master/tsdb"
	"github.com/zgwit/storm/v3"
	"strconv"
	"time"
)

//Device 设备
type Device struct {
	model.Device
	events.EventEmitter

	pollers    []*Poller
	validators []*Alarm

	//命令索引
	commandIndex map[string]*model.Command

	mapper *Mapper
}

func NewDevice(m *model.Device) (*Device, error) {
	dev := &Device{
		Device:       *m,
		commandIndex: make(map[string]*model.Command, 0),
		//mapper: newMapper(m.Points, adapter), TODO 引入协议
		pollers:    make([]*Poller, 0),
		validators: make([]*Alarm, 0),
	}

	//加载模板
	if dev.ElementID != "" {
		var template model.Element
		err := database.Master.One("ID", dev.ElementID, &template)
		if err == storm.ErrNotFound {
			return nil, errors.New("找不到模板")
		} else if err != nil {
			return nil, err
		}
		dev.DeviceContent = template.DeviceContent
	}

	err := dev.initMapper()
	if err != nil {
		return nil, err
	}

	err = dev.initPollers()
	if err != nil {
		return nil, err
	}

	err = dev.initCalculators()
	if err != nil {
		return nil, err
	}

	err = dev.initValidators()
	if err != nil {
		return nil, err
	}

	return dev, nil
}

func (dev *Device) initMapper() error {
	var err error
	//找到链接，导入协议
	link := GetLink(dev.LinkID)
	if link == nil {
		//TODO error
		return nil
	}
	if link.adapter == nil {
		//TODO error
		return nil
	}

	dev.mapper, err = newMapper(dev.Station, dev.Points, link.adapter)
	metric := strconv.Itoa(dev.ID)

	//处理数据变化结果
	dev.mapper.On("data", func(data calc.Context) {
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
				dev.Context[calculator.As] = val
			}
		}

		//处理策略
		for _, validator := range dev.validators {
			err := validator.Execute(dev.Context)
			if err != nil {
				dev.Emit("error", err)
			}
		}

		//向上广播
		dev.Emit("data", data)

		//保存到时序数据库
		//是否有必要起协程 或者 使用单一进程进行写入
		go func() {
			for k, v := range data {
				_ = tsdb.Save(metric, k, v.(float64))
			}
		}()
	})
	return err
}

func (dev *Device) initPollers() error {
	if dev.Pollers == nil {
		return nil
	}
	for _, v := range dev.Pollers {
		dev.pollers = append(dev.pollers, &Poller{Poller: *v})
	}
	return nil
}

func (dev *Device) initValidators() error {
	if dev.Validators == nil {
		return nil
	}
	for _, v := range dev.Validators {
		validator := &Alarm{Alarm: *v}
		validator.On("alarm", func(alarm *model.AlarmContent) {
			da := &model.DeviceAlarm{DeviceID: dev.ID, AlarmContent: *alarm}

			//入库
			_ = database.History.Save(da)
			dev.createEvent("告警：" + alarm.Message)

			//上报
			dev.Emit("alarm", da)
		})
		dev.validators = append(dev.validators, validator)
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
	_ = database.History.Save(model.DeviceEvent{DeviceID: dev.ID, Event: event})
}

//Start 设备启动
func (dev *Device) Start() error {
	dev.createEvent("启动")

	//采集器
	for _, poller := range dev.pollers {
		err := poller.Start()
		if err != nil {
			return err
		}
	}
	return nil
}

//Stop 结束设备
func (dev *Device) Stop() error {
	dev.createEvent("关闭")

	for _, poller := range dev.pollers {
		poller.Stop()
	}
	return nil
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
				err := dev.mapper.Set(directive.Address, val)
				dev.Emit("error", err)
			})
		} else {
			err := dev.mapper.Set(directive.Address, val)
			//dev.events.Publish("error", err)
			return err
		}
	}

	return nil
}
