package service

import (
	"github.com/asdine/storm/v3"
	"github.com/asdine/storm/v3/q"
	"github.com/zgwit/iot-master/database"
	"github.com/zgwit/iot-master/internal"
	events2 "github.com/zgwit/iot-master/internal/events"
	"net"
	"time"
)

type UdpServer struct {
	events2.EventEmitter

	service *internal.Service

	children map[int]*UdpLink
	links    map[string]*UdpLink

	listener *net.UDPConn
}

func NewUdpServer(service *internal.Service) *UdpServer {
	svr := &UdpServer{
		service: service,
	}
	if service.Register != nil {
		svr.children = make(map[int]*UdpLink)
		svr.links = make(map[string]*UdpLink)
	}
	return svr
}

func (server *UdpServer) Open() error {
	addr, err := net.ResolveUDPAddr("udp", server.service.Addr)
	if err != nil {
		return err
	}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		//TODO 需要正确处理接收错误
		return err
	}
	server.listener = conn //共用连接

	go func() {
		for {
			buf := make([]byte, 1024)
			n, addr, err := conn.ReadFromUDP(buf)
			if err != nil {
				_ = conn.Close()
				continue
			}
			data := buf[n:]

			//如果已经保存了链接 TODO 要有超时处理
			link, ok := server.links[addr.String()]
			if ok {
				link.onData(data)
				continue
			}

			lnk := internal.Link{
				ServiceId: server.service.Id,
				Created:   time.Now(),
			}

			if server.service.Register == nil {
				//先结束其他链接
				for _, link := range server.links {
					_ = link.Close()
				}
				err = database.Link.One("ServiceId", server.service.Id, &lnk)
			} else {
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

			link = newUdpLink(conn, addr)
			link.Id = lnk.Id
			server.children[lnk.Id] = link

			//启动对应的设备 发消息
			server.Emit("link", link)

			link.Once("close", func() {
				//TODO 记录离线
				delete(server.children, link.Id)
				delete(server.links, link.addr.String())
			})
		}
	}()

	return nil
}

func (server *UdpServer) Close() (err error) {
	//close links
	if server.children != nil {
		for _, l := range server.children {
			_ = l.Close()
		}
	}
	return server.listener.Close()
}

func (server *UdpServer) GetLink(id int) (Link, error) {
	return server.children[id], nil
}
