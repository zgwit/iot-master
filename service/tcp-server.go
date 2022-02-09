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
}

func NewTcpServer(service *model.Service) *TcpServer {
	svc := &TcpServer{service: service}
	if !service.Single {
		svc.children = make(map[int]*NetLink)
	}
	return svc
}

func (c *TcpServer) Open() error {
	l, err := net.Listen("tcp", c.service.Addr)
	if err != nil {
		return err
	}
	go func() {
		for {
			conn, err := l.Accept()
			if err != nil {
				//TODO 需要正确处理接收错误
				break
			}

			lnk := model.Link{
				ServiceId: c.service.Id,
				Created:   time.Now(),
			}

			if c.service.Single {
				//TODO 等待链接结束，再重新接收

				err = database.Link.One("ServiceId", c.service.Id, &lnk)
			} else {
				buf := make([]byte, 128)
				n, err := conn.Read(buf)
				if err != nil {
					continue
				}
				sn := string(buf[n:])
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

			c.link = NewNetLink(conn)
			c.link.Id = lnk.Id
			if !c.service.Single {
				c.children[lnk.Id] = c.link
			}
			//TODO 启动对应的设备

		}
	}()

	return nil
}

func (c *TcpServer) HasAcceptor() bool {
	return false
}

func (c *TcpServer) Close() error {
	if c.link != nil {
		return c.link.Close()
	}
	return nil //TODO return error
}

func (c *TcpServer) GetLink(id int) (Link, error) {
	if c.service.Single {
		return c.link, nil
	} else {
		return c.children[id], nil
	}
}
