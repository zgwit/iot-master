package service

import "github.com/zgwit/iot-master/common"

type Link interface {
	common.EventEmitterInterface

	ID() int
	Write(data []byte) error
	//Read(data []byte) (int, error)
	Close() error
}
