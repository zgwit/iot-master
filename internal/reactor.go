package interval

import "time"

type Rector struct {
	Disabled bool `json:"disabled"`

	//条件
	Condition Condition `json:"condition"`

	//重复日
	Daily DailyRange `json:"daily"`

	//延迟报警
	Delay DelayChecker `json:"delay"`

	//重复报警
	Repeat RepeatChecker `json:"repeat"`

	//执行命名
	Invokes []Invoke `json:"invokes"`
}

func (r *Rector) Execute() error {

	//条件检查
	if !r.Condition.Evaluate() {
		r.Delay.Reset()
		r.Repeat.Reset()
		return nil
	}

	//时间检查
	if !r.Daily.Check() {
		r.Delay.Reset()
		r.Repeat.Reset()
		return nil
	}

	now := time.Now().UnixMicro()
	//时间检查
	if !r.Delay.Check(now) {
		r.Repeat.Check(now)
		return nil
	}

	//重复检查
	if !r.Repeat.Check(now) {
		return nil
	}

	//执行响应
	for _, i := range r.Invokes {
		if err := i.Execute(); err != nil {
			return err
		}
	}

	return nil
}
