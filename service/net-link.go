package service

import (
	"github.com/zgwit/iot-master/events"
	"net"
)

type NetLink struct {
	events.EventEmitter

	Id     int
	conn   net.Conn
}

func newNetLink(conn net.Conn) *NetLink {
	return &NetLink{
		conn:   conn,
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
		l.Emit("data", buf[n:])
	}
}

func (l *NetLink) Close() error {
	l.onClose()
	return l.conn.Close()
}

func (l *NetLink) onClose() {
	l.Emit("close")
}
