package master

import (
	"iot-master/helper"
	"iot-master/model"
	"iot-master/pkg/events"
	"time"
)

//Validator 规则
type Validator struct {
	model.Validator

	events.EventEmitter
}

//Execute 执行
func (a *Validator) Execute(ctx map[string]interface{}) error {

	//条件检查
	value := helper.ToFloat64(ctx[a.Value])
	val := a.Compare.Eval(value, a.Value1, a.Value2)
	if !val {
		a.DelayChecker.Reset()
		a.RepeatChecker.Reset()
		return nil
	}

	//时间检查
	if !a.DailyChecker.Check() {
		a.DelayChecker.Reset()
		a.RepeatChecker.Reset()
		return nil
	}

	now := time.Now().UnixMicro()
	//时间检查
	if !a.DelayChecker.Check(now) {
		a.RepeatChecker.Check(now)
		return nil
	}

	//重复检查
	if !a.RepeatChecker.Check(now) {
		return nil
	}

	//产生报警
	a.Emit("alarm", &a.AlarmContent)

	return nil
}
