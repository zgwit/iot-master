package protocol

//Addr 地址
//type Addr struct {
//	Slave  uint8  `json:"slave,omitempty"`
//	Code   uint8  `json:"code"`
//	Block  uint16 `json:"block,omitempty"`
//	Offset uint16 `json:"offset"`
//}

type Addr interface {
	String() string
	Diff(base Addr) int
}

