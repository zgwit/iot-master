package influx

import (
	"context"
	"fmt"
	"github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
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

func Write(measurement string, tags map[string]string, fields map[string]interface{}) error {
	point := influxdb2.NewPoint(measurement, tags, fields, time.Now())
	return writeAPI.WritePoint(context.Background(), point)
}

func Query() error {
	result, err := queryAPI.Query(context.Background(), "Flux")
	if err != nil {
		return err
	}
	for result.Next() {
		if result.TableChanged() {
			fmt.Printf("table: %s\n", result.TableMetadata().String())
		}
		fmt.Printf("value: %v\n", result.Record().Value())
	}
	if result.Err() != nil {
		fmt.Printf("query parsing error: %s\n", result.Err().Error())
	}
	return nil
}
