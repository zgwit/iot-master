package config

import "time"

//History 参数
type History struct {
	Enable             bool          `yaml:"enable"`
	DataPath           string        `yaml:"data_path"`
	TimestampPrecision string        `yaml:"timestamp_precision"` //ns us ms s
	RetentionDuration  time.Duration `yaml:"retention_duration"`  //s
	PartitionDuration  time.Duration `yaml:"partition_duration"`  //s
	WriteTimeout       time.Duration `yaml:"write_timeout"`       //s
}

var HistoryDefault = History{
	Enable:             true,
	DataPath:           "history",
	TimestampPrecision: "ms",
}

