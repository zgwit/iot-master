package master

import (
	"fmt"
	"iot-master/aggregator"
	"iot-master/calc"
	"iot-master/db"
	"iot-master/events"
	"iot-master/model"
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

	Context map[string]interface{}
	Devices []*ProjectDevice

	aggregators []aggregator.Aggregator
	alarms      []*Alarm
	jobs        []*Job
	strategies  []*Strategy
	scripts     []*Script

	deviceNameIndex map[string]*Device
	deviceIdIndex   map[int64]*Device

	deviceDataHandler  func(data map[string]interface{})
	deviceAlarmHandler func(alarm *model.DeviceAlarm)

	running bool
}

func NewProject(m *model.Project) (*Project, error) {
	prj := &Project{
		Project:         *m,
		Context:         make(map[string]interface{}),
		Devices:         make([]*ProjectDevice, 0),
		aggregators:     make([]aggregator.Aggregator, 0),
		alarms:          make([]*Alarm, 0),
		jobs:            make([]*Job, 0),
		strategies:      make([]*Strategy, 0),
		scripts:         make([]*Script, 0),
		deviceNameIndex: make(map[string]*Device),
		deviceIdIndex:   make(map[int64]*Device),
	}

	//加载模板
	if prj.TemplateId != "" {
		var template model.Template
		has, err := db.Engine.ID(prj.TemplateId).Get(&template)
		if err != nil {
			return nil, err
		}
		if !has {
			return nil, fmt.Errorf("找不到模板 %s", prj.TemplateId)
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

	err = prj.initAlarms()
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
		dev := GetDevice(d.Id)
		if dev == nil {
			//如果找不到设备，该怎么处理
			return fmt.Errorf("device %d not found", d.Id)
		}

		prj.deviceNameIndex[d.Name] = dev
		prj.deviceIdIndex[d.Id] = dev
		prj.Context[d.Name] = dev.Context //两级上下文

		prj.Devices = append(prj.Devices, &ProjectDevice{ProjectDevice: *d, device: dev})
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
			CreateProjectEvent(prj.Id, "执行定时任务："+job.String())
		})
		prj.jobs = append(prj.jobs, job)
	}
	return nil
}

func (prj *Project) initAlarms() error {
	if prj.Alarms == nil {
		return nil
	}
	for _, v := range prj.Alarms {
		a := &Alarm{Alarm: *v}
		err := a.Init()
		if err != nil {
			return err
		}
		a.On("alarm", func(alarm *model.AlarmContent) {
			pa := &model.ProjectAlarm{ProjectId: prj.Id, AlarmContent: *alarm}

			//入库
			_, _ = db.Engine.InsertOne(pa)

			//事件
			CreateProjectEvent(prj.Id, "告警："+alarm.Message)

			//上报
			prj.Emit("alarm", pa)
		})
		prj.alarms = append(prj.alarms, a)
	}
	return nil
}

func (prj *Project) initStrategies() error {
	if prj.Strategies == nil {
		return nil
	}
	for _, v := range prj.Strategies {
		strategy := &Strategy{Strategy: *v}
		err := strategy.Init()
		if err != nil {
			return err
		}

		strategy.On("invoke", func() {
			for _, invoke := range strategy.Invokes {
				err := prj.execute(invoke)
				if err != nil {
					prj.Emit("error", err)
				}
			}

			//保存历史
			CreateProjectEvent(prj.Id, "执行控制策略："+strategy.Name)
		})
		prj.strategies = append(prj.strategies, strategy)
	}
	return nil
}

func (prj *Project) initScripts() error {
	if prj.Scripts == nil {
		return nil
	}
	for _, v := range prj.Scripts {
		script := &Script{Script: *v}
		err := script.Init(prj.Context)
		if err != nil {
			return err
		}
		prj.scripts = append(prj.scripts, script)
	}
	return nil
}

//initHandler 项目初始化
func (prj *Project) initHandler() error {
	//设备数据变化的处理函数
	prj.deviceDataHandler = func(data map[string]interface{}) {
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

		//处理脚本
		for _, script := range prj.scripts {
			err := script.OnData(prj.Context)
			if err != nil {
				prj.Emit("error", err)
			}
		}
	}

	//设备告警的处理函数
	prj.deviceAlarmHandler = func(alarm *model.DeviceAlarm) {
		pa := &model.ProjectAlarm{ProjectId: prj.Id, AlarmContent: alarm.AlarmContent}

		//历史入库
		_, _ = db.Engine.InsertOne(pa)

		//上报
		prj.Emit("alarm", pa)
	}

	//初始化设备
	for _, d := range prj.Project.Devices {
		dev := GetDevice(d.Id)
		if dev == nil {
			//TODO 如果找不到设备，该怎么处理
			continue
		}
		prj.deviceNameIndex[d.Name] = dev
		prj.deviceIdIndex[d.Id] = dev
		prj.Context[d.Name] = dev.Context //两级上下文
	}

	return nil
}

//Start 项目启动
func (prj *Project) Start() error {
	CreateProjectEvent(prj.Id, "启动")

	//订阅设备的数据变化和报警
	for _, dev := range prj.Devices {
		dev.device.On("data", prj.deviceDataHandler)
		dev.device.On("alarm", prj.deviceAlarmHandler)
	}

	//定时任务
	for _, job := range prj.jobs {
		err := job.Start()
		if err != nil {
			//TODO 需要关闭
			//_ = prj.Stop()
			return err
		}
	}

	prj.running = true

	return nil
}

//Stop 项目结束
func (prj *Project) Stop() error {
	CreateProjectEvent(prj.Id, "关闭")

	for _, dev := range prj.Devices {
		dev.device.Off("data", prj.deviceDataHandler)
		dev.device.Off("alarm", prj.deviceAlarmHandler)
	}
	for _, job := range prj.jobs {
		job.Stop()
	}

	prj.running = false

	return nil
}

func (prj *Project) Running() bool {
	return prj.running
}

func (prj *Project) Set(name string, value interface{}) error {
	//prj.Context[name] = value
	index := strings.Index(name, ".")
	if index == -1 {
		prj.Context[name] = value
	} else {
		dev := name[:index]
		key := name[index+1:]
		if d, ok := prj.deviceNameIndex[dev]; ok {
			return d.Set(key, value)
		} //else return error ??
	}
	return nil
}

func (prj *Project) execute(in *model.Invoke) error {
	args := make([]interface{}, 0)
	for _, d := range in.Arguments {
		//tp := reflect.TypeOf(d).Kind()
		//if tp == reflect.String {
		//} else if tp == reflect.Float64 {
		//	args = append(args, d.(float64))
		//}
		val, err := calc.Language.Evaluate(d, prj.Context)
		if err != nil {
			return err
		}
		args = append(args, val)
	}

	for _, d := range prj.Devices {
		if d.belongTargets(in.Targets) {
			err := d.device.Execute(in.Command, args)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
