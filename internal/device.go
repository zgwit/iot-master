package interval

import (
	"github.com/antonmedv/expr"
	"github.com/asaskevich/EventBus"
	"time"
)

type Device struct {
	Disabled bool `json:"disabled"`

	Id   string   `json:"id"`
	Name string   `json:"name"`
	Tags []string `json:"tags"`

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

func (c *Device) Start() error {

	//处理数据变化结果
	_ = c.adapter.events.Subscribe("data", func(data Context) {
		//更新上下文
		for k, v := range data {
			c.Context[k] = v
		}

		//数据变化后，更新计算
		for _, v := range c.Calculators {
			val, err := v.Evaluate()
			if err != nil {
				c.events.Publish("error", err)
			} else {
				c.Context[v.Variable] = val
			}
		}

		//处理响应
		for _, v := range c.Reactors {
			err := v.Execute()
			if err != nil {
				c.events.Publish("error", err)
			}
		}

		//向上广播
		c.events.Publish("data", data)
	})

	//采集器数据变化
	for _, v := range c.Collectors {
		err := v.Start()
		if err != nil {
			return err
		}
	}

	//计算器数据变化
	for _, v := range c.Calculators {
		_ = v.Init(c.Context)
	}

	//定时任务
	for _, v := range c.Jobs {
		err := v.Start()
		if err != nil {
			return err
		}

		_ = v.events.Subscribe("invoke", func() {
			//TODO 日志
			for _, i := range v.Invokes {
				err := c.Execute(i.Command, i.Argv)
				if err != nil {
					c.events.Publish("error", err)
				}
			}
		})
	}

	//订阅告警
	for _, v := range c.Reactors {
		v.Init()

		_ = v.events.Subscribe("alarm", func(alarm *Alarm) {
			da := &DeviceAlarm{
				Alarm:    *alarm,
				DeviceId: 0,
				Created:  time.Now(),
			}
			//TODO 入库

			//上报
			c.events.Publish("alarm", da)
		})

		_ = v.events.Subscribe("invoke", func() {
			//TODO 日志
			for _, i := range v.Invokes {
				err := c.Execute(i.Command, i.Argv)
				if err != nil {
					c.events.Publish("error", err)
				}
			}
		})
	}

	return nil
}

func (c *Device) Stop() error {
	return nil
}

func (c *Device) Execute(command string, argv []float64) error {
	cmd := c.commandIndex[command]
	//直接执行
	for _, d := range cmd.Directives {
		val := d.Value
		//优先级：参数 > 表达式 > 初始值
		if d.Arg > 0 {
			val = argv[d.Arg-1]
		} else if d.Expression != "" {
			//TODO 参数加入环境变量
			v, err := expr.Eval(d.Expression, c.Context)
			if err != nil {
				//c.events.Publish("error", err)
				return err
			} else {
				val = v.(float64)
			}
		}
		//延迟执行
		if d.Delay > 0 {
			time.AfterFunc(time.Duration(d.Delay)*time.Millisecond, func() {
				err := c.adapter.Set(d.Point, val)
				c.events.Publish("error", err)
			})
		} else {
			err := c.adapter.Set(d.Point, val)
			//c.events.Publish("error", err)
			return err
		}
	}

	return nil
}
