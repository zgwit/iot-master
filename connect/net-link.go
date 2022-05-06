package connect

import (
	"net"
)

//NetLink 网络连接
type NetLink struct {
	baseLink
}

func newNetLink(conn net.Conn) *NetLink {
	return &NetLink{baseLink: baseLink{link: conn}}
}

//Write 写
func (l *NetLink) Write(data []byte) error {
	_, err := l.link.Write(data)
	if err != nil {
		l.onClose()
	}
	return err
}

//Read 读
func (l *NetLink) Read(data []byte) (int, error) {
	n, err := l.link.Read(data)
	if err != nil {
		l.onClose()
	}
	return n, err
}

func (l *NetLink) receive() {
	l.running = true
	buf := make([]byte, 1024)
	for {
		n, err := l.link.Read(buf)
		if err != nil {
			l.onClose()
			break
		}
		if n == 0 {
			continue
		}
		//透传转发
		if l.pipe != nil {
			_, err = l.pipe.Write(buf[:n])
			if err != nil {
				l.pipe = nil
			} else {
				continue
			}
		}
		l.Emit("data", buf[:n])
	}
	l.running = false
}
