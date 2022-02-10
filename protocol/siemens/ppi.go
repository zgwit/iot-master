package siemens

import "github.com/zgwit/iot-master/service"

type PPI struct {
	link service.Link
}

//Read 读到数据
func (t *PPI)Read(address string, length int) ([]byte, error) {

}

//Write 写入数据
func (t *PPI)Write(address string, values []byte) error {

}