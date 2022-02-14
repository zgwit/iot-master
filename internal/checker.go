package internal

import "time"

type TimeRange struct {
	Start int `json:"start"`
	End   int `json:"end"`
}

func (tr *TimeRange) Check(tm *time.Time) bool {
	min := tm.Hour()*60 + tm.Minute()
	return tr.Start <= min && min <= tr.End
}

type DailyChecker struct {
	Times    []TimeRange    `json:"times"`
	Weekdays []time.Weekday `json:"weekdays"`
}

func (dr *DailyChecker) Check() bool {
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

type DelayChecker struct {
	Delay int64 `json:"delay"`

	start int64
}

func (d *DelayChecker) Reset() {
	d.start = 0
}

func (d *DelayChecker) Check(now int64) bool {
	if d.Delay <= 0 {
		return true
	}

	if d.start == 0 {
		d.start = now
		return false
	}

	return d.start+d.Delay < now
}

type RepeatChecker struct {
	Interval int64 `json:"interval"`
	Total    int   `json:"total,omitempty"`

	last int64

	raised     bool
	resetTimes int
}

func (d *RepeatChecker) Reset() {
	d.raised = false
	d.resetTimes = 0
}

func (d *RepeatChecker) Check(now int64) bool {
	//初次
	if !d.raised {
		d.raised = true
		d.last = now
		d.resetTimes = 0

		return true
	}

	//重置间隔
	if d.Interval <= 0 {
		return false
	}

	//最大重置次数限制
	if d.Total > 0 && d.resetTimes > d.Total {
		return false
	}

	//如果还没到重置时间，则不提醒
	if d.last+d.Interval > now {
		return false
	}

	//重置
	d.last = now
	d.resetTimes++
	return true
}
