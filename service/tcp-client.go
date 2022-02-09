package service

import (
	"github.com/asaskevich/EventBus"
	"github.com/asdine/storm/v3"
	"github.com/zgwit/iot-master/database"
	"github.com/zgwit/iot-master/model"
	"net"
	"time"
)

type TcpClient struct {
	service *model.Service
	link    *TcpLink
	events  EventBus.Bus
}

func NewTcpClient(service *model.Service) *TcpClient {
	return &TcpClient{
		service: service,
		events: EventBus.New(),
	}
}

func (client *TcpClient) Open() error {
	addr ,err := net.ResolveTCPAddr("udp", client.service.Addr)
	if err != nil {
		return err
	}
	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		return err
	}
	client.link = newTcpLink(conn)
	//Store link
	lnk := model.Link{
		ServiceId: client.service.Id,
		Created:   time.Now(),
	}
	err = database.Link.One("ServiceId", client.service.Id, &lnk)
	if err == storm.ErrNotFound {
		//保存一条新记录
		_ = database.Link.Save(&lnk)
	} else if err != nil {
		return err
	}
	client.link.Id = lnk.Id

	client.events.Publish("link", client.link)

	return nil
}

func (client *TcpClient) Close() error {
	if client.link != nil {
		return client.link.Close()
	}
	return nil //TODO return error
}

func (client *TcpClient) GetLink(id int) (Link, error) {
	return client.link, nil
}

func (client *TcpClient) OnLink(fn func(link Link)) {
	_ = client.events.Subscribe("link", fn)
}