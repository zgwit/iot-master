package interval

import (
	"github.com/asaskevich/EventBus"
	"time"
)

type Alarm struct {
	Code    string `json:"code"`
	Level   int    `json:"level"`
	Message string `json:"message"`
}

type DeviceAlarm struct {
	Alarm
	DeviceId int       `json:"device_id"`
	Created  time.Time `json:"time"`
}

type ProjectAlarm struct {
	DeviceAlarm
	ProjectId int `json:"project_id"`
}

type Reactor struct {
	Disabled bool `json:"disabled"`

	//条件
	Condition string `json:"condition"`
	condition *Expression

	//重复日
	Daily *DailyRange `json:"daily,omitempty"`

	//延迟报警
	Delay *DelayChecker `json:"delay,omitempty"`

	//重复报警
	Repeat *RepeatChecker `json:"repeat,omitempty"`

	//产生告警
	Alarm *Alarm `json:"alarm,omitempty"`

	//执行命名
	Invokes []Invoke `json:"invokes,omitempty"`

	events EventBus.Bus
}

func (a *Reactor) Init(){
	a.events = EventBus.New()
}

func (a *Reactor) Execute() error {

	//条件检查
	val, err := a.condition.Evaluate()
	if err != nil {
		return err
	}
	if !val.(bool) {
		a.Delay.Reset()
		a.Repeat.Reset()
		return nil
	}

	//时间检查
	if a.Daily != nil && !a.Daily.Check() {
		a.Delay.Reset()
		a.Repeat.Reset()
		return nil
	}

	now := time.Now().UnixMicro()
	//时间检查
	if a.Delay != nil && !a.Delay.Check(now) {
		a.Repeat.Check(now)
		return nil
	}

	//重复检查
	if a.Repeat != nil && !a.Repeat.Check(now) {
		return nil
	}

	//TODO 产生报警
	if a.Alarm != nil {
		//TODO 使用事件机制？？？
	}

	//执行响应
	for _, i := range a.Invokes {
		a.events.Publish("invoke", i)
	}

	return nil
}