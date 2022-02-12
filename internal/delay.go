package internal

type Checker interface {
	Reset()
	Check(now int64) bool
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
