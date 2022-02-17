package connect

import (
	"github.com/zgwit/iot-master/events"
)

//Tunnel 通道
type Tunnel interface {
	events.EventInterface

	Open() error
	Close() error
	GetLink(id int) Link
}
