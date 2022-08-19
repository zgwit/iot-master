package link

import (
	"io"
	"iot-master/model"
	"iot-master/pkg/events"
	"time"
)

//Tunnel 通道
type Tunnel interface {
	events.EventInterface

	Model() *model.Tunnel

	Write(data []byte) error

	Open() error

	Close() error

	Running() bool

	Online() bool

	First() bool

	//Pipe 透传
	Pipe(pipe io.ReadWriteCloser)

	//Ask 发送指令，接收数据
	Ask(cmd []byte, timeout time.Duration) ([]byte, error)
}
