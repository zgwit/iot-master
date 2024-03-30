package protocol

import (
	"github.com/zgwit/iot-master/v4/connect"
	"github.com/zgwit/iot-master/v4/types"
)

type Factory func(conn connect.Tunnel, opts types.Options) (Adapter, error)

type Protocol struct {
	Name    string  `json:"name"`
	Label   string  `json:"label"`
	Factory Factory `json:"-"`

	//通道参数
	Options []types.FormItem `json:"-"`

	//设备参数
	Stations []types.FormItem `json:"-"`

	//轮询器
	Pollers []types.FormItem `json:"-"`

	//映射表
	Mappers []types.FormItem `json:"-"`
}
