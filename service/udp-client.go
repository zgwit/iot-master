package service

import (
	"github.com/asaskevich/EventBus"
	"github.com/asdine/storm/v3"
	"github.com/zgwit/iot-master/database"
	"github.com/zgwit/iot-master/model"
	"net"
	"time"
)

type UdpClient struct {
	service *model.Service
	link    *UdpLink
	events  EventBus.Bus
}

func NewUdpClient(service *model.Service) *UdpClient {
	return &UdpClient{
		service: service,
		events: EventBus.New(),
	}
}

func (client *UdpClient) Open() error {
	addr ,err := net.ResolveUDPAddr("udp", client.service.Addr)
	if err != nil {
		return err
	}
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		return err
	}
	client.link = newUdpLink(conn)
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


func (client *UdpClient) Close() error {
	if client.link != nil {
		return client.link.Close()
	}
	return nil //TODO return error
}

func (client *UdpClient) GetLink(id int) (Link, error) {
	return client.link, nil
}

func (client *UdpClient) OnLink(fn func(link Link)) {
	_ = client.events.Subscribe("link", fn)
}