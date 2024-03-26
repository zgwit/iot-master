package protocol

import (
	"github.com/zgwit/iot-master/v4/connect"
	"github.com/zgwit/iot-master/v4/device"
	"github.com/zgwit/iot-master/v4/types"
)

type Factory func(tunnel string, conn connect.Conn, opts types.Options) (device.Adapter, error)

type Protocol struct {
	Name    string           `json:"name"`
	Label   string           `json:"label"`
	Factory Factory          `json:"-"`
	Options []types.FormItem `json:"-"`
	Pollers []types.FormItem `json:"-"`
	Mappers []types.FormItem `json:"-"`
}
