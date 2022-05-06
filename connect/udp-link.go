package connect

import (
	"net"
)

//UdpLink UDP链接
type UdpLink struct {
	baseLink

	conn *net.UDPConn
	addr *net.UDPAddr
}

func newUdpLink(conn *net.UDPConn, addr *net.UDPAddr) *UdpLink {
	return &UdpLink{
		baseLink: baseLink{link: conn},
		conn:     conn,
		addr:     addr,
	}
}

//Write 写
func (l *UdpLink) Write(data []byte) error {
	_, err := l.conn.WriteToUDP(data, l.addr)
	if err != nil {
		l.onClose()
	}
	return err
}

func (l *UdpLink) onData(data []byte) {
	l.running = true
	l.Emit("data", data)
}
