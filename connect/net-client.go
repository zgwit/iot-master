package connect

import (
	"github.com/asdine/storm/v3"
	"github.com/zgwit/iot-master/database"
	"github.com/zgwit/iot-master/events"
	"net"
	"time"
)

type NetClient struct {
	events.EventEmitter

	service *TunnelModel
	link    *NetConn
	net     string
}

func newNetClient(service *TunnelModel, net string) *NetClient {
	return &NetClient{
		service: service,
		net:     net,
	}
}

func (client *NetClient) Open() error {
	client.Emit("open")

	//发起连接
	conn, err := net.Dial(client.net, client.service.Addr)
	if err != nil {
		return err
	}
	client.link = newNetConn(conn)
	go client.link.receive()

	//Store link
	lnk := LinkModel{
		TunnelId: client.service.Id,
		Created:  time.Now(),
	}
	err = database.Link.One("TunnelId", client.service.Id, &lnk)
	if err == storm.ErrNotFound {
		//保存一条新记录
		_ = database.Link.Save(&lnk)
	} else if err != nil {
		return err
	} else {
		//上线
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
	//记录启动
	client.Emit("close")

	if client.link != nil {
		return client.link.Close()
	}
	return nil //TODO return error
}

func (client *NetClient) GetLink(id int) (Link, error) {
	return client.link, nil
}
