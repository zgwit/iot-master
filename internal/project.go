package interval

import (
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

	deviceDataHandler  func(data Context)
	deviceAlarmHandler func(alarm *DeviceAlarm)
}

func (prj *Project) Init() error {
	//设备数据变化的处理函数
	prj.deviceDataHandler = func(data Context) {
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

	//设备告警的处理函数
	prj.deviceAlarmHandler = func(alarm *DeviceAlarm) {
		pa := &ProjectAlarm{
			DeviceAlarm: *alarm,
			ProjectId:   prj.Id,
		}
		//TODO 入库

		//上报
		prj.events.Publish("alarm", pa)
	}

	//初始化设备
	for i := 0; i < len(prj.Devices); i++ {
		d := &prj.Devices[i]
		dev := GetDevice(d.Id)
		//TODO 如果找不到设备，该怎么处理
		d.device = dev
		prj.deviceNameIndex[d.Name] = dev
		prj.deviceIdIndex[d.Id] = dev
		prj.Context[d.Name] = dev.Context //两级上下文
		//_ = dev.events.Subscribe("data", prj.deviceDataHandler)
	}

	//定时任务
	for i := 0; i < len(prj.Jobs); i++ {
		v := &prj.Jobs[i]
		v.Init()

		_ = v.events.Subscribe("invoke", func() {
			//TODO 日志
			for _, invoke := range v.Invokes {
				err := prj.execute(&invoke)
				if err != nil {
					prj.events.Publish("error", err)
				}
			}
		})
	}

	//初始化聚合器
	for i := 0; i < len(prj.Aggregators); i++ {
		a := &prj.Aggregators[i]
		a.Init()
		for _, d := range prj.Devices {
			if a.Select.has(&d) {
				err := a.Push(d.device.Context)
				if err != nil {
					return err
				}
			}
		}
	}

	//订阅告警
	//for _, v := range dev.Reactors {
	for i := 0; i < len(prj.Reactors); i++ {
		reactor := &prj.Reactors[i]
		reactor.Init()

		_ = reactor.events.Subscribe("alarm", func(alarm *Alarm) {
			pa := &ProjectAlarm{
				DeviceAlarm: DeviceAlarm{
					Alarm:   *alarm,
					Created: time.Now(),
				},
				ProjectId: prj.Id,
			}
			//TODO 入库

			//上报
			prj.events.Publish("alarm", pa)
		})

		_ = reactor.events.Subscribe("invoke", func() {
			//TODO 日志
			for _, invoke := range reactor.Invokes {
				err := prj.execute(&invoke)
				if err != nil {
					prj.events.Publish("error", err)
				}
			}
		})
	}

	return nil
}

func (prj *Project) Start() error {
	//订阅设备的数据变化和报警
	for i := 0; i < len(prj.Devices); i++ {
		dev := prj.Devices[i].device
		_ = dev.events.Subscribe("data", prj.deviceDataHandler)
		_ = dev.events.Subscribe("alarm", prj.deviceAlarmHandler)
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
		dev := prj.Devices[i].device
		_ = dev.events.Unsubscribe("data", prj.deviceDataHandler)
		_ = dev.events.Unsubscribe("alarm", prj.deviceAlarmHandler)
	}
	for i := 0; i < len(prj.Jobs); i++ {
		prj.Jobs[i].Stop()
	}
	return nil
}

func (prj *Project) execute(in *Invoke) error {
	for _, d := range prj.Devices {
		if in.Select.has(&d) {
			err := d.device.Execute(in.Command, in.Argv)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
