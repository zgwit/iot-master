package model

//Validator 数据校验
type Validator struct {
	Disabled bool `json:"disabled"`

	//条件
	Condition string `json:"condition"`

	//重复日
	DailyChecker

	//延迟报警
	DelayChecker

	//重复报警
	RepeatChecker

	//产生告警
	Alarm
}
