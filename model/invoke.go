package model

//Invoke 执行
type Invoke struct {
	Targets   []string      `json:"targets,omitempty"`
	Command   string        `json:"command"`
	Arguments []interface{} `json:"arguments"`
}
