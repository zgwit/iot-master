package influx

import (
	"context"
	"github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"time"
)

var writeAPI api.WriteAPIBlocking
var queryAPI api.QueryAPI

var _opts *Options
var client influxdb2.Client

func Open(opts *Options) {
	_opts = opts
	client = influxdb2.NewClient(opts.URL, opts.Token)
	writeAPI = client.WriteAPIBlocking(opts.ORG, opts.Bucket)
	queryAPI = client.QueryAPI(opts.ORG)
}

func Close() {
	client.Close()
}

func Write(tags map[string]string, fields map[string]interface{}) error {
	point := influxdb2.NewPoint(_opts.Measurement, tags, fields, time.Now())
	return writeAPI.WritePoint(context.Background(), point)
}

type Point struct {
	Time  time.Time   `json:"time"`
	Value interface{} `json:"value"`
}

func Query(window, start, end string, tags map[string]string, field string) ([]Point, error) {
	flux := "from(bucket: \"" + _opts.Bucket + "\")\n"
	flux += "|> range(start: " + start + ", stop: " + end + ")\n"
	for k, v := range tags {
		flux += "|> filter(fn: (r) => r[\"" + k + "\"] == \"" + v + "\")\n"
	}
	flux += "|> filter(fn: (r) => r[\"_field\"] == \"" + field + "\")"
	flux += "|> aggregateWindow(every: " + window + ", fn: mean, createEmpty: false)\n"
	flux += "|> yield(name: \"mean\")"

	result, err := queryAPI.Query(context.Background(), flux)
	if err != nil {
		return nil, err
	}

	records := make([]Point, 0)
	for result.Next() {
		//result.TableChanged()
		records = append(records, Point{
			Time:  result.Record().Time(),
			Value: result.Record().Value(),
		})
	}
	return records, result.Err()
}
