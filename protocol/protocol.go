package protocol

import "github.com/zgwit/iot-master/connect"

//Adapter 协议接口
type Adapter interface {
	//Address 解析地址
	Address(addr string) (Addr, error)

	//Write 写数据
	Write(station int, addr Addr, data []byte) error

	//Read 读数据
	Read(station int, addr Addr, size int) ([]byte, error)

	//Immediate 立即读，高优先级
	Immediate(station int, addr Addr, size int) ([]byte, error)
}

type Options map[string]interface{}

func (opts Options) GetInt(name string, value int) int {
	v, ok := opts[name]
	if ok {
		return v.(int)
	}
	return value
}

type Factory func(link connect.Link, opts Options) Adapter

type Code struct {
	Name  string `json:"name"`
	Label string `json:"label"`
}

type Protocol struct {
	Name    string  `json:"name"`
	Label   string  `json:"label"`
	Version string  `json:"version"`
	Codes   []Code  `json:"codes"`
	Factory Factory `json:"-"`
}
