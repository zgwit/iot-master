package tunnel

import (
	"github.com/zgwit/iot-master/events"
	"net"
)

type NetConn struct {
	events.EventEmitter

	Id     int
	conn   net.Conn
}

func newNetConn(conn net.Conn) *NetConn {
	return &NetConn{
		conn:   conn,
	}
}

func (l *NetConn) ID() int {
	return l.Id
}

func (l *NetConn) Write(data []byte) error {
	_, err := l.conn.Write(data)
	if err != nil {
		l.onClose()
	}
	return err
}

func (l *NetConn) Read(data []byte) (int, error) {
	n, err := l.conn.Read(data)
	if err != nil {
		l.onClose()
	}
	return n, err
}

func (l *NetConn) receive() {
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

func (l *NetConn) Close() error {
	l.onClose()
	return l.conn.Close()
}

func (l *NetConn) onClose() {
	l.Emit("close")
}
