package service

import (
	"github.com/asaskevich/EventBus"
	"github.com/asdine/storm/v3"
	"github.com/asdine/storm/v3/q"
	"github.com/zgwit/iot-master/database"
	"github.com/zgwit/iot-master/model"
	"net"
	"time"
)

type UdpServer struct {
	service  *model.Service
	link     *UdpLink
	children map[int]*UdpLink

	listener *net.UDPConn
	events  EventBus.Bus
}

func NewUdpServer(service *model.Service) *UdpServer {
	svr := &UdpServer{
		service: service,
		events: EventBus.New(),
	}
	if service.Register != nil {
		svr.children = make(map[int]*UdpLink)
	}
	return svr
}

func (server *UdpServer) Open() error {
	addr, err := net.ResolveUDPAddr("udp", server.service.Addr)
	if err != nil {
		return err
	}
	go func() {
		for {
			conn, err := net.ListenUDP("udp", addr)
			if err != nil {
				//TODO 需要正确处理接收错误
				break
			}

			lnk := model.Link{
				ServiceId: server.service.Id,
				Created:   time.Now(),
			}

			if server.service.Register == nil {
				//TODO 等待链接结束，再重新接收
				if server.link != nil {
					_ = server.link.Close()
				}
				err = database.Link.One("ServiceId", server.service.Id, &lnk)
			} else {
				buf := make([]byte, 128)
				//n, err := conn.Read(buf)
				n, r, err := conn.ReadFromUDP(buf)
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

			link := newUdpLink(conn, addr)
			link.Id = lnk.Id
			if server.service.Register == nil {
				server.link = link
			} else {
				server.children[lnk.Id] = link
			}
			//TODO 启动对应的设备 发消息

			server.events.Publish("link", link)

			link.OnClose(func() {
				//TODO 记录

				if server.service.Register == nil {
					server.link = nil
				} else {
					delete(server.children, link.Id)
				}
			})
		}
	}()

	return nil
}


func (server *UdpServer) Close() (err error) {
	//close links
	if server.link != nil {
		_ = server.link.Close()
	}
	if server.children != nil {
		for _, l := range server.children {
			_ = l.Close()
		}
	}
	return server.listener.Close()
}

func (server *UdpServer) GetLink(id int) (Link, error) {
	if server.service.Register == nil {
		return server.link, nil
	} else {
		return server.children[id], nil
	}
}

func (server *UdpServer) OnLink(fn func(link Link)) {
	_ = server.events.Subscribe("link", fn)
}
