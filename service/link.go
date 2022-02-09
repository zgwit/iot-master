package service

import (
	"github.com/asaskevich/EventBus"
	"net"
)

type Link interface {
	Write(data []byte) error
	Read(data []byte) (int, error)
	Close() error
	OnClose(fn func())
}

type NetLink struct {
	Id     int
	conn   net.Conn
	events EventBus.Bus
}

func NewNetLink(conn net.Conn) *NetLink {
	return &NetLink{
		conn:   conn,
		events: EventBus.New(),
	}
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

func (l *NetLink) Close() error {
	l.onClose()
	return l.conn.Close()
}

func (l *NetLink) onClose() {
	l.events.Publish("close")
}

func (l *NetLink) OnClose(fn func()) {
	_ = l.events.Subscribe("close", fn)
}
