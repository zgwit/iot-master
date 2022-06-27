package history

import (
	"fmt"
	"iot-master/config"
	"iot-master/history/influxdb2"
	"iot-master/history/storage"
	"iot-master/history/tstorage"
)

var Storage storage.Storage

func Open(options *config.History) error {
	switch options.Type {
	case "embed":
		Storage = &tstorage.TStorage{}
	case "influxdb2":
		Storage = &influxdb2.Influxdb{}
	default:
		return fmt.Errorf("未支持的历史数据库类型：%s", options.Type)
	}
	return Storage.Open(options.Options)
}

func Close() error {
	return Storage.Close()
}
