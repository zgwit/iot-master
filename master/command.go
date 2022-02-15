package master

import "github.com/zgwit/iot-master/master/select"

//Command 命令
type Command struct {
	Name  string `json:"name"`
	Label string `json:"label,omitempty"`
	Argc  int    `json:"argc,omitempty"`

	Directives []*Directive `json:"directives"`
}

//Directive 指令
type Directive struct {
	Point string  `json:"point"`
	Value float64 `json:"value,omitempty"`
	Delay int64   `json:"delay"`

	//使用参数
	Arg int `json:"arg,omitempty"` //0:默认参数 1 2 3

	//使用表达式
	Expression string `json:"expression,omitempty"`
}

//Invoke 执行
type Invoke struct {
	Command string    `json:"command"`
	Argv    []float64 `json:"argv"`

	//目标设备（只在Project中使用）
	Select _select.Select `json:"select,omitempty"`
}
