package master

import (
	"errors"
	"fmt"
	"github.com/zgwit/iot-master/aggregator"
	"github.com/zgwit/iot-master/calc"
	"github.com/zgwit/iot-master/database"
	"github.com/zgwit/iot-master/events"
	"github.com/zgwit/iot-master/model"
	"github.com/zgwit/storm/v3"
	"strings"
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

func (d *ProjectDevice) belongTargets(targets []string) bool {
	for _, target := range targets {
		if target == d.Name {
			return true
		}
		for _, tag := range d.device.Tags {
			//strings.EqualFold
			if target == tag {
				return true
			}
		}
	}
	return false
}

//Project 项目
type Project struct {
	model.Project
	events.EventEmitter

	Devices []*ProjectDevice

	aggregators []aggregator.Aggregator
	jobs        []*Job
	strategies  []*Strategy

	deviceNameIndex map[string]*Device
	deviceIDIndex   map[int]*Device

	deviceDataHandler  func(data calc.Context)
	deviceAlarmHandler func(alarm *model.DeviceAlarm)
}

func NewProject(m *model.Project) (*Project, error) {
	prj := &Project{
		Project:         *m,
		Devices:         make([]*ProjectDevice, 0),
		aggregators:     make([]aggregator.Aggregator, 0),
		jobs:            make([]*Job, 0),
		strategies:      make([]*Strategy, 0),
		deviceNameIndex: make(map[string]*Device),
		deviceIDIndex:   make(map[int]*Device),
	}

	//加载模板
	if prj.TemplateID != "" {
		var template model.Template
		err := database.Master.One("ID", prj.TemplateID, &template)
		if err == storm.ErrNotFound {
			return nil, errors.New("找不到模板")
		} else if err != nil {
			return nil, err
		}
		prj.ProjectContent = template.ProjectContent
	}

	err := prj.initDevices()
	if err != nil {
		return nil, err
	}

	err = prj.initAggregators()
	if err != nil {
		return nil, err
	}

	err = prj.initHandler()
	if err != nil {
		return nil, err
	}

	err = prj.initJobs()
	if err != nil {
		return nil, err
	}

	err = prj.initStrategies()
	if err != nil {
		return nil, err
	}

	return prj, nil
}

func (prj *Project) initDevices() error {
	if prj.Devices == nil {
		return nil
	}
	for _, d := range prj.Project.Devices {
		dev := GetDevice(d.ID)
		if dev == nil {
			//如果找不到设备，该怎么处理
			return fmt.Errorf("device %d not found", d.ID)
		}
		prj.deviceNameIndex[d.Name] = dev
		prj.deviceIDIndex[d.ID] = dev
		prj.Context[d.Name] = dev.Context //两级上下文

		prj.Devices = append(prj.Devices, &ProjectDevice{ProjectDevice: *d})
	}
	return nil
}

func (prj *Project) initAggregators() error {
	if prj.Aggregators == nil {
		return nil
	}
	for _, v := range prj.Aggregators {
		agg, err := aggregator.New(v)
		if err != nil {
			return err
		}
		err = agg.Init()
		if err != nil {
			return err
		}
		for _, d := range prj.Devices {
			if d.belongTargets(agg.Model().Targets) {
				agg.Push(d.device.Context)
			}
		}
		prj.aggregators = append(prj.aggregators, agg)
	}
	return nil
}

func (prj *Project) initJobs() error {
	if prj.Jobs == nil {
		return nil
	}
	for _, v := range prj.Jobs {
		job := &Job{Job: *v}
		job.On("invoke", func() {
			var err error
			for _, invoke := range job.Invokes {
				err = prj.execute(&invoke)
				if err != nil {
					prj.Emit("error", err)
				}
			}

			//日志
			prj.createEvent("执行定时任务：" + job.String())
		})
		prj.jobs = append(prj.jobs, job)
	}
	return nil
}

func (prj *Project) initStrategies() error {
	if prj.Strategies == nil {
		return nil
	}
	for _, v := range prj.Strategies {
		strategy := &Strategy{Strategy: *v}
		strategy.On("alarm", func(alarm *model.Alarm) {
			pa := &model.ProjectAlarm{ProjectID: prj.ID, Alarm: *alarm}

			//入库
			_ = database.History.Save(pa)

			//事件
			prj.createEvent("告警：" + alarm.Message)

			//上报
			prj.Emit("alarm", pa)
		})

		strategy.On("invoke", func() {
			for _, invoke := range strategy.Invokes {
				err := prj.execute(invoke)
				if err != nil {
					prj.Emit("error", err)
				}
			}

			//保存历史
			prj.createEvent("执行控制策略：" + strategy.Name)
		})
		prj.strategies = append(prj.strategies, strategy)
	}
	return nil
}

//initHandler 项目初始化
func (prj *Project) initHandler() error {
	//设备数据变化的处理函数
	prj.deviceDataHandler = func(data calc.Context) {
		//数据变化后，更新计算
		for _, agg := range prj.aggregators {
			val, err := agg.Evaluate()
			if err != nil {
				prj.Emit("error", err)
			} else {
				prj.Context[agg.Model().As] = val
			}
		}

		//处理响应
		for _, strategy := range prj.strategies {
			err := strategy.Execute(prj.Context)
			if err != nil {
				prj.Emit("error", err)
			}
		}
	}

	//设备告警的处理函数
	prj.deviceAlarmHandler = func(alarm *model.DeviceAlarm) {
		pa := &model.ProjectAlarm{ProjectID: prj.ID, Alarm: alarm.Alarm}

		//历史入库
		_ = database.History.Save(pa)

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

	return nil
}

func (prj *Project) createEvent(event string) {
	_ = database.History.Save(model.ProjectEvent{ProjectID: prj.ID, Event: event})
}

//Start 项目启动
func (prj *Project) Start() error {
	prj.createEvent("启动")

	//订阅设备的数据变化和报警
	for _, dev := range prj.Devices {
		dev.device.On("data", prj.deviceDataHandler)
		dev.device.On("alarm", prj.deviceAlarmHandler)
	}

	//定时任务
	for _, job := range prj.jobs {
		err := job.Start()
		if err != nil {
			return err
		}
	}
	return nil
}

//Stop 项目结束
func (prj *Project) Stop() error {
	prj.createEvent("关闭")

	for _, dev := range prj.Devices {
		dev.device.Off("data", prj.deviceDataHandler)
		dev.device.Off("alarm", prj.deviceAlarmHandler)
	}
	for _, job := range prj.jobs {
		job.Stop()
	}
	return nil
}

func (prj *Project) execute(in *model.Invoke) error {

	for _, d := range prj.Devices {
		if d.belongTargets(in.Targets) {
			err := d.device.Execute(in.Command, in.Argv)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
