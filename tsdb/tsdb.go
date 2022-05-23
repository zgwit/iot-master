package tsdb

import (
	"github.com/nakabonne/tstorage"
	"github.com/zgwit/iot-master/config"
	"github.com/zgwit/iot-master/model"
	"strconv"
	"time"
)

//Storage 引擎
var Storage tstorage.Storage

func Opened() bool  {
	return Storage != nil
}

//Open 打开
func Open(opts *config.History) error {
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

func Write(metric string, values map[string]interface{}) error {
	rows := make([]tstorage.Row, 0)
	timestamp := time.Now().UnixMilli()
	for k, v := range values {
		value := v.(float64)
		rows = append(rows, tstorage.Row{
			Metric:    metric,
			Labels:    []tstorage.Label{{Name: "key", Value: k}},
			DataPoint: tstorage.DataPoint{Value: value, Timestamp: timestamp},
		})
	}
	return Storage.InsertRows(rows)
}

//Query 加载数据（毫秒精度）
func Query(metric string, key string, start, end, window int64) ([]model.DataPoint, error) {
	points, err := Storage.Select(metric, []tstorage.Label{{Name: "key", Value: key}}, start, end)

	if err != nil {
		//无数据
		if err == tstorage.ErrNoDataPoints {
			return make([]model.DataPoint, 0), nil
		}
		return nil, err
	}

	results := make([]model.DataPoint, 0)
	var total float64 = 0
	var count float64 = 0
	var timestamp int64

	for _, p := range points {
		//按窗口划分
		for p.Timestamp > start+window {
			start += window
			if count > 0 {
				results = append(results, model.DataPoint{
					Value: total / count,
					Time:  time.UnixMilli(timestamp),
				})
				total = 0
				count = 0
			}
		}

		total += p.Value
		count++
		timestamp = p.Timestamp
	}
	//最后一组
	if count > 0 {
		results = append(results, model.DataPoint{
			Value: total / count,
			Time:  time.UnixMilli(timestamp),
		})
	}

	return results, nil
}

//Close 关闭
func Close() error {
	err := Storage.Close()
	Storage = nil
	return err
}
