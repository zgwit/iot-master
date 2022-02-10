package service

import (
	"github.com/asdine/storm/v3"
	"github.com/zgwit/iot-master/common"
	"github.com/zgwit/iot-master/database"
	"github.com/zgwit/iot-master/model"
	"net"
	"time"
)

type NetClient struct {
	common.EventEmitter

	service *model.Service
	link    *NetLink
	net     string
}

func newNetClient(service *model.Service, net string) *NetClient {
	return &NetClient{
		service: service,
		net:     net,
	}
}

func (client *NetClient) Open() error {
	conn, err := net.Dial(client.net, client.service.Addr)
	if err != nil {
		return err
	}
	client.link = newNetLink(conn)
	go client.link.receive()

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

	client.Emit("link", client.link)

	client.link.Once("close", func() {
		retry := client.service.Retry
		if retry == 0 {
			retry = 10 //默认10秒重试
		}
		time.AfterFunc(time.Second*time.Duration(retry), func() {
			_ = client.Open()
		})
	})

	return nil
}

func (client *NetClient) Close() error {
	if client.link != nil {
		return client.link.Close()
	}
	return nil //TODO return error
}

func (client *NetClient) GetLink(id int) (Link, error) {
	return client.link, nil
}
