package protocol

import (
	"fmt"
	"github.com/zgwit/iot-master/v4/connect"
	"github.com/zgwit/iot-master/v4/device"
	"github.com/zgwit/iot-master/v4/types"
)

var protocols = map[string]*Protocol{}

func Protocols() []*Protocol {
	var ps []*Protocol
	for _, p := range protocols {
		ps = append(ps, p)
	}
	return ps
}

func Register(proto *Protocol) {
	protocols[proto.Name] = proto
}

func Create(tunnel string, conn connect.Conn, name string, opts types.Options) (device.Adapter, error) {
	if p, ok := protocols[name]; ok {
		return p.Factory(tunnel, conn, opts)
	}
	return nil, fmt.Errorf("协议 %s 找不到", name)
}
