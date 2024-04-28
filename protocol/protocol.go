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
	OptionForm []types.FormItem `json:"-"`

	//设备参数
	StationForm []types.FormItem `json:"-"`

	//轮询器
	PollersForm []types.FormItem `json:"-"`

	//映射表
	MapperForm []types.FormItem `json:"-"`
}
