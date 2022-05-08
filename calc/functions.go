package calc

import (
	"fmt"
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

func CreateContext() Context {
	 return map[string]interface{}{
		"E": math.E,
		"PI": math.Pi,
		"ABS": tarFunc(math.Abs),
		"CEIL": tarFunc(math.Ceil),
		"FLOOR": tarFunc(math.Floor),
		"TRUNC": tarFunc(math.Trunc),
		"POW": tarFunc2(math.Pow),
		"ROUND": tarFunc(math.Round),
		"SQRT": tarFunc(math.Sqrt),
		"CBRT": tarFunc(math.Cbrt),
		"EXP": tarFunc(math.Exp),
		"SIN": tarFunc(math.Sin),
		"SINH": tarFunc(math.Sinh),
		"ASIN": tarFunc(math.Asin),
		"ASINH": tarFunc(math.Asinh),
		"COS": tarFunc(math.Cos),
		"COSH": tarFunc(math.Cosh),
		"ACOS": tarFunc(math.Acos),
		"ACOSH": tarFunc(math.Acosh),
		"TAN": tarFunc(math.Tan),
		"TANH": tarFunc(math.Tanh),
		"ATAN": tarFunc(math.Atan),
		"ATANH": tarFunc(math.Atanh),
		"LN": tarFunc(math.Log),
		"LOG": tarFunc(math.Log10),
	}
}
