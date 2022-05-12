package connect

import (
	"github.com/zgwit/iot-master/events"
	"io"
)

//Link 链接
type Link interface {
	events.EventInterface

	Id() int

	Write(data []byte) error

	Close() error

	Running() bool

	First() bool

	Pipe(pipe io.ReadWriteCloser)
}


func WriteAndRead(data []byte) ([]byte, error) {
	//link.read channel
	//link.Once("data", )

	return nil, nil
}

type baseLink struct {
	events.EventEmitter

	link io.ReadWriteCloser

	id      int
	running bool
	first   bool

	pipe io.ReadWriteCloser
}

func (l *baseLink) Id() int {
	return l.id
}

func (l *baseLink) Running() bool {
	return l.running
}

func (l *baseLink) First() bool {
	return l.first
}

//Close 关闭
func (l *baseLink) Close() error {
	l.onClose()
	return l.link.Close()
}

func (l *baseLink) onClose() {
	l.running = false
	if l.pipe != nil {
		_ = l.pipe.Close()
	}
	l.Emit("close")
}

func (l *baseLink) Pipe(pipe io.ReadWriteCloser) {
	//关闭之前的透传
	if l.pipe != nil {
		_ = l.pipe.Close()
	}

	l.pipe = pipe
	//传入空，则关闭
	if pipe == nil {
		return
	}

	go func() {
		buf := make([]byte, 1024)
		for {
			n ,err := pipe.Read(buf)
			if err != nil {
				//pipe关闭，则不再透传
				break
			}
			//将收到的数据转发出去
			n, err = l.link.Write(buf[:n])
			if err != nil {
				//发送失败，说明连接失效
				_ = pipe.Close()
				break
			}
		}
		l.pipe = nil
	}()
}

