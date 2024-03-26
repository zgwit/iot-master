package types

type FormItem struct {
	Key         string             `json:"key"`
	Label       string             `json:"label"`
	Type        string             `json:"type,omitempty"` //type object array
	Default     any                `json:"default,omitempty"`
	Placeholder string             `json:"placeholder,omitempty"`
	Tips        string             `json:"tips,omitempty"`
	Pattern     string             `json:"pattern,omitempty"`
	Options     []FormSelectOption `json:"options,omitempty"`
	Required    bool               `json:"required,omitempty"`
	Min         float64            `json:"min,omitempty"`
	Max         float64            `json:"max,omitempty"`
	Step        float64            `json:"step,omitempty"`

	Children []FormItem `json:"children,omitempty"` //子级？
	Array    bool       `json:"array,omitempty"`
}

type FormSelectOption struct {
	Value    any    `json:"value"`
	Label    string `json:"label"`
	Disabled bool   `json:"disabled,omitempty"`
}
