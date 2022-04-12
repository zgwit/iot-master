package tsdb

import (
	"github.com/nakabonne/tstorage"
	"strconv"
	"time"
)

//Storage 引擎
var Storage tstorage.Storage

//Open 打开
func Open(opts *Options) error {
	if opts == nil {
		opts = DefaultOptions()
	}

	options := make([]tstorage.Option, 0)

	if opts.DataPath != "" {
		options = append(options, tstorage.WithDataPath(opts.DataPath))
	}
	if opts.TimestampPrecision != "" {
		options = append(options, tstorage.WithTimestampPrecision(tstorage.TimestampPrecision(opts.TimestampPrecision)))
	}
	if opts.RetentionDuration > 0 {
		options = append(options, tstorage.WithRetention(opts.RetentionDuration*time.Second))
	}
	if opts.PartitionDuration > 0 {
		options = append(options, tstorage.WithPartitionDuration(opts.PartitionDuration*time.Second))
	}
	if opts.WriteTimeout > 0 {
		options = append(options, tstorage.WithPartitionDuration(opts.WriteTimeout*time.Second))
	}
	if opts.BufferedSize > 0 {
		options = append(options, tstorage.WithWALBufferedSize(opts.BufferedSize))
	}
	if opts.Log {
		//options = append(options, tstorage.WithLogger(nil))
	}

	var err error
	Storage, err = tstorage.NewStorage(options...)

	return err
}

//Save 保存数据
func Save(metric string, key string, point float64) error {
	rows := []tstorage.Row{{
		Metric:    metric,
		Labels:    []tstorage.Label{{Name: "key", Value: key}},
		DataPoint: tstorage.DataPoint{Value: point, Timestamp: time.Now().Unix()},
	}}
	return Storage.InsertRows(rows)
}

//Load 加载数据
func Load(metric string, id int, start, end int64) ([]*tstorage.DataPoint, error) {
	//TODO 简单查询，并作结果整合
	return Storage.Select(metric, []tstorage.Label{{Name: "key", Value: strconv.Itoa(id)}}, start, end)
	//TODO 处理数据
}

//Close 关闭
func Close() error {
	err := Storage.Close()
	Storage = nil
	return err
}
