package protocol

import (
	"github.com/zgwit/iot-master/v4/connect"
	"github.com/zgwit/iot-master/v4/device"
	"github.com/zgwit/iot-master/v4/types"
)

type Factory func(tunnel string, conn connect.Conn, opts types.Options) (device.Adapter, error)

type Protocol struct {
	Name    string
	Label   string
	Factory Factory
	Options []types.FormItem
	Mappers []types.FormItem
}
