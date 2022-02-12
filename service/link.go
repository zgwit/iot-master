package service

import (
	"github.com/zgwit/iot-master/events"
)

type Link interface {
	events.EventEmitterInterface

	ID() int
	Write(data []byte) error
	//Read(data []byte) (int, error)
	Close() error
}
