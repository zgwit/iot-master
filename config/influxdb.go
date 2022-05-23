package config

//Influxdb 参数
type Influxdb struct {
	Enable      bool   `yaml:"enable"`
	Bucket      string `yaml:"bucket"`
	ORG         string `yaml:"org"`
	Token       string `yaml:"token"`
	URL         string `yaml:"url"`
	Measurement string `yaml:"measurement"`
}

var InfluxdbDefault = Influxdb{
	Enable: false,
}