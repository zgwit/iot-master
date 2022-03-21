package master

import (
	"github.com/zgwit/iot-master/model"
	"github.com/zgwit/iot-master/protocol"
)

type Point struct {
	model.Point
	Addr protocol.Addr
}
