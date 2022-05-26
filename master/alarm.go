package master

import (
	"github.com/zgwit/iot-master/calc"
	"github.com/zgwit/iot-master/events"
	"github.com/zgwit/iot-master/model"
	"time"
)

//Alarm 规则
type Alarm struct {
	model.Alarm

	condition *calc.Expression
	events.EventEmitter
}

//Execute 执行
func (a *Alarm) Execute(ctx map[string]interface{}) error {

	//条件检查
	val, err := a.condition.Eval(ctx)
	if err != nil {
		return err
	}
	if !val.(bool) {
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
