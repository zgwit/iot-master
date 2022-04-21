package model

//Command 命令
type Command struct {
	Name string `json:"name"`
	//Label string `json:"label,omitempty"`

	Directives []Directive `json:"directives"`
}

//Directive 指令
type Directive struct {
	Address string  `json:"address"`
	Value   float64 `json:"value,omitempty"`
	Delay   int64   `json:"delay,omitempty"`

	//使用表达式
	Expression string `json:"expression,omitempty"`
}

//Invoke 执行
type Invoke struct {
	Targets   []string `json:"targets,omitempty"`
	Command   string   `json:"command"`
	Arguments []string `json:"arguments"`
}
