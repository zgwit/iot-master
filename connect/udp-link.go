package connect

import (
	"io"
	"net"
	"time"
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
	if l.pipe != nil {
		return nil //透传模式下，直接抛弃
	}
	_, err := l.conn.WriteToUDP(data, l.addr)
	if err != nil {
		l.onClose()
	}
	return err
}

func (l *UdpLink) Poll(cmd []byte, timeout time.Duration) ([]byte, error)  {
	//堵塞
	l.lock.Lock()
	defer l.lock.Unlock() //自动解锁

	err := l.Write(cmd)
	if err != nil {
		return nil, err
	}
	return l.wait(timeout)
}

func (l *UdpLink) Pipe(pipe io.ReadWriteCloser) {
	//关闭之前的透传
	if l.pipe != nil {
		_ = l.pipe.Close()
	}
	l.pipe = pipe

	//传入空，则关闭
	if l.pipe == nil {
		return
	}

	go func() {
		buf := make([]byte, 1024)
		for {
			n ,err := pipe.Read(buf)
			if err != nil {
				//pipe关闭，则不再透传
				break
			}
			//将收到的数据转发出去
			//n, err = l.link.Write(buf[:n])
			_, err = l.conn.WriteToUDP(buf[:n], l.addr)
			if err != nil {
				//发送失败，说明连接失效
				_ = pipe.Close()
				break
			}
		}
		l.pipe = nil
	}()
}

func (l *UdpLink) onData(data []byte) {
	l.running = true

	//透传
	if l.pipe != nil {
		_, err := l.pipe.Write(data)
		if err == nil {
			return
		}
		l.pipe = nil
	}

	l.Emit("data", data)
}

