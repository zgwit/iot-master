package calc

import (
	"fmt"
	"github.com/PaesslerAG/gval"
	"math"
)

func toFloat64(val interface{}) float64 {
	switch x := val.(type) {
	case float32:
		return float64(x)
	case float64:
		return x
	case int:
		return float64(x)
	case int8:
		return float64(x)
	case int16:
		return float64(x)
	case int32:
		return float64(x)
	case int64:
		return float64(x)

	case uint:
		return float64(x)
	case uint8:
		return float64(x)
	case uint16:
		return float64(x)
	case uint32:
		return float64(x)
	case uint64:
		return float64(x)

	default:
		panic(fmt.Sprintf("invalid operation: float64(%T)", x))
	}
}

func tarFunc(fun func(x float64) float64) func(x interface{}) float64 {
	return func(x interface{}) float64 {
		xx := toFloat64(x)
		return fun(xx)
	}
}

func tarFunc2(fun func(x, y float64) float64) func(x, y interface{}) float64 {
	return func(x, y interface{}) float64 {
		xx := toFloat64(x)
		yy := toFloat64(y)
		return fun(xx, yy)
	}
}

//Language 添加了数据函数的gval
var Language = gval.Full(
	gval.Constant("E", math.E),
	gval.Constant("PI", math.Pi),
	gval.Function("ABS", tarFunc(math.Abs)),
	gval.Function("CEIL", tarFunc(math.Ceil)),
	gval.Function("FLOOR", tarFunc(math.Floor)),
	gval.Function("TRUNC", tarFunc(math.Trunc)),
	gval.Function("POW", tarFunc2(math.Pow)),
	gval.Function("ROUND", tarFunc(math.Round)),
	gval.Function("SQRT", tarFunc(math.Sqrt)),
	gval.Function("CBRT", tarFunc(math.Cbrt)),
	gval.Function("EXP", tarFunc(math.Exp)),
	gval.Function("SIN", tarFunc(math.Sin)),
	gval.Function("SINH", tarFunc(math.Sinh)),
	gval.Function("ASIN", tarFunc(math.Asin)),
	gval.Function("ASINH", tarFunc(math.Asinh)),
	gval.Function("COS", tarFunc(math.Cos)),
	gval.Function("COSH", tarFunc(math.Cosh)),
	gval.Function("ACOS", tarFunc(math.Acos)),
	gval.Function("ACOSH", tarFunc(math.Acosh)),
	gval.Function("TAN", tarFunc(math.Tan)),
	gval.Function("TANH", tarFunc(math.Tanh)),
	gval.Function("ATAN", tarFunc(math.Atan)),
	gval.Function("ATANH", tarFunc(math.Atanh)),
	gval.Function("LN", tarFunc(math.Log)),
	gval.Function("LOG", tarFunc(math.Log10)),
)
