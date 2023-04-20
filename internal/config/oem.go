package config

type OEM struct {
	Title     string `yaml:"title,omitempty" json:"title,omitempty"`
	Logo      string `yaml:"logo,omitempty" json:"logo,omitempty"`
	Company   string `yaml:"company,omitempty" json:"company,omitempty"`
	Copyright string `yaml:"copyright,omitempty" json:"copyright,omitempty"`
}
