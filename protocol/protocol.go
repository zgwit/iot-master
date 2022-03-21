package protocol

import "github.com/zgwit/iot-master/connect"

//type Protocol interface {
//	Write(slave, code, address int, data []byte) error
//	Read(slave, code, address, length int) ([]byte, error)
//
//	// Immediate 立即读，高优先级
//	Immediate(slave, code, address, length int) ([]byte, error)
//}

//Protocol 协议接口
type Protocol interface {
	Write(addr Addr, data []byte) error
	Read(addr Addr, size uint16) ([]byte, error)

	// Immediate 立即读，高优先级
	Immediate(addr Addr, size uint16) ([]byte, error)
}

type Options map[string]interface{}

type Factory func(link connect.Link, opts Options) Protocol

type Address func(addr string) (Addr, error)

type Describer struct {
	Name    string `json:"name"`
	Label   string `json:"label"`
	Version string `json:"version"`
	Factory Factory
	Address Address
}
