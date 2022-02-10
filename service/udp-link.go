package service

import (
	"github.com/asaskevich/EventBus"
	"net"
)

type UdpLink struct {
	Id     int
	conn   *net.UDPConn
	addr   *net.UDPAddr
	events EventBus.Bus
}

func newUdpLink(conn *net.UDPConn, addr *net.UDPAddr) *UdpLink {
	return &UdpLink{
		conn:   conn,
		addr:   addr,
		events: EventBus.New(),
	}
}

func (l *UdpLink) ID() int {
	return l.Id
}

func (l *UdpLink) Write(data []byte) error {
	_, err := l.conn.WriteToUDP(data, l.addr)
	if err != nil {
		l.onClose()
	}
	return err
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

func (l *UdpLink) onData(data []byte) {
	l.events.Publish("data", data)
}

func (l *UdpLink) OnData(fn func(data []byte)) {
	_ = l.events.Subscribe("data", fn)
}
