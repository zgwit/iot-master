package interval

type Command struct {
	Name  string `json:"name"`
	Label string `json:"label,omitempty"`
	Argc  int    `json:"argc,omitempty"`

	Directives []Directive `json:"directives"`
}

type Directive struct {
	Point string  `json:"point"`
	Value float64 `json:"value,omitempty"`
	Delay int64   `json:"delay"`

	//使用参数
	Arg int `json:"arg,omitempty"` //0:默认参数 1 2 3

	//使用表达式
	Expression string `json:"expression,omitempty"`
}

type Invoke struct {
	Command string    `json:"command"`
	Argv    []float64 `json:"argv"`

	//目标设备（只在Project中使用）
	Select Select `json:"select"`
}

type Select struct {
	Names []string `json:"device,omitempty"` //name
	Ids   []int    `json:"ids,omitempty"`
	Tags  []string `json:"tags,omitempty"`
}

func (s *Select) has(d *ProjectDevice) bool {
	for _, name := range s.Names {
		if name == d.Name {
			return true
		}
	}
	for _, name := range s.Ids {
		if name == d.Id {
			return true
		}
	}
	if hasTag(s.Tags, d.device.Tags) {
		return true
	}
	return false
}
