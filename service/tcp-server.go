package service

import (
	"github.com/asdine/storm/v3"
	"github.com/asdine/storm/v3/q"
	"github.com/zgwit/iot-master/database"
	"github.com/zgwit/iot-master/model"
	"net"
	"time"
)

type TcpServer struct {
	service  *model.Service
	link     *NetLink
	children map[int]*NetLink

	listener net.Listener
}

func NewTcpServer(service *model.Service) *TcpServer {
	svc := &TcpServer{service: service}
	if service.Register != nil {
		svc.children = make(map[int]*NetLink)
	}
	return svc
}

func (c *TcpServer) Open() error {
	var err error
	c.listener, err = net.Listen("tcp", c.service.Addr)
	if err != nil {
		return err
	}
	go func() {
		for {
			conn, err := c.listener.Accept()
			if err != nil {
				//TODO 需要正确处理接收错误
				break
			}

			lnk := model.Link{
				ServiceId: c.service.Id,
				Created:   time.Now(),
			}

			if c.service.Register == nil {
				//TODO 等待链接结束，再重新接收
				if c.link != nil {
					_ = c.link.Close()
				}
				err = database.Link.One("ServiceId", c.service.Id, &lnk)
			} else {
				buf := make([]byte, 128)
				n, err := conn.Read(buf)
				if err != nil {
					_ = conn.Close()
					continue
				}
				data := buf[n:]
				if !c.service.Register.Check(data) {
					_ = conn.Close()
					continue
				}
				sn := string(data)
				lnk.SN = sn
				err = database.Link.Select(
					q.And(
						q.Eq("ServiceId", c.service.Id),
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

			link := NewNetLink(conn)
			link.Id = lnk.Id
			if c.service.Register == nil {
				c.link = link
			} else {
				c.children[lnk.Id] = link
			}
			//TODO 启动对应的设备 发消息

			link.OnClose(func() {
				//TODO 记录

				if c.service.Register == nil {
					c.link = nil
				} else {
					delete(c.children, link.Id)
				}
			})
		}
	}()

	return nil
}


func (c *TcpServer) Close() (err error) {
	//TODO close links
	return c.listener.Close()
}

func (c *TcpServer) GetLink(id int) (Link, error) {
	if c.service.Register == nil {
		return c.link, nil
	} else {
		return c.children[id], nil
	}
}
