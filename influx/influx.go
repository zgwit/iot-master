package influx

import (
	"context"
	"github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"github.com/influxdata/influxdb-client-go/v2/api/query"
	"time"
)

var writeAPI api.WriteAPIBlocking
var queryAPI api.QueryAPI

func Open() {
	bucket := "example-bucket"
	org := "example-org"
	token := "example-token"
	// Store the URL of your InfluxDB instance
	url := "http://localhost:8086"
	client := influxdb2.NewClient(url, token)
	writeAPI = client.WriteAPIBlocking(org, bucket)
	queryAPI = client.QueryAPI(org)
}

func Close() {
	//client.Close
}

func Write(measurement string, tags map[string]string, fields map[string]interface{}) error  {
	point := influxdb2.NewPoint(measurement, tags, fields, time.Now())
	return writeAPI.WritePoint(context.Background(), point)
}

func Query(window, start, end string, tags map[string]string, field string) ([]*query.FluxRecord, error) {
	flux := "from(bucket: \"${bucket}\")\n"
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

	records := make([]*query.FluxRecord, 0)
	for result.Next() {
		//result.TableChanged()
		records = append(records, result.Record())
	}
	return records, result.Err()
}
