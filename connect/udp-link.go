package connect

import (
	"github.com/zgwit/iot-master/events"
	"net"
)

type UdpLink struct {
	events.EventEmitter

	Id     int
	conn   *net.UDPConn
	addr   *net.UDPAddr
}

func newUdpLink(conn *net.UDPConn, addr *net.UDPAddr) *UdpLink {
	return &UdpLink{
		conn:   conn,
		addr:   addr,
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
	l.Emit("close")
}

func (l *UdpLink) onData(data []byte) {
	l.Emit("data", data)
}