package interval

import (
	"github.com/antonmedv/expr"
	"github.com/asaskevich/EventBus"
	"sync"
	"time"
)


type Device struct {
	Disabled bool     `json:"disabled"`

	Id       int      `json:"id" storm:"id,increment"`
	Name     string   `json:"name"`
	Tags     []string `json:"tags"`

	//从机号
	Slave int `json:"slave"`

	Points      []Point      `json:"points"`
	Collectors  []Collector  `json:"collectors"`
	Calculators []Calculator `json:"calculators"`
	Commands    []Command    `json:"commands"`
	Reactors    []Reactor    `json:"reactors"`
	Jobs        []Job        `json:"jobs"`

	//上下文
	Context Context `json:"context"`

	//命令索引
	commandIndex map[string]*Command

	events EventBus.Bus

	adapter *Adapter
}

func (dev *Device) Init() error {

	//处理数据变化结果
	_ = dev.adapter.events.Subscribe("data", func(data Context) {
		//更新上下文
		for k, v := range data {
			dev.Context[k] = v
		}

		//数据变化后，更新计算
		//for _, v := range dev.Calculators {
		for i := 0; i < len(dev.Calculators); i++ {
			calculator := &dev.Calculators[i]
			val, err := calculator.Evaluate()
			if err != nil {
				dev.events.Publish("error", err)
			} else {
				dev.Context[calculator.Variable] = val
			}
		}

		//处理响应
		//for _, v := range dev.Reactors {
		for i := 0; i < len(dev.Reactors); i++ {
			reactor := &dev.Reactors[i]
			err := reactor.Execute()
			if err != nil {
				dev.events.Publish("error", err)
			}
		}

		//向上广播
		dev.events.Publish("data", data)
	})


	//初始化计算器
	//for _, v := range dev.Calculators {
	for i := 0; i < len(dev.Calculators); i++ {
		calculator := &dev.Calculators[i]
		_ = calculator.Init(dev.Context)
	}

	//定时任务
	//for _, v := range dev.Jobs {
	for i := 0; i < len(dev.Jobs); i++ {
		v := &dev.Jobs[i]
		err := v.Start()
		if err != nil {
			return err
		}

		_ = v.events.Subscribe("invoke", func() {
			//TODO 日志
			for _, invoke := range v.Invokes {
				err := dev.Execute(invoke.Command, invoke.Argv)
				if err != nil {
					dev.events.Publish("error", err)
				}
			}
		})
	}

	//订阅告警
	//for _, v := range dev.Reactors {
	for i := 0; i < len(dev.Reactors); i++ {
		reactor := &dev.Reactors[i]
		reactor.Init()

		_ = reactor.events.Subscribe("alarm", func(alarm *Alarm) {
			da := &DeviceAlarm{
				Alarm:    *alarm,
				DeviceId: dev.Id,
				Created:  time.Now(),
			}
			//TODO 入库

			//上报
			dev.events.Publish("alarm", da)
		})

		_ = reactor.events.Subscribe("invoke", func() {
			//TODO 日志
			for _, invoke := range reactor.Invokes {
				err := dev.Execute(invoke.Command, invoke.Argv)
				if err != nil {
					dev.events.Publish("error", err)
				}
			}
		})
	}

	return nil
}

func (dev *Device) Start() error {
	//采集器
	for i := 0; i < len(dev.Collectors); i++ {
		err := dev.Collectors[i].Start()
		if err != nil {
			return err
		}
	}
	//定时任务
	for i := 0; i < len(dev.Jobs); i++ {
		err := dev.Jobs[i].Start()
		if err != nil {
			return err
		}
	}
	return nil
}

func (dev *Device) Stop() error {
	for i := 0; i < len(dev.Collectors); i++ {
		dev.Collectors[i].Stop()
	}
	for i := 0; i < len(dev.Jobs); i++ {
		dev.Jobs[i].Stop()
	}
	return nil
}

func (dev *Device) Execute(command string, argv []float64) error {
	cmd := dev.commandIndex[command]
	//直接执行
	//for _, d := range cmd.Directives {
	for i := 0; i < len(cmd.Directives); i++ {
		directive := &cmd.Directives[i]
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
				dev.events.Publish("error", err)
			})
		} else {
			err := dev.adapter.Set(directive.Point, val)
			//dev.events.Publish("error", err)
			return err
		}
	}

	return nil
}
