package master

import (
	"github.com/zgwit/iot-master/calc"
	"github.com/zgwit/iot-master/events"
	"github.com/zgwit/iot-master/model"
	"time"
)

//Reactor 响应
type Reactor struct {
	model.Reactor

	condition *calc.Expression
	events.EventEmitter
}

//Execute 执行
func (a *Reactor) Execute(ctx calc.Context) error {

	//条件检查
	val, err := a.condition.Evaluate(ctx)
	if err != nil {
		return err
	}
	if !val.(bool) {
		a.Delay.Reset()
		a.Repeat.Reset()
		return nil
	}

	//时间检查
	if a.Daily != nil && !a.Daily.Check() {
		a.Delay.Reset()
		a.Repeat.Reset()
		return nil
	}

	now := time.Now().UnixMicro()
	//时间检查
	if a.Delay != nil && !a.Delay.Check(now) {
		a.Repeat.Check(now)
		return nil
	}

	//重复检查
	if a.Repeat != nil && !a.Repeat.Check(now) {
		return nil
	}

	//产生报警
	if a.Alarm != nil {
		a.Emit("alarm", &a.Alarm)
	}

	//执行响应
	//for _, i := range a.Invokes {
	//	a.events.Publish("invoke", i)
	//}
	if a.Invokes != nil && len(a.Invokes) > 0 {
		a.Emit("invoke")
	}

	return nil
}
