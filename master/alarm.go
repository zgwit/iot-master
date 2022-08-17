package master

import (
	"context"
	"github.com/PaesslerAG/gval"
	"iot-master/model"
	"iot-master/pkg/calc"
	"iot-master/pkg/events"
	"time"
)

//Alarm 规则
type Alarm struct {
	model.Alarm

	condition gval.Evaluable
	events.EventEmitter
}

func (a *Alarm) Init() (err error) {
	a.condition, err = calc.Language.NewEvaluable(a.Condition)
	return
}

//Execute 执行
func (a *Alarm) Execute(ctx map[string]interface{}) error {

	//条件检查
	val, err := a.condition.EvalBool(context.Background(), ctx)
	if err != nil {
		return err
	}
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
