package tsdb

import (
	"github.com/nakabonne/tstorage"
	"github.com/zgwit/iot-master/config"
	"time"
)


var Storage tstorage.Storage

func Open() error {

	cfg := &config.Config.History

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

func Close() error {
	err := Storage.Close()
	Storage = nil
	return err
}
