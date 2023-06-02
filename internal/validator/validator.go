package validator

import (
	"github.com/PaesslerAG/gval"
	"github.com/zgwit/iot-master/v3/model"
	"github.com/zgwit/iot-master/v3/pkg/calc"
)

type Validator struct {
	*model.ModValidator

	Expression gval.Evaluable

	//again      bool
	Start int64 //开始时间s
	Count uint  //报警次数
}

// New 新建
func New(m *model.ModValidator) (v *Validator, err error) {
	v = &Validator{ModValidator: m}
	v.Expression, err = calc.New(m.Expression)
	return
}
