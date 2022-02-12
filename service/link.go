package service

import (
	events2 "github.com/zgwit/iot-master/internal/events"
)

type Link interface {
	events2.EventEmitterInterface

	ID() int
	Write(data []byte) error
	//Read(data []byte) (int, error)
	Close() error
}
