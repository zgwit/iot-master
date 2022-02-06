package database

import (
	"github.com/asdine/storm/v3"
	"github.com/zgwit/iot-master/config"
)

var Storm *storm.DB


func Open() (err error) {
	Storm, err = storm.Open(config.Config.Database.Path)
	return
}

func Close() error {
	return Storm.Close()
}
