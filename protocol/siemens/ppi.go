package siemens

import "github.com/zgwit/iot-master/tunnel"

type PPI struct {
	link tunnel.Conn
}

//Read 读到数据
func (t *PPI)Read(address string, length int) ([]byte, error) {
	return nil,nil
}

//Write 写入数据
func (t *PPI)Write(address string, values []byte) error {
	return nil
}