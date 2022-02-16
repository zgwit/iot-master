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
func (s *Strategy) Execute(ctx calc.Context) error {

	//条件检查
	val, err := s.condition.Evaluate(ctx)
	if err != nil {
		return err
	}
	if !val.(bool) {
		s.Delay.Reset()
		s.Repeat.Reset()
		return nil
	}

	//时间检查
	if s.Daily != nil && !s.Daily.Check() {
		s.Delay.Reset()
		s.Repeat.Reset()
		return nil
	}

	now := time.Now().UnixMicro()
	//时间检查
	if s.Delay != nil && !s.Delay.Check(now) {
		s.Repeat.Check(now)
		return nil
	}

	//重复检查
	if s.Repeat != nil && !s.Repeat.Check(now) {
		return nil
	}

	//产生报警
	if s.Alarm != nil {
		s.Emit("alarm", &s.Alarm)
	}

	//执行响应
	//for _, i := range s.Invokes {
	//	s.events.Publish("invoke", i)
	//}
	if s.Invokes != nil && len(s.Invokes) > 0 {
		s.Emit("invoke")
	}

	return nil
}
