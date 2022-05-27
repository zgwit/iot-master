package connect

import (
	"errors"
	"github.com/zgwit/iot-master/model"
	"net"
)

//ServerTcpTunnel 网络连接
type ServerTcpTunnel struct {
	tunnelBase
}

func newServerTcpTunnel(tunnel *model.Tunnel, conn net.Conn) *ServerTcpTunnel {
	return &ServerTcpTunnel{tunnelBase: tunnelBase{
		tunnel: tunnel,
		link:   conn,
	}}
}

func (l *ServerTcpTunnel) Open() error {
	return errors.New("ServerTcpTunnel cannot open")
}

//Write 写
func (l *ServerTcpTunnel) Write(data []byte) error {
	if l.pipe != nil {
		return nil //透传模式下，直接抛弃
	}
	_, err := l.link.Write(data)
	if err != nil {
		l.onClose()
	}
	return err
}

//Read 读
func (l *ServerTcpTunnel) Read(data []byte) (int, error) {
	n, err := l.link.Read(data)
	if err != nil {
		l.onClose()
	}
	return n, err
}

func (l *ServerTcpTunnel) receive() {
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
