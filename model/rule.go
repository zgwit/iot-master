package model

//Rule 响应，待改名
type Rule struct {
	Disabled bool `json:"disabled"`

	//名称
	Name string `json:"name"`

	//条件
	Condition string `json:"condition"`

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
}
