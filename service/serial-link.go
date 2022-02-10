package service

import (
	"github.com/asaskevich/EventBus"
	"io"
)

type SerialLink struct {
	Id     int
	port   io.ReadWriteCloser
	events EventBus.Bus
}

func newSerialLink(port   io.ReadWriteCloser) *SerialLink {
	return &SerialLink{
		port:   port,
		events: EventBus.New(),
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
		l.events.Publish("data", buf[n:])
	}
}

func (l *SerialLink) Close() error {
	l.onClose()
	return l.port.Close()
}

func (l *SerialLink) onClose() {
	l.events.Publish("close")
}

func (l *SerialLink) OnClose(fn func()) {
	_ = l.events.SubscribeOnce("close", fn)
}

func (l *SerialLink) OnData(fn func(data []byte)) {
	_ = l.events.Subscribe("data", fn)
}
