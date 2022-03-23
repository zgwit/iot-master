package master

import (
	"github.com/zgwit/iot-master/model"
	"github.com/zgwit/iot-master/protocol"
	"sync"
)

type Link struct {
	model.Link
	adapter protocol.Adapter
}

var links sync.Map
