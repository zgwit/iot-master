package connect

import (
	"github.com/asdine/storm/v3"
	"github.com/asdine/storm/v3/q"
	"github.com/zgwit/iot-master/database"
	"github.com/zgwit/iot-master/events"
	"github.com/zgwit/iot-master/model"
	"net"
	"time"
)

//TcpServer TCP服务器
type TcpServer struct {
	events.EventEmitter

	service *model.Tunnel

	children map[int]*NetLink

	listener *net.TCPListener
}

func newTcpServer(service *model.Tunnel) *TcpServer {
	svr := &TcpServer{
		service: service,
	}
	if service.Register != nil {
		svr.children = make(map[int]*NetLink)
	}
	return svr
}

//Open 打开
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

			lnk := model.Link{
				TunnelId: server.service.Id,
				Created:  time.Now(),
			}

			if server.service.Register == nil {
				//先结束历史链接
				for _, link := range server.children {
					_ = link.Close()
				}
				err = database.Master.One("TunnelId", server.service.Id, &lnk)
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
				err = database.Master.Select(
					q.And(
						q.Eq("TunnelId", server.service.Id),
						q.Eq("SN", sn),
					),
				).First(&lnk)
			}

			if err == storm.ErrNotFound {
				//保存一条新记录
				_ = database.Master.Save(&lnk)
			} else if err != nil {
				//return err
				continue
			}

			link := newNetLink(conn)
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

//Close 关闭
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

//GetLink 获取连接
func (server *TcpServer) GetLink(id int) (Link, error) {
	return server.children[id], nil
}
