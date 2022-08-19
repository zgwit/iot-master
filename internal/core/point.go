package core

import (
	"iot-master/model"
	"iot-master/protocols/protocol"
)

type Point struct {
	model.Point
	Addr protocol.Addr
}
