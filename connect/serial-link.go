package connect

import (
	"io"
)

//SerialLink 串口连接
type SerialLink struct {
	baseLink
}

func newSerialLink(port io.ReadWriteCloser) *SerialLink {
	return &SerialLink{baseLink: baseLink{link: port}}
}

//Write 写
func (l *SerialLink) Write(data []byte) error {
	_, err := l.link.Write(data)
	if err != nil {
		l.onClose()
	}
	return err
}

func (l *SerialLink) receive() {
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
