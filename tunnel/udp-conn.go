package tunnel

import (
	"github.com/zgwit/iot-master/events"
	"net"
)

type UdpConn struct {
	events.EventEmitter

	Id     int
	conn   *net.UDPConn
	addr   *net.UDPAddr
}

func newUdpLink(conn *net.UDPConn, addr *net.UDPAddr) *UdpConn {
	return &UdpConn{
		conn:   conn,
		addr:   addr,
	}
}

func (l *UdpConn) ID() int {
	return l.Id
}

func (l *UdpConn) Write(data []byte) error {
	_, err := l.conn.WriteToUDP(data, l.addr)
	if err != nil {
		l.onClose()
	}
	return err
}

func (l *UdpConn) Close() error {
	l.onClose()
	return l.conn.Close()
}

func (l *UdpConn) onClose() {
	l.Emit("close")
}

func (l *UdpConn) onData(data []byte) {
	l.Emit("data", data)
}