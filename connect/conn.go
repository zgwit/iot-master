package connect

import (
	"github.com/zgwit/iot-master/events"
)

type Conn interface {
	events.EventInterface

	ID() int

	Write(data []byte) error

	Close() error
}
