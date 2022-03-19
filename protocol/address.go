package protocol

//Address 地址
type Address struct {
	Slave  byte   `json:"slave,omitempty"`
	Code   byte   `json:"code"`
	Block  uint16 `json:"block,omitempty"`
	Offset uint16 `json:"offset"`
}
