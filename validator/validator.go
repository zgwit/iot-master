package validator

import (
	"context"
	"github.com/PaesslerAG/gval"
	"github.com/zgwit/iot-master/v3/model"
	"github.com/zgwit/iot-master/v3/pkg/calc"
	"github.com/zgwit/iot-master/v3/pkg/log"
	"time"
)

type Validator struct {
	*model.ModValidator

	Expression gval.Evaluable

	//again      bool
	start int64 //开始时间s
	count uint  //报警次数
	//reported bool
}

// New 新建
func New(m *model.ModValidator) (v *Validator, err error) {
	v = &Validator{ModValidator: m}
	v.Expression, err = calc.New(m.Expression)
	return
}

func (v *Validator) Validate(values map[string]any) bool {
	ret, err := v.Expression.EvalBool(context.Background(), values)
	if err != nil {
		log.Error(err)
		return false
	}

	if !ret {
		//约束OK，检查下一个
		v.start = 0
		v.count = 0
		return false
	}

	//起始时间
	now := time.Now().Unix()
	if v.start == 0 {
		v.start = now
	}

	//延迟报警
	if v.Delay > 0 {
		if now < v.start+int64(v.Delay) {
			return false
		}
	}

	if v.count > 0 {
		//重复报警
		if !v.Repeat {
			return false
		}

		//超过最大次数
		if v.RepeatTotal > 0 && v.count >= v.RepeatTotal {
			return false
		}

		//还没到时间
		if now < v.start+int64(v.RepeatDelay) {
			return false
		}

		v.start = now
	}

	v.count++

	return true
}
