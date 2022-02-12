package service

import (
	"github.com/zgwit/iot-master/events"
)

type Link interface {
	events.EventInterface

	ID() int
	Write(data []byte) error
	//Read(data []byte) (int, error)
	Close() error
}
