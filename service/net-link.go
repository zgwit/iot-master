package service

import (
	"github.com/asaskevich/EventBus"
	"net"
)

type NetLink struct {
	Id     int
	conn   net.Conn
	events EventBus.Bus
}

func newNetLink(conn net.Conn) *NetLink {
	return &NetLink{
		conn:   conn,
		events: EventBus.New(),
	}
}

func (l *NetLink) ID() int {
	return l.Id
}

func (l *NetLink) Write(data []byte) error {
	_, err := l.conn.Write(data)
	if err != nil {
		l.onClose()
	}
	return err
}

func (l *NetLink) Read(data []byte) (int, error) {
	n, err := l.conn.Read(data)
	if err != nil {
		l.onClose()
	}
	return n, err
}

func (l *NetLink) receive() {
	buf := make([]byte, 1024)
	for {
		n, err := l.conn.Read(buf)
		if err != nil {
			l.onClose()
			break
		}
		l.events.Publish("data", buf[n:])
	}
}

func (l *NetLink) Close() error {
	l.onClose()
	return l.conn.Close()
}

func (l *NetLink) onClose() {
	l.events.Publish("close")
}

func (l *NetLink) OnClose(fn func()) {
	_ = l.events.SubscribeOnce("close", fn)
}

func (l *NetLink) OnData(fn func(data []byte)) {
	_ = l.events.Subscribe("data", fn)
}
