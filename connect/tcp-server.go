package connect

import (
	"github.com/asdine/storm/v3"
	"github.com/asdine/storm/v3/q"
	"github.com/zgwit/iot-master/database"
	"github.com/zgwit/iot-master/events"
	"net"
	"time"
)

type TcpServer struct {
	events.EventEmitter

	service *Tunnel

	children map[int]*NetConn

	listener *net.TCPListener
}

func newTcpServer(service *Tunnel) *TcpServer {
	svr := &TcpServer{
		service: service,
	}
	if service.Register != nil {
		svr.children = make(map[int]*NetConn)
	}
	return svr
}

func (server *TcpServer) Open() error {
	server.Emit("open")

	addr, err := net.ResolveTCPAddr("tcp", server.service.Addr)
	if err != nil {
		return err
	}
	server.listener, err = net.ListenTCP("tcp", addr)
	if err != nil {
		return err
	}
	go func() {
		for {
			conn, err := server.listener.AcceptTCP()
			if err != nil {
				//TODO 需要正确处理接收错误
				break
			}

			lnk := Link{
				ServiceId: server.service.Id,
				Created:   time.Now(),
			}

			if server.service.Register == nil {
				//先结束历史链接
				for _, link := range server.children {
					_ = link.Close()
				}
				err = database.Link.One("ServiceId", server.service.Id, &lnk)
			} else {
				buf := make([]byte, 128)
				n, err := conn.Read(buf)
				if err != nil {
					_ = conn.Close()
					continue
				}
				data := buf[n:]
				if !server.service.Register.Check(data) {
					_ = conn.Close()
					continue
				}
				sn := string(data)
				lnk.SN = sn
				err = database.Link.Select(
					q.And(
						q.Eq("ServiceId", server.service.Id),
						q.Eq("SN", sn),
					),
				).First(&lnk)
			}

			if err == storm.ErrNotFound {
				//保存一条新记录
				_ = database.Link.Save(&lnk)
			} else if err != nil {
				//return err
				continue
			}

			link := newNetConn(conn)
			go link.receive()

			link.Id = lnk.Id
			server.children[lnk.Id] = link

			//启动对应的设备 发消息
			server.Emit("link", link)

			link.Once("close", func() {
				delete(server.children, link.Id)
			})
		}
	}()

	return nil
}

func (server *TcpServer) Close() (err error) {
	server.Emit("close")

	//close links
	if server.children != nil {
		for _, l := range server.children {
			_ = l.Close()
		}
	}
	return server.listener.Close()
}

func (server *TcpServer) GetLink(id int) (Conn, error) {
	return server.children[id], nil
}
