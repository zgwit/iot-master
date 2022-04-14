package model

//Invoke 执行
type Invoke struct {
	Targets []string  `json:"targets,omitempty"`
	Command string    `json:"command"`
	Argv    []float64 `json:"argv"`
}
