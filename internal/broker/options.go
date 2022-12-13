package broker

// Options 参数
type Options struct {
	Url      string `yaml:"url" json:"url"`
	ClientId string `yaml:"client_id" json:"client_id,omitempty"`
	Username string `yaml:"username" json:"username,omitempty"`
	Password string `yaml:"password" json:"password,omitempty"`
}
