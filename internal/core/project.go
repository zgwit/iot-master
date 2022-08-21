package core

import (
	"fmt"
	"github.com/zgwit/iot-master/model"
	"github.com/zgwit/iot-master/pkg/calc"
	"github.com/zgwit/iot-master/pkg/events"
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
		for _, tag := range d.device.product.Tags {
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

	deviceNameIndex map[string]*Device
	deviceIdIndex   map[uint64]*Device

	deviceDataHandler func(data map[string]interface{})

	running bool
}

func NewProject(m *model.Project) (*Project, error) {
	prj := &Project{
		Project:         *m,
		Context:         make(map[string]interface{}),
		Devices:         make([]*ProjectDevice, 0),
		deviceNameIndex: make(map[string]*Device),
		deviceIdIndex:   make(map[uint64]*Device),
	}

	err := prj.initDevices()
	if err != nil {
		return nil, err
	}

	err = prj.initHandler()
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

//initHandler 项目初始化
func (prj *Project) initHandler() error {
	//设备数据变化的处理函数
	prj.deviceDataHandler = func(data map[string]interface{}) {

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

	//订阅设备的数据变化和报警
	for _, dev := range prj.Devices {
		dev.device.On("data", prj.deviceDataHandler)
	}

	prj.running = true

	return nil
}

//Stop 项目结束
func (prj *Project) Stop() error {
	for _, dev := range prj.Devices {
		dev.device.Off("data", prj.deviceDataHandler)
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
