package service

import (
	"github.com/asaskevich/EventBus"
	"net"
)

type TcpLink struct {
	Id     int
	conn   *net.TCPConn
	events EventBus.Bus
}

func newTcpLink(conn *net.TCPConn) *TcpLink {
	return &TcpLink{
		conn:   conn,
		events: EventBus.New(),
	}
}

func (l *TcpLink) ID() int {
	return l.Id
}

func (l *TcpLink) Write(data []byte) error {
	_, err := l.conn.Write(data)
	if err != nil {
		l.onClose()
	}
	return err
}

func (l *TcpLink) Read(data []byte) (int, error) {
	n, err := l.conn.Read(data)
	if err != nil {
		l.onClose()
	}
	return n, err
}

func (l *TcpLink) Close() error {
	l.onClose()
	return l.conn.Close()
}

func (l *TcpLink) onClose() {
	l.events.Publish("close")
}

func (l *TcpLink) OnClose(fn func()) {
	_ = l.events.SubscribeOnce("close", fn)
}
