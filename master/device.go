package master

import (
	"github.com/antonmedv/expr"
	"github.com/asdine/storm/v3"
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
	events.EventEmitter

	pollers    []*Poller
	strategies []*Strategy
	jobs       []*Job
	timers     []*Timer

	//命令索引
	commandIndex map[string]*model.Command

	adapter *Adapter
}

func NewDevice(m *model.Device) *Device {
	dev := &Device{
		Device:       *m,
		commandIndex: make(map[string]*model.Command, 0),
	}

	if m.Pollers != nil {
		dev.pollers = make([]*Poller, len(m.Pollers))
		for _, v := range m.Pollers {
			dev.pollers = append(dev.pollers, &Poller{Poller: *v})
		}
	} else {
		dev.pollers = make([]*Poller, 0)
	}

	if m.Jobs != nil {
		dev.jobs = make([]*Job, len(m.Jobs))
		for _, v := range m.Jobs {
			dev.jobs = append(dev.jobs, &Job{Job: *v})
		}
	} else {
		dev.jobs = make([]*Job, 0)
	}

	if m.Strategies != nil {
		dev.strategies = make([]*Strategy, len(m.Strategies))
		for _, v := range m.Strategies {
			dev.strategies = append(dev.strategies, &Strategy{Strategy: *v})
		}
	} else {
		dev.strategies = make([]*Strategy, 0)
	}

	dev.timers = make([]*Timer, 0)

	return dev
}

//Init 设备初始化
func (dev *Device) Init() error {

	metric := strconv.Itoa(dev.ID)

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
		for _, reactor := range dev.strategies {
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
	for _, job := range dev.jobs {
		job.On("invoke", func() {
			var err error
			for _, invoke := range job.Invokes {
				err = dev.Execute(invoke.Command, invoke.Argv)
				if err != nil {
					dev.Emit("error", err)
				}
			}

			//日志
			_ = database.History.Save(model.DeviceHistoryJob{
				DeviceHistory: model.DeviceHistory{
					DeviceID: dev.ID,
					History:  "action",
					Created:  time.Now(),
				},
				Job: job.String(),
			})
		})
	}

	//订阅告警
	for _, strategy := range dev.strategies {
		strategy.On("alarm", func(alarm *model.Alarm) {
			da := &model.DeviceAlarm{
				Alarm:    *alarm,
				DeviceID: dev.ID,
				Created:  time.Now(),
			}

			//入库
			_ = database.History.Save(model.DeviceHistoryAlarm{
				DeviceHistory: model.DeviceHistory{
					DeviceID: dev.ID,
					History:  "action",
					Created:  time.Now(),
				},
				Code:     alarm.Code,
				Level:    alarm.Level,
				Message:  alarm.Message,
			})

			//上报
			dev.Emit("alarm", da)
		})

		strategy.On("invoke", func() {
			for _, invoke := range strategy.Invokes {
				err := dev.Execute(invoke.Command, invoke.Argv)
				if err != nil {
					dev.Emit("error", err)
				}
			}

			//保存历史
			history := model.DeviceHistoryReactor{
				DeviceHistory: model.DeviceHistory{
					DeviceID: dev.ID,
					History:  "action",
					Created:  time.Now(),
				},
				Name:     strategy.Name,
			}
			if history.Name == "" {
				history.Name = strategy.Condition
			}
			_ = database.History.Save(history)
		})
	}

	return nil
}

//Start 设备启动
func (dev *Device) Start() error {
	_ = database.History.Save(model.DeviceHistory{DeviceID: dev.ID, History: "start", Created: time.Now()})

	//采集器
	for _, collector := range dev.pollers {
		err := collector.Start()
		if err != nil {
			return err
		}
	}
	//定时任务
	for _, job := range dev.jobs {
		err := job.Start()
		if err != nil {
			return err
		}
	}
	//用户定时任务
	for _, timer := range dev.timers {
		err := timer.Start()
		if err != nil {
			return err
		}
	}
	return nil
}

//Stop 结束设备
func (dev *Device) Stop() error {
	_ = database.History.Save(model.DeviceHistory{DeviceID: dev.ID, History: "stop", Created: time.Now()})

	for _, collector := range dev.pollers {
		collector.Stop()
	}
	for _, job := range dev.jobs {
		job.Stop()
	}
	for _, timer := range dev.timers {
		timer.Stop()
	}
	return nil
}

//Execute 执行命令
func (dev *Device) Execute(command string, argv []float64) error {
	_ = database.History.Save(model.DeviceHistoryCommand{
		DeviceHistory: model.DeviceHistory{
			DeviceID: dev.ID,
			History:  "action",
			Created:  time.Now(),
		},
		Command: command,
		Argv: argv,
	})

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

func (dev *Device) LoadTimers() error {
	var timers []model.DeviceTimer
	err := database.Master.All(&timers) //TODO 判断disabled

	if err != storm.ErrNotFound {
		return nil
	} else if err != nil {
		return err
	}

	for _, t := range timers {
		timer := &Timer{Timer: t.Timer}
		dev.timers = append(dev.timers, timer)

		timer.On("invoke", func() {
			var err error
			for _, invoke := range timer.Invokes {
				err = dev.Execute(invoke.Command, invoke.Argv)
				if err != nil {
					dev.Emit("error", err)
				}
			}

			//日志
			_ = database.History.Save(model.DeviceHistoryTimer{
				DeviceHistory: model.DeviceHistory{
					DeviceID: dev.ID,
					History:  "action",
					Created:  time.Now(),
				},
				TimerID: timer.ID,
			})
		})
	}

	return nil
}
