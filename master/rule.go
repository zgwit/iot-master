package master

import (
	"github.com/zgwit/iot-master/calc"
	"github.com/zgwit/iot-master/events"
	"github.com/zgwit/iot-master/model"
	"time"
)

//Rule 规则
type Rule struct {
	model.Rule

	condition *calc.Expression
	events.EventEmitter
}

//Execute 执行
func (r *Rule) Execute(ctx calc.Context) error {

	//条件检查
	val, err := r.condition.Evaluate(ctx)
	if err != nil {
		return err
	}
	if !val.(bool) {
		r.Delay.Reset()
		r.Repeat.Reset()
		return nil
	}

	//时间检查
	if r.Daily != nil && !r.Daily.Check() {
		r.Delay.Reset()
		r.Repeat.Reset()
		return nil
	}

	now := time.Now().UnixMicro()
	//时间检查
	if r.Delay != nil && !r.Delay.Check(now) {
		r.Repeat.Check(now)
		return nil
	}

	//重复检查
	if r.Repeat != nil && !r.Repeat.Check(now) {
		return nil
	}

	//产生报警
	if r.Alarm != nil {
		r.Emit("alarm", &r.Alarm)
	}

	//执行响应
	//for _, i := range r.Invokes {
	//	r.events.Publish("invoke", i)
	//}
	if r.Invokes != nil && len(r.Invokes) > 0 {
		r.Emit("invoke")
	}

	return nil
}
