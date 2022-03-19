package protocol

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

type Factory func(interface{}) Protocol

type Desc struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Factory Factory
}
