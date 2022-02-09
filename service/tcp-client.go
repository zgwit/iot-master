package service

import (
	"github.com/asdine/storm/v3"
	"github.com/zgwit/iot-master/database"
	"github.com/zgwit/iot-master/model"
	"net"
	"time"
)

type TcpClient struct {
	service *model.Service
	link    *NetLink
}

func NewTcpClient(service *model.Service) *TcpClient {
	return &TcpClient{service: service}
}

func (c *TcpClient) Open() error {
	conn, err := net.Dial("tcp", c.service.Addr)
	if err != nil {
		return err
	}
	c.link = NewNetLink(conn)
	//TODO store link
	lnk := model.Link{
		ServiceId: c.service.Id,
		Created:   time.Now(),
	}
	err = database.Link.Find("ServiceId", c.service.Id, &lnk)
	if err == storm.ErrNotFound {
		//保存一条新记录
		_ = database.Link.Save(&lnk)
	} else if err != nil {
		return err
	}
	c.link.Id = lnk.Id

	//TODO 启动对应的设备

	return nil
}

func (c *TcpClient) HasAcceptor() bool {
	return false
}

func (c *TcpClient) Close() error {
	if c.link != nil {
		return c.link.Close()
	}
	return nil //TODO return error
}

func (c *TcpClient) GetLink(id int) (Link, error) {
	return c.link, nil
}
