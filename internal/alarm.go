package interval

import "time"

type Alarm struct {
	Disabled bool `json:"disabled"`

	//告警
	Code    string `json:"code"`
	Level   int    `json:"level"`
	Message string `json:"message"`

	//条件
	Condition Condition `json:"condition"`

	//重复日
	Daily DailyRange `json:"daily"`

	//延迟报警
	Delay  DelayChecker  `json:"delay"`

	//重复报警
	Repeat RepeatChecker `json:"repeat"`
}

func (a *Alarm) Execute() {

	//条件检查
	if !a.Condition.Evaluate() {
		a.Delay.Reset()
		a.Repeat.Reset()
		return
	}

	//时间检查
	if !a.Daily.Check() {
		a.Delay.Reset()
		a.Repeat.Reset()
		return
	}


	now := time.Now().UnixMicro()
	//时间检查
	if !a.Delay.Check(now) {
		a.Repeat.Check(now)
		return
	}

	//重复检查
	if !a.Repeat.Check(now) {
		return
	}

	//TODO 产生报警


}
