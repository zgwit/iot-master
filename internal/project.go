package interval

import (
	"github.com/antonmedv/expr"
	"github.com/asaskevich/EventBus"
	"github.com/zgwit/iot-master/internal/aggregator"
	"time"
)

type ProjectDevice struct {
	Id     int    `json:"id" storm:"id,increment"`
	Name   string `json:"name"`
	device *Device
}

type Project struct {
	Id       int  `json:"id"`
	Disabled bool `json:"disabled"`

	Devices []ProjectDevice `json:"devices"`

	Aggregators []aggregator.Aggregator `json:"aggregators"`
	Commands    []Command               `json:"commands"`
	Reactors    []Reactor               `json:"reactors"`
	Jobs        []Job                   `json:"jobs"`

	Context Context `json:"context"`

	deviceNameIndex map[string]*Device
	deviceIdIndex   map[int]*Device

	events EventBus.Bus

	handler func(data Context)
}

func (prj *Project) Init() error {
	prj.handler = func(data Context) {
		//数据变化后，更新计算
		for i := 0; i < len(prj.Aggregators); i++ {
			agg := &prj.Aggregators[i]
			val, err := agg.Evaluate()
			if err != nil {
				prj.events.Publish("error", err)
			} else {
				prj.Context[agg.As] = val
			}
		}

		//处理响应
		for i := 0; i < len(prj.Reactors); i++ {
			reactor := &prj.Reactors[i]
			err := reactor.Execute()
			if err != nil {
				prj.events.Publish("error", err)
			}
		}
	}

	for i := 0; i < len(prj.Devices); i++ {
		d := &prj.Devices[i]
		dev := GetDevice(d.Id)
		d.device = dev
		prj.deviceNameIndex[d.Name] = dev
		prj.deviceIdIndex[d.Id] = dev
		prj.Context[d.Name] = dev.Context //两级上下文
		//_ = dev.events.Subscribe("data", prj.handler)
	}

	//定时任务
	//for _, v := range dev.Jobs {
	for i := 0; i < len(prj.Jobs); i++ {
		v := &prj.Jobs[i]
		err := v.Start()
		if err != nil {
			return err
		}

		_ = v.events.Subscribe("invoke", func() {
			//TODO 日志
			for _, invoke := range v.Invokes {
				err := prj.Execute(invoke.Command, invoke.Argv)
				if err != nil {
					prj.events.Publish("error", err)
				}
			}
		})
	}

	//订阅告警
	//for _, v := range dev.Reactors {
	for i := 0; i < len(prj.Reactors); i++ {
		reactor := &prj.Reactors[i]
		reactor.Init()

		_ = reactor.events.Subscribe("alarm", func(alarm *DeviceAlarm) {
			pa := &ProjectAlarm{
				DeviceAlarm: *alarm,
				ProjectId:   prj.Id,
			}
			//TODO 入库

			//上报
			prj.events.Publish("alarm", pa)
		})

		_ = reactor.events.Subscribe("invoke", func() {
			//TODO 日志
			for _, invoke := range reactor.Invokes {
				err := prj.Execute(invoke.Command, invoke.Argv)
				if err != nil {
					prj.events.Publish("error", err)
				}
			}
		})
	}

	return nil
}

func (prj *Project) Start() error {

	for i := 0; i < len(prj.Devices); i++ {
		_ = prj.Devices[i].device.events.Subscribe("data", prj.handler)
	}

	//定时任务
	for i := 0; i < len(prj.Jobs); i++ {
		err := prj.Jobs[i].Start()
		if err != nil {
			return err
		}
	}
	return nil
}

func (prj *Project) Stop() error {
	for i := 0; i < len(prj.Devices); i++ {
		_ = prj.Devices[i].device.events.Unsubscribe("data", prj.handler)
	}
	for i := 0; i < len(prj.Jobs); i++ {
		prj.Jobs[i].Stop()
	}
	return nil
}

func (prj *Project) execute(in *Invoke) error {
	devices := make(map[*Device]bool) //set
	for _, id:= range in.Ids {
		d := prj.deviceIdIndex[id]
		devices[d] = true
	}
	for _, name:= range in.Names {
		d := prj.deviceNameIndex[name]
		devices[d] = true
	}
	for _, d := range prj.Devices {
		//TODO 判断tags交积
		devices[d.device] = true
	}
	if len(devices)==0 {
		return nil
	}

	//让设备依次执行
	for dev, _ := range devices {
		err := dev.Execute(in.Command, in.Argv)
		if err != nil {
			return err
		}
	}

	//TODO 以上代码可以再整合，外循环为遍历prj.Devices

	return nil
}
