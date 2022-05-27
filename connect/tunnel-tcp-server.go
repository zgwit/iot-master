package connect

import (
	"errors"
	"github.com/zgwit/iot-master/db"
	"github.com/zgwit/iot-master/events"
	"github.com/zgwit/iot-master/model"
	"net"
	"time"
)

//TunnelTcpServer TCP服务器
type TunnelTcpServer struct {
	events.EventEmitter
	tunnelBase

	listener *net.TCPListener

	running bool
}

func newTunnelTcpServer(tunnel *model.Tunnel) *TunnelTcpServer {
	svr := &TunnelTcpServer{
		tunnelBase: tunnelBase{tunnel: tunnel},
	}
	return svr
}

//Open 打开
func (server *TunnelTcpServer) Open() error {
	if server.running {
		return errors.New("server is opened")
	}
	server.Emit("open")

	addr, err := net.ResolveTCPAddr("tcp", resolvePort(server.tunnel.Addr))
	if err != nil {
		return err
	}
	server.listener, err = net.ListenTCP("tcp", addr)
	if err != nil {
		return err
	}

	server.running = true
	go func() {
		for {
			conn, err := server.listener.AcceptTCP()
			if err != nil {
				//TODO 需要正确处理接收错误
				break
			}

			server.link = conn
			//上线
			server.tunnel.Last = time.Now()
			server.tunnel.Remote = conn.RemoteAddr().String()
			_, _ = db.Engine.ID(server.tunnel.Id).Cols("last", "remote").Update(server.tunnel)

			server.receive()
		}

		server.running = false
	}()

	return nil
}

func (server *TunnelTcpServer) receive() {
	server.running = true
	buf := make([]byte, 1024)
	for {
		n, err := server.link.Read(buf)
		if err != nil {
			server.onClose()
			break
		}
		if n == 0 {
			continue
		}
		//透传转发
		if server.pipe != nil {
			_, err = server.pipe.Write(buf[:n])
			if err != nil {
				server.pipe = nil
			} else {
				continue
			}
		}
		server.Emit("data", buf[:n])
	}
	server.running = false
}

//Close 关闭
func (server *TunnelTcpServer) Close() (err error) {
	server.Emit("close")
	_ = server.link.Close()
	return server.listener.Close()
}
