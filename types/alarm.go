package types

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
