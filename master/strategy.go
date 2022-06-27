package master

import (
	"context"
	"github.com/PaesslerAG/gval"
	"iot-master/calc"
	"iot-master/events"
	"iot-master/model"
	"time"
)

//Strategy 规则
type Strategy struct {
	model.Strategy

	condition gval.Evaluable
	events.EventEmitter
}

func (s *Strategy) Init() (err error) {
	s.condition, err = calc.Language.NewEvaluable(s.Condition)
	return
}

//Execute 执行
func (s *Strategy) Execute(ctx map[string]interface{}) error {

	//条件检查
	val, err := s.condition.EvalBool(context.Background(), ctx)
	if err != nil {
		return err
	}
	if !val {
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
