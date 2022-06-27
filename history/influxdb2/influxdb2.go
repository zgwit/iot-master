package influxdb2

import (
	"context"
	"fmt"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"iot-master/helper"
	"iot-master/history/storage"
	"time"
)

type Influxdb struct {
	writeAPI api.WriteAPIBlocking
	queryAPI api.QueryAPI

	client      influxdb2.Client
	measurement string
	bucket      string
}

func (i *Influxdb) Open(opts helper.Options) error {
	i.measurement = opts.GetDefaultString("measurement", "table")
	i.bucket = opts.GetDefaultString("bucket", "")

	i.client = influxdb2.NewClient(opts.GetDefaultString("url", ""), opts.GetDefaultString("token", ""))
	i.writeAPI = i.client.WriteAPIBlocking(opts.GetDefaultString("org", ""), opts.GetDefaultString("bucket", ""))
	i.queryAPI = i.client.QueryAPI(opts.GetDefaultString("org", ""))
	return nil
}

func (i *Influxdb) Close() error {
	i.client.Close()
	return nil
}

func (i *Influxdb) Write(id int64, values map[string]interface{}) error {
	metric := fmt.Sprintf("%d", id)
	//TODO tags 中添加 {node:config.Node}
	point := influxdb2.NewPoint(i.measurement, map[string]string{"id": metric}, values, time.Now())
	return i.writeAPI.WritePoint(context.Background(), point)
}

func (i *Influxdb) Query(id int64, field string, start, end, window string) ([]storage.Point, error) {
	metric := fmt.Sprintf("%d", id)

	flux := "from(bucket: \"" + i.bucket + "\")\n"
	flux += "|> range(start: " + start + ", stop: " + end + ")\n"
	flux += "|> filter(fn: (r) => r[\"id\"] == \"" + metric + "\")\n"
	flux += "|> filter(fn: (r) => r[\"_field\"] == \"" + field + "\")"
	flux += "|> aggregateWindow(every: " + window + ", fn: mean, createEmpty: false)\n"
	flux += "|> yield(name: \"mean\")"

	result, err := i.queryAPI.Query(context.Background(), flux)
	if err != nil {
		return nil, err
	}

	records := make([]storage.Point, 0)
	for result.Next() {
		//result.TableChanged()
		records = append(records, storage.Point{
			Value: result.Record().Value(),
			Time:  result.Record().Time(),
		})
	}
	return records, result.Err()
}
