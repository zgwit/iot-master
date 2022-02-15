package tsdb

import "time"

//Option 参数
type Option struct {
	DataPath           string        `yaml:"data_path"`
	TimestampPrecision string        `yaml:"timestamp_precision"` //ns us ms s
	RetentionDuration  time.Duration `yaml:"retention_duration"`  //s
	PartitionDuration  time.Duration `yaml:"partition_duration"`  //s
	WriteTimeout       time.Duration `yaml:"write_timeout"`       //s
	BufferedSize       int           `yaml:"buffered_size"`
	Log                bool          `yaml:"log"`
}
