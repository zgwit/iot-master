package tsdb

import (
	"github.com/nakabonne/tstorage"
	"strconv"
	"time"
)

var Storage tstorage.Storage

func Open(cfg *Option) error {

	opts := make([]tstorage.Option, 0)

	if cfg.DataPath != "" {
		opts = append(opts, tstorage.WithDataPath(cfg.DataPath))
	}
	if cfg.TimestampPrecision != "" {
		opts = append(opts, tstorage.WithTimestampPrecision(tstorage.TimestampPrecision(cfg.TimestampPrecision)))
	}
	if cfg.RetentionDuration > 0 {
		opts = append(opts, tstorage.WithRetention(cfg.RetentionDuration*time.Second))
	}
	if cfg.PartitionDuration > 0 {
		opts = append(opts, tstorage.WithPartitionDuration(cfg.PartitionDuration*time.Second))
	}
	if cfg.WriteTimeout > 0 {
		opts = append(opts, tstorage.WithPartitionDuration(cfg.WriteTimeout*time.Second))
	}
	if cfg.BufferedSize > 0 {
		opts = append(opts, tstorage.WithWALBufferedSize(cfg.BufferedSize))
	}
	if cfg.Log {
		//opts = append(opts, tstorage.WithLogger(nil))
	}

	var err error
	Storage, err = tstorage.NewStorage(opts...)

	return err
}

func Save(metric string, id int, point float64) error {
	rows := []tstorage.Row{{
		Metric:    metric,
		Labels:    []tstorage.Label{{Name: "key", Value: strconv.Itoa(id)}},
		DataPoint: tstorage.DataPoint{Value: point, Timestamp: time.Now().Unix()},
	}}
	return Storage.InsertRows(rows)
}

func Load(metric string, id int, start, end int64) ([]*tstorage.DataPoint, error) {
	//TODO 简单查询，并作结果整合
	return Storage.Select(metric, []tstorage.Label{{Name: "key", Value: strconv.Itoa(id)}}, start, end)
	//TODO 处理数据
}

func Close() error {
	err := Storage.Close()
	Storage = nil
	return err
}
