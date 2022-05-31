package tstorage

import (
	"errors"
	"fmt"
	"github.com/nakabonne/tstorage"
	"github.com/zgwit/iot-master/helper"
	"github.com/zgwit/iot-master/history/storage"
	"regexp"
	"strconv"
	"time"
)

type TStorage struct {
	storage tstorage.Storage
}

func (t *TStorage) Open(opts helper.Options) error {
	options := make([]tstorage.Option, 0)
	options = append(options, tstorage.WithDataPath(opts.GetDefaultString("path", "tstorage")))

	if precision, ok := opts.GetString("precision"); ok {
		tstorage.WithTimestampPrecision(tstorage.TimestampPrecision(precision))
	}
	if retention, ok := opts.GetInt("retention"); ok && retention > 0 {
		options = append(options, tstorage.WithRetention(time.Duration(retention)*time.Second))
	}
	if partition, ok := opts.GetInt("partition"); ok && partition > 0 {
		options = append(options, tstorage.WithPartitionDuration(time.Duration(partition)*time.Second))
	}
	//options = append(options, tstorage.WithPartitionDuration(opts.WriteTimeout*time.Second))

	var err error
	t.storage, err = tstorage.NewStorage(options...)
	return err
}

func (t *TStorage) Close() error {
	return t.storage.Close()
}

func (t *TStorage) Write(id int64, values map[string]interface{}) error {
	metric := fmt.Sprintf("%d", id)
	rows := make([]tstorage.Row, 0)
	for k, v := range values {
		rows = append(rows, tstorage.Row{
			Metric:    metric,
			Labels:    []tstorage.Label{{Name: "id", Value: k}},
			DataPoint: tstorage.DataPoint{Value: helper.ToFloat64(v), Timestamp: time.Now().Unix()},
		})
	}
	return t.storage.InsertRows(rows)
}

func (t *TStorage) Query(id int64, field string, start, end, window string) ([]storage.Point, error) {
	//相对时间转化为时间戳
	s, err := parseTime(start)
	if err != nil {
		return nil, err
	}
	s += time.Now().UnixMilli()

	e, err := parseTime(end)
	if err != nil {
		return nil, err
	}
	e += time.Now().UnixMilli()

	w, err := parseTime(window)
	if err != nil {
		return nil, err
	}

	metric := fmt.Sprintf("%d", id)
	points, err := t.storage.Select(metric, []tstorage.Label{{Name: "key", Value: field}}, s, e)
	if err != nil {
		//无数据
		if err == tstorage.ErrNoDataPoints {
			return make([]storage.Point, 0), nil
		}
		return nil, err
	}

	results := make([]storage.Point, 0)
	var total float64 = 0
	var count float64 = 0
	var timestamp int64

	for _, p := range points {
		//按窗口划分
		for p.Timestamp > s+w {
			start += window
			if count > 0 {
				results = append(results, storage.Point{
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
		results = append(results, storage.Point{
			Value: total / count,
			Time:  time.UnixMilli(timestamp),
		})
	}

	return results, nil
}

var timeReg *regexp.Regexp

func init() {
	timeReg = regexp.MustCompile(`^(-?\d+)(h|m|s)$`)
}

func parseTime(tm string) (int64, error) {
	ss := timeReg.FindStringSubmatch(tm)
	if ss == nil || len(ss) != 3 {
		return 0, errors.New("错误时间")
	}
	val, _ := strconv.ParseInt(ss[1], 10, 64)
	switch ss[2] {
	case "d":
		val *= 24 * 60 * 60 * 1000
	case "h":
		val *= 60 * 60 * 1000
	case "m":
		val *= 60 * 1000
	case "s":
		val *= 1000
	}
	return val, nil
}
