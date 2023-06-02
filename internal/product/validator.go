package product

import (
	"github.com/PaesslerAG/gval"
	"github.com/zgwit/iot-master/v3/model"
)

type Validator struct {
	*model.Validator

	Expression gval.Evaluable

	//again      bool
	Start int64 //开始时间s
	Count uint  //报警次数
}
