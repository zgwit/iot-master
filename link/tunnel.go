package link

import (
	"github.com/zgwit/iot-master/model"
	"github.com/zgwit/iot-master/pkg/events"
	"io"
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
