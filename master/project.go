package master

import (
	"github.com/zgwit/iot-master/aggregator"
	"github.com/zgwit/iot-master/calc"
	"github.com/zgwit/iot-master/database"
	"github.com/zgwit/iot-master/events"
	"github.com/zgwit/iot-master/model"
	"strings"
	"time"
)

//ProjectDevice 项目的设备
type ProjectDevice struct {
	model.ProjectDevice

	device *Device
}

func hasTag(a, b []string) bool {
	if a != nil && b != nil {
		for i := len(a); i >= 0; i-- {
			for j := len(b); j >= 0; j-- {
				if strings.EqualFold(a[i], b[j]) {
					return true
				}
			}
		}
	}
	return false
}

func (d *ProjectDevice) belongSelector(s *model.Selector) bool {
	if s.Names != nil {
		for _, name := range s.Names {
			if name == d.Name {
				return true
			}
		}
	}
	if s.IDs != nil {
		for _, id := range s.IDs {
			if id == d.ID {
				return true
			}
		}
	}
	if s.Tags != nil && len(s.Tags) > 0 && d.device.Tags != nil && len(d.device.Tags) > 0 {
		for i := len(s.Tags); i >= 0; i-- {
			for j := len(d.device.Tags); j >= 0; j-- {
				if strings.EqualFold(s.Tags[i], d.device.Tags[j]) {
					return true
				}
			}
		}
	}
	return false
}

//Project 项目
type Project struct {
	model.Project

	Devices []*ProjectDevice

	Aggregators []aggregator.Aggregator
	Jobs        []*Job
	Strategies  []*Strategy
	UserJobs    []*UserJob

	deviceNameIndex map[string]*Device
	deviceIDIndex   map[int]*Device

	events.EventEmitter

	deviceDataHandler  func(data calc.Context)
	deviceAlarmHandler func(alarm *model.DeviceAlarm)
}

func NewProject(m *model.Project) *Project {
	prj := &Project{
		Project:         *m,
		deviceNameIndex: make(map[string]*Device),
		deviceIDIndex:   make(map[int]*Device),
	}

	if m.Devices != nil {
		prj.Devices = make([]*ProjectDevice, len(m.Devices))
		for _, v := range m.Devices {
			prj.Devices = append(prj.Devices, &ProjectDevice{ProjectDevice: *v})
		}
	} else {
		prj.Devices = make([]*ProjectDevice, 0)
	}

	if m.Aggregators != nil {
		prj.Aggregators = make([]aggregator.Aggregator, len(m.Aggregators))
		for _, v := range m.Aggregators {
			agg, err := aggregator.New(v)
			if err != nil {
				return nil //TODO err
			}
			prj.Aggregators = append(prj.Aggregators, agg)
		}
	} else {
		prj.Aggregators = make([]aggregator.Aggregator, 0)
	}

	if m.Jobs != nil {
		prj.Jobs = make([]*Job, len(m.Jobs))
		for _, v := range m.Jobs {
			prj.Jobs = append(prj.Jobs, &Job{Job: *v})
		}
	} else {
		prj.Jobs = make([]*Job, 0)
	}

	if m.Strategies != nil {
		prj.Strategies = make([]*Strategy, len(m.Strategies))
		for _, v := range m.Strategies {
			prj.Strategies = append(prj.Strategies, &Strategy{Strategy: *v})
		}
	} else {
		prj.Strategies = make([]*Strategy, 0)
	}

	return prj
}

//Init 项目初始化
func (prj *Project) Init() error {
	//设备数据变化的处理函数
	prj.deviceDataHandler = func(data calc.Context) {
		//数据变化后，更新计算
		for _, agg := range prj.Aggregators {
			val, err := agg.Evaluate()
			if err != nil {
				prj.Emit("error", err)
			} else {
				prj.Context[agg.Model().As] = val
			}
		}

		//处理响应
		for _, reactor := range prj.Strategies {
			err := reactor.Execute(prj.Context)
			if err != nil {
				prj.Emit("error", err)
			}
		}
	}

	//设备告警的处理函数
	prj.deviceAlarmHandler = func(alarm *model.DeviceAlarm) {
		pa := &model.ProjectAlarm{
			DeviceAlarm: *alarm,
			ProjectID:   prj.ID,
		}
		//TODO 入库

		//上报
		prj.Emit("alarm", pa)
	}

	//初始化设备
	for _, d := range prj.Project.Devices {
		dev := GetDevice(d.ID)
		if dev == nil {
			//TODO 如果找不到设备，该怎么处理
			continue
		}
		prj.deviceNameIndex[d.Name] = dev
		prj.deviceIDIndex[d.ID] = dev
		prj.Context[d.Name] = dev.Context //两级上下文
	}

	//定时任务
	for _, job := range prj.Jobs {
		job.On("invoke", func() {
			var err error
			for _, invoke := range job.Invokes {
				err = prj.execute(&invoke)
				if err != nil {
					prj.Emit("error", err)
				}
			}

			//日志
			_ = database.History.Save(model.ProjectHistoryJob{
				ProjectID: prj.ID,
				Job:       job.String(),
				History:   "action",
				Created:   time.Now()})
		})
	}

	//初始化聚合器
	for _, agg := range prj.Aggregators {
		err := agg.Init()
		if err != nil {
			return err
		}
		for _, d := range prj.Devices {
			if d.belongSelector(&agg.Model().Selector) {
				agg.Push(d.device.Context)
			}
		}
	}

	//订阅告警
	for _, reactor := range prj.Strategies {
		reactor.On("alarm", func(alarm *model.Alarm) {
			pa := &model.ProjectAlarm{
				DeviceAlarm: model.DeviceAlarm{
					Alarm:   *alarm,
					Created: time.Now(),
				},
				ProjectID: prj.ID,
			}

			//入库
			_ = database.History.Save(model.ProjectHistoryAlarm{
				ProjectID: prj.ID,
				Code:      alarm.Code,
				Level:     alarm.Level,
				Message:   alarm.Message,
				Created:   time.Now(),
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
			history := model.ProjectHistoryReactor{
				ProjectID: prj.ID,
				Name:      reactor.Name,
				History:   "action",
				Created:   time.Now(),
			}
			if history.Name == "" {
				history.Name = reactor.Condition
			}
			_ = database.History.Save(history)
		})
	}

	return nil
}

//Start 项目启动
func (prj *Project) Start() error {
	_ = database.History.Save(model.ProjectHistory{ProjectID: prj.ID, History: "start", Created: time.Now()})

	//订阅设备的数据变化和报警
	for _, dev := range prj.Devices {
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

//Stop 项目结束
func (prj *Project) Stop() error {
	_ = database.History.Save(model.ProjectHistory{ProjectID: prj.ID, History: "stop", Created: time.Now()})

	for _, dev := range prj.Devices {
		dev.device.Off("data", prj.deviceDataHandler)
		dev.device.Off("alarm", prj.deviceAlarmHandler)
	}
	for _, job := range prj.Jobs {
		job.Stop()
	}
	return nil
}

func (prj *Project) execute(in *model.Invoke) error {

	for _, d := range prj.Devices {
		if d.belongSelector(&in.Selector) {
			err := d.device.Execute(in.Command, in.Argv)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
