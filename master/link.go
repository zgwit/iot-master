package master

import (
	"github.com/zgwit/iot-master/model"
	"github.com/zgwit/iot-master/protocol"
)

type Link struct {
	model.Link
	adapter protocol.Adapter
}

