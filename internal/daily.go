package interval

import "time"

type TimeRange struct {
	Start int `json:"start"`
	End   int `json:"end"`
}

func (tr *TimeRange) Check(tm *time.Time) bool {
	min := tm.Hour()*60 + tm.Minute()
	return tr.Start <= min && min <= tr.End
}

type DailyRange struct {
	Times    []TimeRange    `json:"times"`
	Weekdays []time.Weekday `json:"weekdays"`
}

func (dr *DailyRange) Check() bool {
	tm := time.Now()

	//检查时间
	has := false

	for _, tr := range dr.Times {
		if tr.Check(&tm) {
			has = true
		}
	}
	if !has {
		return false
	}

	//检查星期
	//不设置星期，则为真
	if len(dr.Weekdays) == 0 {
		return true
	}

	week := tm.Weekday()
	for _, w := range dr.Weekdays {
		if w == week {
			return true
		}
	}
	return false
}
