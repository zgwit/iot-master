package master

import (
	"github.com/antonmedv/expr"
	"github.com/zgwit/iot-master/calc"
	"github.com/zgwit/iot-master/database"
	"github.com/zgwit/iot-master/events"
	"github.com/zgwit/iot-master/model"
	"github.com/zgwit/iot-master/tsdb"
	"strconv"
	"time"
)

//Device 设备
type Device struct {
	model.Device

	Pollers  []*Poller  `json:"pollers"`
	Reactors []*Reactor `json:"reactors"`
	Jobs     []*Job     `json:"jobs"`

	//命令索引
	commandIndex map[string]*model.Command

	adapter *Adapter

	events.EventEmitter
}

func NewDevice(m *model.Device) *Device {
	dev := &Device{
		Device:       *m,
		commandIndex: make(map[string]*model.Command, 0),
	}

	if m.Pollers != nil {
		dev.Pollers = make([]*Poller, len(m.Pollers))
		for i, v := range m.Pollers {
			dev.Pollers[i] = &Poller{Poller: *v}
		}
	} else {
		dev.Pollers = make([]*Poller, 0)
	}

	if m.Jobs != nil {
		dev.Jobs = make([]*Job, len(m.Jobs))
		for i, v := range m.Jobs {
			dev.Jobs[i] = &Job{Job: *v}
		}
	} else {
		dev.Jobs = make([]*Job, 0)
	}

	if m.Reactors != nil {
		dev.Reactors = make([]*Reactor, len(m.Reactors))
		for i, v := range m.Reactors {
			dev.Reactors[i] = &Reactor{Reactor: *v}
		}
	} else {
		dev.Reactors = make([]*Reactor, 0)
	}

	return dev
}

//Init 设备初始化
func (dev *Device) Init() error {

	metric := strconv.Itoa(dev.Id)

	//处理数据变化结果
	dev.adapter.On("data", func(data calc.Context) {
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
				dev.Context[calculator.Variable] = val
			}
		}

		//处理响应
		for _, reactor := range dev.Reactors {
			err := reactor.Execute(dev.Context)
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

	//初始化计算器
	for _, calculator := range dev.Calculators {
		err := calculator.Init()
		if err != nil {
			return err
		}
	}

	//定时任务
	for _, job := range dev.Jobs {
		err := job.Start()
		if err != nil {
			return err
		}

		job.On("invoke", func() {
			var err error
			for _, invoke := range job.Invokes {
				err = dev.Execute(invoke.Command, invoke.Argv)
				if err != nil {
					dev.Emit("error", err)
				}
			}

			//日志
			_ = database.DeviceHistoryJob.Save(model.DeviceHistoryJob{
				DeviceId: dev.Id,
				Job:      job.String(),
				History:  "action",
				Created:  time.Now()})
		})
	}

	//订阅告警
	for _, reactor := range dev.Reactors {
		reactor.On("alarm", func(alarm *model.Alarm) {
			da := &model.DeviceAlarm{
				Alarm:    *alarm,
				DeviceId: dev.Id,
				Created:  time.Now(),
			}

			//入库
			_ = database.DeviceHistoryAlarm.Save(model.DeviceHistoryAlarm{
				DeviceId: dev.Id,
				Code:     alarm.Code,
				Level:    alarm.Level,
				Message:  alarm.Message,
				Created:  time.Now(),
			})

			//上报
			dev.Emit("alarm", da)
		})

		reactor.On("invoke", func() {
			for _, invoke := range reactor.Invokes {
				err := dev.Execute(invoke.Command, invoke.Argv)
				if err != nil {
					dev.Emit("error", err)
				}
			}

			//保存历史
			history := model.DeviceHistoryReactor{
				DeviceId: dev.Id,
				Name:     reactor.Name,
				History:  "action",
				Created:  time.Now(),
			}
			if history.Name == "" {
				history.Name = reactor.Condition
			}
			_ = database.DeviceHistoryReactor.Save(history)
		})
	}

	return nil
}

//Start 设备启动
func (dev *Device) Start() error {
	_ = database.DeviceHistory.Save(model.DeviceHistory{DeviceId: dev.Id, History: "start", Created: time.Now()})

	//采集器
	for _, collector := range dev.Pollers {
		err := collector.Start()
		if err != nil {
			return err
		}
	}
	//定时任务
	for _, job := range dev.Jobs {
		err := job.Start()
		if err != nil {
			return err
		}
	}
	return nil
}

//Stop 结束设备
func (dev *Device) Stop() error {
	_ = database.DeviceHistory.Save(model.DeviceHistory{DeviceId: dev.Id, History: "stop", Created: time.Now()})

	for _, collector := range dev.Pollers {
		collector.Stop()
	}
	for _, job := range dev.Jobs {
		job.Stop()
	}
	return nil
}

//Execute 执行命令
func (dev *Device) Execute(command string, argv []float64) error {
	_ = database.DeviceHistoryCommand.Save(model.DeviceHistoryCommand{DeviceId: dev.Id, Command: command, Argv: argv, History: "execute", Created: time.Now()})

	cmd := dev.commandIndex[command]
	//直接执行
	for _, directive := range cmd.Directives {
		val := directive.Value
		//优先级：参数 > 表达式 > 初始值
		if directive.Arg > 0 {
			val = argv[directive.Arg-1]
		} else if directive.Expression != "" {
			//TODO 参数加入环境变量
			v, err := expr.Eval(directive.Expression, dev.Context)
			if err != nil {
				//dev.events.Publish("error", err)
				return err
			} else {
				val = v.(float64)
			}
		}
		//延迟执行
		if directive.Delay > 0 {
			time.AfterFunc(time.Duration(directive.Delay)*time.Millisecond, func() {
				err := dev.adapter.Set(directive.Point, val)
				dev.Emit("error", err)
			})
		} else {
			err := dev.adapter.Set(directive.Point, val)
			//dev.events.Publish("error", err)
			return err
		}
	}

	return nil
}
