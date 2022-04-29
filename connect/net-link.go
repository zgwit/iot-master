package connect

import (
	"github.com/zgwit/iot-master/events"
	"net"
)

//NetLink 网络连接
type NetLink struct {
	events.EventEmitter

	id   int
	conn net.Conn
}

func newNetLink(conn net.Conn) *NetLink {
	return &NetLink{
		conn: conn,
	}
}

//Id Id
func (l *NetLink) Id() int {
	return l.id
}

//Write 写
func (l *NetLink) Write(data []byte) error {
	_, err := l.conn.Write(data)
	if err != nil {
		l.onClose()
	}
	return err
}

//Read 读
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
		l.Emit("data", buf[:n])
	}
}

//Close 关闭
func (l *NetLink) Close() error {
	l.onClose()
	return l.conn.Close()
}

func (l *NetLink) onClose() {
	l.Emit("close")
}
