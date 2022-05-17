package connect

import (
	"errors"
	"github.com/zgwit/iot-master/events"
	"io"
	"sync"
	"time"
)

//Link 链接
type Link interface {
	events.EventInterface

	Id() int

	Write(data []byte) error

	Close() error

	Running() bool

	First() bool

	//Pipe 透传
	Pipe(pipe io.ReadWriteCloser)

	//Poll 发送指令，接收数据
	Poll(cmd []byte, timeout time.Duration) ([]byte, error)
}

type baseLink struct {
	events.EventEmitter

	lock sync.Mutex

	link  io.ReadWriteCloser

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

func (l *baseLink) wait(duration time.Duration) ([]byte, error) {
	resp := make(chan []byte, 1)
	l.Once("data", func(data []byte) {
		resp <- data
	})
	select {
	case <-time.After(duration):
		return nil, errors.New("超时")
	case buf := <-resp:
		return buf, nil
	}
}

func (l *baseLink) Poll(cmd []byte, timeout time.Duration) ([]byte, error) {
	//堵塞
	l.lock.Lock()
	defer l.lock.Unlock() //自动解锁

	_, err := l.link.Write(cmd)
	if err != nil {
		return nil, err
	}
	return l.wait(timeout)
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
			n, err := pipe.Read(buf)
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
