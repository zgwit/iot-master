package master

import (
	"github.com/zgwit/iot-master/calc"
	"github.com/zgwit/iot-master/events"
	"github.com/zgwit/iot-master/model"
	"time"
)

//Strategy 规则
type Strategy struct {
	model.Strategy

	condition *calc.Expression
	events.EventEmitter
}

//Execute 执行
func (s *Strategy) Execute(ctx map[string]interface{}) error {

	//条件检查
	val, err := s.condition.Eval(ctx)
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

	//执行响应
	s.Emit("invoke")

	return nil
}
