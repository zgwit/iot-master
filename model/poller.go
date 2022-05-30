package model

//Poller 采集器
type Poller struct {
	Disabled bool   `json:"disabled"`
	Type     string `json:"type"` //interval, clock, crontab

	Interval int    `json:"interval,omitempty"`
	Clock    int    `json:"clock,omitempty"`
	Crontab  string `json:"crontab,omitempty"`

	Code    string `json:"code"`
	Address string `json:"address"`
	Length  int    `json:"length"`

	//TODO Filters

	//等待结果
	//Parallel bool `json:"parallel,omitempty"`
}
