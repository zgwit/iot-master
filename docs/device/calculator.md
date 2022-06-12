有些传感器的数据需要进一步计算才能获取，比如：溶氧值需要结合盐度、温度等

公式引擎

https://github.com/PaesslerAG/gval

## 内置函数

| 函数    | golang函数映射 | 说明          |
|-------|------------|-------------|
| E常量   | math.E     | 欧拉常数 2.718  |
| PI常量  | math.Pi    | 圆周率 3.14159 |
| ABS   | math.Abs   | 绝对值         |
| CEIL  | math.Ceil  | 上限整数        |
| FLOOR | math.Floor | 下限整数        |
| TRUNC | math.Trunc | 去掉小数位       |
| POW   | math.Pow   | 幂次方         |
| ROUND | math.Round | 四舍五入        |
| SQRT  | math.Sqrt  | 平方根         |
| CBRT  | math.Cbrt  | 立方根         |
| EXP   | math.Exp   | 欧拉常数的幂次方    |
| SIN   | math.Sin   | 正弦          |
| SINH  | math.Sinh  | 双曲正弦        |
| ASIN  | math.Asin  | 反正弦         |
| ASINH | math.Asinh | 反双曲正弦       |
| COS   | math.Cos   | 余弦          |
| COSH  | math.Cosh  | 双曲余弦        |
| ACOS  | math.Acos  | 反余弦         |
| ACOSH | math.Acosh | 反双曲余弦       |
| TAN   | math.Tan   | 正切          |
| TANH  | math.Tanh  | 双曲正切        |
| ATAN  | math.Atan  | 反切          |
| ATANH | math.Atanh | 反双曲正切       |
| LN    | math.Log   | 对数          |
| LOG   | math.Log10 | 10的自然对数     |
| HYPOT | math.Hypot | 平方和的平方根     |
| MAX   | math.Max   | 最大          |
| MIN   | math.Min   | 最小          |
