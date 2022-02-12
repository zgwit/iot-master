package service

import (
	"github.com/zgwit/iot-master/events"
	"io"
)

type SerialConn struct {
	events.EventEmitter

	Id   int
	port io.ReadWriteCloser
}

func newSerialLink(port io.ReadWriteCloser) *SerialConn {
	return &SerialConn{
		port: port,
	}
}

func (l *SerialConn) ID() int {
	return l.Id
}

func (l *SerialConn) Write(data []byte) error {
	_, err := l.port.Write(data)
	if err != nil {
		l.onClose()
	}
	return err
}

func (l *SerialConn) receive() {
	buf := make([]byte, 1024)
	for {
		n, err := l.port.Read(buf)
		if err != nil {
			l.onClose()
			break
		}
		l.Emit("data", buf[n:])
	}
}

func (l *SerialConn) Close() error {
	l.onClose()
	return l.port.Close()
}

func (l *SerialConn) onClose() {
	l.Emit("close")
}
