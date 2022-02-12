package internal

import (
	"github.com/zgwit/iot-master/events"
	"github.com/zgwit/iot-master/internal/aggregator"
	"github.com/zgwit/iot-master/internal/calc"
	"time"
)

type ProjectDevice struct {
	Id     int    `json:"id" storm:"id,increment"`
	Name   string `json:"name"`
	device *Device
}

type Project struct {
	Id       int  `json:"id" storm:"id,increment"`
	Disabled bool `json:"disabled"`

	Devices []*ProjectDevice `json:"devices"`

	Aggregators []*aggregator.Aggregator `json:"aggregators"`
	Commands    []*Command               `json:"commands"`
	Reactors    []*Reactor               `json:"reactors"`
	Jobs        []*Job                   `json:"jobs"`

	Context calc.Context `json:"context"`

	deviceNameIndex map[string]*Device
	deviceIdIndex   map[int]*Device

	events.EventEmitter

	deviceDataHandler  func(data calc.Context)
	deviceAlarmHandler func(alarm *DeviceAlarm)
}

func (prj *Project) Init() error {
	//设备数据变化的处理函数
	prj.deviceDataHandler = func(data calc.Context) {
		//数据变化后，更新计算
		for _, agg := range prj.Aggregators{
			val, err := agg.Evaluate()
			if err != nil {
				prj.Emit("error", err)
			} else {
				prj.Context[agg.As] = val
			}
		}

		//处理响应
		for _,reactor := range prj.Reactors {
			err := reactor.Execute(prj.Context)
			if err != nil {
				prj.Emit("error", err)
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
		prj.Emit("alarm", pa)
	}

	//初始化设备
	for _,d:= range prj.Devices {
		dev := GetDevice(d.Id)
		//TODO 如果找不到设备，该怎么处理
		d.device = dev
		prj.deviceNameIndex[d.Name] = dev
		prj.deviceIdIndex[d.Id] = dev
		prj.Context[d.Name] = dev.Context //两级上下文
		//_ = dev.events.Subscribe("data", prj.deviceDataHandler)
	}

	//定时任务
	for _, job := range prj.Jobs {
		job.On("invoke", func() {
			//TODO 日志
			for _, invoke := range job.Invokes {
				err := prj.execute(invoke)
				if err != nil {
					prj.Emit("error", err)
				}
			}
		})
	}

	//初始化聚合器
	for _, agg := range prj.Aggregators {
		err := agg.Init()
		if err != nil {
			return err
		}
		for _, d := range prj.Devices {
			if hasSelect(&agg.Select, d) {
				agg.Push(d.device.Context)
			}
		}
	}

	//订阅告警
	for _, reactor := range prj.Reactors {
		reactor.On("alarm", func(alarm *Alarm) {
			pa := &ProjectAlarm{
				DeviceAlarm: DeviceAlarm{
					Alarm:   *alarm,
					Created: time.Now(),
				},
				ProjectId: prj.Id,
			}
			//TODO 入库

			//上报
			prj.Emit("alarm", pa)
		})

		reactor.On("invoke", func() {
			//TODO 日志
			for _, invoke := range reactor.Invokes {
				err := prj.execute(invoke)
				if err != nil {
					prj.Emit("error", err)
				}
			}
		})
	}

	return nil
}

func (prj *Project) Start() error {
	//订阅设备的数据变化和报警
	for _,dev := range prj.Devices {
		dev.device.On("data", prj.deviceDataHandler)
		dev.device.On("alarm", prj.deviceAlarmHandler)
	}

	//定时任务
	for _, job := range prj.Jobs {
		err := job.Start()
		if err != nil {
			return err
		}
	}
	return nil
}

func (prj *Project) Stop() error {
	for _,dev := range prj.Devices {
		dev.device.Off("data", prj.deviceDataHandler)
		dev.device.Off("alarm", prj.deviceAlarmHandler)
	}
	for _, job := range prj.Jobs {
		job.Stop()
	}
	return nil
}

func (prj *Project) execute(in *Invoke) error {
	for _, d := range prj.Devices {
		if hasSelect(&in.Select, d) {
			err := d.device.Execute(in.Command, in.Argv)
			if err != nil {
				return err
			}
		}
	}
	return nil
}


type ProjectHistory struct {
	Id        int       `json:"id" storm:"id,increment"`
	ProjectId int       `json:"project_id"`
	History   string    `json:"history"`
	Created   time.Time `json:"created"`
}

type ProjectHistoryAlarm struct {
	Id int `json:"id" storm:"id,increment"`

	ProjectId int    `json:"project_id"`
	DeviceId  int    `json:"device_id"`
	Code      string `json:"code"`
	Level     int    `json:"level"`
	Message   string `json:"message"`

	Created time.Time `json:"created"`
}

type ProjectHistoryJob struct {
	Id      int       `json:"id" storm:"id,increment"`
	Job     string    `json:"job"`
	Result  string    `json:"result"`
	Created time.Time `json:"created"`
}