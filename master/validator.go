package master

import (
	"github.com/zgwit/iot-master/calc"
	"github.com/zgwit/iot-master/events"
	"github.com/zgwit/iot-master/model"
	"time"
)

//Validator 规则
type Validator struct {
	model.Validator

	condition *calc.Expression
	events.EventEmitter
}

//Execute 执行
func (s *Validator) Execute(ctx calc.Context) error {

	//条件检查
	val, err := s.condition.Evaluate(ctx)
	if err != nil {
		return err
	}
	if !val.(bool) {
		s.DelayChecker.Reset()
		s.RepeatChecker.Reset()
		return nil
	}

	//时间检查
	if !s.DailyChecker.Check() {
		s.DelayChecker.Reset()
		s.RepeatChecker.Reset()
		return nil
	}

	now := time.Now().UnixMicro()
	//时间检查
	if !s.DelayChecker.Check(now) {
		s.RepeatChecker.Check(now)
		return nil
	}

	//重复检查
	if !s.RepeatChecker.Check(now) {
		return nil
	}

	//产生报警
	s.Emit("alarm", &s.Alarm)

	return nil
}
