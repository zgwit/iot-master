package core

import (
	"github.com/zgwit/iot-master/model"
	"github.com/zgwit/iot-master/protocols/protocol"
)

type Point struct {
	model.Point
	Addr protocol.Addr
}
