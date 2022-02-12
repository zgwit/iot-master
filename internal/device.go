package internal

import (
	"github.com/antonmedv/expr"
	"github.com/zgwit/iot-master/events"
	"github.com/zgwit/iot-master/internal/calc"
	"time"
)


type Device struct {
	Disabled bool     `json:"disabled"`

	Id       int      `json:"id" storm:"id,increment"`
	Name     string   `json:"name"`
	Tags     []string `json:"tags"`

	//从机号
	Slave int `json:"slave"`

	Points      []*Point      `json:"points"`
	Collectors  []*Collector  `json:"collectors"`
	Calculators []*Calculator `json:"calculators"`
	Commands    []*Command    `json:"commands"`
	Reactors    []*Reactor    `json:"reactors"`
	Jobs        []*Job        `json:"jobs"`

	//上下文
	Context calc.Context `json:"context"`

	//命令索引
	commandIndex map[string]*Command

	adapter *Adapter

	events.EventEmitter
}

func (dev *Device) Init() error {

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
			//TODO 日志
			for _, invoke := range job.Invokes {
				err := dev.Execute(invoke.Command, invoke.Argv)
				if err != nil {
					dev.Emit("error", err)
				}
			}
		})
	}

	//订阅告警
	for _, reactor := range dev.Reactors {
		reactor.On("alarm", func(alarm *Alarm) {
			da := &DeviceAlarm{
				Alarm:    *alarm,
				DeviceId: dev.Id,
				Created:  time.Now(),
			}
			//TODO 入库

			//上报
			dev.Emit("alarm", da)
		})

		reactor.On("invoke", func() {
			//TODO 日志
			for _, invoke := range reactor.Invokes {
				err := dev.Execute(invoke.Command, invoke.Argv)
				if err != nil {
					dev.Emit("error", err)
				}
			}
		})
	}

	return nil
}

func (dev *Device) Start() error {
	//采集器
	for _, collector := range dev.Collectors {
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

func (dev *Device) Stop() error {
	for _, collector := range dev.Collectors {
		collector.Stop()
	}
	for _, job := range dev.Jobs {
		job.Stop()
	}
	return nil
}

func (dev *Device) Execute(command string, argv []float64) error {
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


type DeviceHistory struct {
	Id       int       `json:"id" storm:"id,increment"`
	DeviceId int       `json:"device_id"`
	History  string    `json:"history"`
	Created  time.Time `json:"created"`
}

type DeviceHistoryAlarm struct {
	Id int `json:"id" storm:"id,increment"`

	DeviceId int    `json:"Device_id"`
	Code     string `json:"code"`
	Level    int    `json:"level"`
	Message  string `json:"message"`

	Created time.Time `json:"created"`
}

type DeviceHistoryJob struct {
	Id      int       `json:"id" storm:"id,increment"`
	Job     string    `json:"job"`
	Result  string    `json:"result"`
	Created time.Time `json:"created"`
}

type DeviceHistoryCommand struct {
	Id      int       `json:"id" storm:"id,increment"`
	Command string    `json:"command"`
	Argv    string    `json:"argv"`
	Result  string    `json:"result"`
	Created time.Time `json:"created"`
}

