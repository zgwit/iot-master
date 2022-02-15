package master

import (
	"github.com/zgwit/iot-master/events"
	"github.com/zgwit/iot-master/master/calc"
	"time"
)

//Alarm 告警内容
type Alarm struct {
	Code    string `json:"code"`
	Level   int    `json:"level"`
	Message string `json:"message"`
}

//DeviceAlarm 设备告警
type DeviceAlarm struct {
	Alarm
	DeviceId int       `json:"device_id"`
	Created  time.Time `json:"time"`
}

//ProjectAlarm 项目告警
type ProjectAlarm struct {
	DeviceAlarm
	ProjectId int `json:"project_id"`
}

//Reactor 响应
type Reactor struct {
	Disabled bool `json:"disabled"`

	//名称
	Name string `json:"name"`

	//条件
	Condition string `json:"condition"`
	condition *calc.Expression

	//重复日
	Daily *DailyChecker `json:"daily,omitempty"`

	//延迟报警
	Delay *DelayChecker `json:"delay,omitempty"`

	//重复报警
	Repeat *RepeatChecker `json:"repeat,omitempty"`

	//产生告警
	Alarm *Alarm `json:"alarm,omitempty"`

	//执行命名
	Invokes []*Invoke `json:"invokes,omitempty"`

	events.EventEmitter
}

//Execute 执行
func (a *Reactor) Execute(ctx calc.Context) error {

	//条件检查
	val, err := a.condition.Evaluate(ctx)
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

	//产生报警
	if a.Alarm != nil {
		a.Emit("alarm", &a.Alarm)
	}

	//执行响应
	//for _, i := range a.Invokes {
	//	a.events.Publish("invoke", i)
	//}
	if a.Invokes != nil && len(a.Invokes) > 0 {
		a.Emit("invoke")
	}

	return nil
}
