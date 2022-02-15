package master

import (
	"github.com/zgwit/iot-master/database"
	"github.com/zgwit/iot-master/events"
	"github.com/zgwit/iot-master/master/aggregator"
	"github.com/zgwit/iot-master/master/calc"
	_select "github.com/zgwit/iot-master/master/select"
	"strings"
	"time"
)

type ProjectDevice struct {
	Id     int    `json:"id" storm:"id,increment"`
	Name   string `json:"name"`
	device *Device
}

func hasTag(a, b []string) bool {
	for i := len(a); i >= 0; i-- {
		for j := len(b); j >= 0; j-- {
			if strings.EqualFold(a[i], b[j]) {
				return true
			}
		}
	}
	return false
}

func (d *ProjectDevice) checkSelect(s *_select.Select) bool {
	for _, name := range s.Names {
		if name == d.Name {
			return true
		}
	}
	for _, name := range s.Ids {
		if name == d.Id {
			return true
		}
	}
	if hasTag(s.Tags, d.device.Tags) {
		return true
	}
	return false
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
			var err error
			for _, invoke := range job.Invokes {
				err = prj.execute(invoke)
				if err != nil {
					prj.Emit("error", err)
				}
			}
			
			//日志
			_ = database.ProjectHistoryJob.Save(ProjectHistoryJob{
				ProjectId: prj.Id,
				Job:      job.String(),
				History:  "action",
				Created:  time.Now()})
		})
	}

	//初始化聚合器
	for _, agg := range prj.Aggregators {
		err := agg.Init()
		if err != nil {
			return err
		}
		for _, d := range prj.Devices {
			if d.checkSelect(&agg.Select) {
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

			//入库
			_ = database.ProjectHistoryAlarm.Save(ProjectHistoryAlarm{
				ProjectId: prj.Id,
				Code:     alarm.Code,
				Level:    alarm.Level,
				Message:  alarm.Message,
				Created:  time.Now(),
			})

			//上报
			prj.Emit("alarm", pa)
		})

		reactor.On("invoke", func() {
			for _, invoke := range reactor.Invokes {
				err := prj.execute(invoke)
				if err != nil {
					prj.Emit("error", err)
				}
			}

			//保存历史
			history := ProjectHistoryReactor{
				ProjectId: prj.Id,
				Name:     reactor.Name,
				History:  "action",
				Created:  time.Now(),
			}
			if history.Name == "" {
				history.Name = reactor.Condition
			}
			_ = database.ProjectHistoryReactor.Save(history)
		})
	}

	return nil
}

func (prj *Project) Start() error {
	_ = database.ProjectHistory.Save(ProjectHistory{ProjectId: prj.Id, History: "start", Created: time.Now()})
	
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
	_ = database.ProjectHistory.Save(ProjectHistory{ProjectId: prj.Id, History: "stop", Created: time.Now()})

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
		if d.checkSelect(&in.Select) {
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

type ProjectHistoryReactor struct {
	Id       int       `json:"id" storm:"id,increment"`
	ProjectId int       `json:"project_id"`
	Name     string    `json:"name"`
	History  string    `json:"result"`
	Created  time.Time `json:"created"`
}


type ProjectHistoryJob struct {
	Id      int       `json:"id" storm:"id,increment"`
	ProjectId int       `json:"project_id"`
	Job     string    `json:"job"`
	History  string    `json:"result"`
	Created time.Time `json:"created"`
}