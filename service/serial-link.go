package service

import (
	events2 "github.com/zgwit/iot-master/internal/events"
	"io"
)

type SerialLink struct {
	events2.EventEmitter

	Id     int
	port   io.ReadWriteCloser
}

func newSerialLink(port   io.ReadWriteCloser) *SerialLink {
	return &SerialLink{
		port:   port,
	}
}

func (l *SerialLink) ID() int {
	return l.Id
}

func (l *SerialLink) Write(data []byte) error {
	_, err := l.port.Write(data)
	if err != nil {
		l.onClose()
	}
	return err
}

func (l *SerialLink) receive() {
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

func (l *SerialLink) Close() error {
	l.onClose()
	return l.port.Close()
}

func (l *SerialLink) onClose() {
	l.Emit("close")
}
