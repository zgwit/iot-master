package model

//Invoke 执行
type Invoke struct {
	Command string    `json:"command"`
	Argv    []float64 `json:"argv"`

	//目标设备（只在Project中使用）
	Selector Selector `json:"selector,omitempty"`
}
