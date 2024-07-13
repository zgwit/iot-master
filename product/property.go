package product

// Property 属性
type Property struct {
	Name        string        `json:"name,omitempty"`        //变量名称
	Label       string        `json:"label,omitempty"`       //显示名称
	Unit        string        `json:"unit,omitempty"`        //单位
	Type        string        `json:"type,omitempty"`        //bool string number array object
	Default     any           `json:"default,omitempty"`     //默认值
	Writable    bool          `json:"writable,omitempty"`    //是否可写
	Remember    bool          `json:"remember,omitempty"`    //是否保存历史
	Aggregators []*Aggregator `json:"aggregators,omitempty"` //聚合计算

	//Children *Property
}
