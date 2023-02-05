package model

type App struct {
	Name    string `json:"name"`
	Icon    string `json:"icon,omitempty"`
	Label   string `json:"label"`
	Desc    string `json:"desc"`
	Type    string `json:"type"` //tcp unix
	Address string `json:"address"`
	Hidden  bool   `json:"hidden,omitempty"` //隐藏，适用于服务
}
