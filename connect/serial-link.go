package connect

import (
	"github.com/zgwit/iot-master/events"
	"io"
)

//SerialLink 串口连接
type SerialLink struct {
	events.EventEmitter

	id      int
	port    io.ReadWriteCloser
	running bool
	first bool
}

func newSerialLink(port io.ReadWriteCloser) *SerialLink {
	return &SerialLink{
		port: port,
	}
}

//Id Id
func (l *SerialLink) Id() int {
	return l.id
}

//Write 写
func (l *SerialLink) Write(data []byte) error {
	_, err := l.port.Write(data)
	if err != nil {
		l.onClose()
	}
	return err
}

func (l *SerialLink) receive() {
	l.running = true
	buf := make([]byte, 1024)
	for {
		n, err := l.port.Read(buf)
		if err != nil {
			l.onClose()
			break
		}
		if n == 0 {
			continue
		}
		l.Emit("data", buf[:n])
	}
	l.running = false
}

//Close 关闭
func (l *SerialLink) Close() error {
	l.onClose()
	return l.port.Close()
}

func (l *SerialLink) onClose() {
	l.Emit("close")
}

func (l *SerialLink) Running() bool {
	return l.running
}

func (l *SerialLink) First() bool {
	return l.first
}