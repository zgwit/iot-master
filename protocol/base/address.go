package base

type Code byte

type Address struct {
	Code   Code   `json:"code"`
	Block  uint16 `json:"block,omitempty"`
	Offset uint16 `json:"offset"`
}

type Protocol interface {
	Write(slave, addr Address, data []byte) error
	Read(slave, addr Address, size int) ([]byte, error)

	// ImmediateRead 立即读，高优先级
	ImmediateRead(slave, addr Address, size int) ([]byte, error)
}
