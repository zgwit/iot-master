package connect

import (
	"github.com/zgwit/iot-master/events"
)

//Link 链接
type Link interface {
	events.EventInterface

	Id() int

	Write(data []byte) error

	Close() error

	Running() bool
}
