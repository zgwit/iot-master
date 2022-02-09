package service

import (
	"github.com/asaskevich/EventBus"
	"net"
)

type UdpLink struct {
	Id     int
	conn   *net.UDPConn
	events EventBus.Bus
}

func newUdpLink(conn *net.UDPConn) *UdpLink {
	return &UdpLink{
		conn:   conn,
		events: EventBus.New(),
	}
}

func (l *UdpLink) ID() int {
	return l.Id
}

func (l *UdpLink) Write(data []byte) error {
	_, err := l.conn.Write(data)
	if err != nil {
		l.onClose()
	}
	return err
}

func (l *UdpLink) Read(data []byte) (int, error) {
	//n, err := l.conn.Read(data)
	n, _, err := l.conn.ReadFromUDP(data)
	if err != nil {
		l.onClose()
	}
	return n, err
}

func (l *UdpLink) Close() error {
	l.onClose()
	return l.conn.Close()
}

func (l *UdpLink) onClose() {
	l.events.Publish("close")
}

func (l *UdpLink) OnClose(fn func()) {
	_ = l.events.SubscribeOnce("close", fn)
}
