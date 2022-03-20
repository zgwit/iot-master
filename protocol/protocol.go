package protocol

import "github.com/zgwit/iot-master/connect"

//type Protocol interface {
//	Write(slave, code, address int, data []byte) error
//	Read(slave, code, address, length int) ([]byte, error)
//
//	// ImmediateRead 立即读，高优先级
//	ImmediateRead(slave, code, address, length int) ([]byte, error)
//}

//Protocol 协议接口
type Protocol interface {
	Write(addr Address, data []byte) error
	Read(addr Address, size uint16) ([]byte, error)

	// ImmediateRead 立即读，高优先级
	ImmediateRead(addr Address, size uint16) ([]byte, error)
}

type Options map[string]interface{}

type Factory func(link connect.Link, opts Options) Protocol

type AddressParser func(addr string) (Address, error)

type Item struct {
	Name    string `json:"name"`
	Label   string `json:"label"`
	Version string `json:"version"`
	Factory Factory
	Address AddressParser
}
