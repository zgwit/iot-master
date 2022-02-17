package connect

import (
	"github.com/asdine/storm/v3"
	"github.com/zgwit/iot-master/database"
	"github.com/zgwit/iot-master/events"
	"github.com/zgwit/iot-master/model"
	"net"
	"time"
)

//NetClient 网络链接
type NetClient struct {
	events.EventEmitter

	service *model.Tunnel
	link    *NetLink
	net     string
}

func newNetClient(service *model.Tunnel, net string) *NetClient {
	return &NetClient{
		service: service,
		net:     net,
	}
}

//Open 打开
func (client *NetClient) Open() error {
	client.Emit("open")

	//发起连接
	conn, err := net.Dial(client.net, client.service.Addr)
	if err != nil {
		return err
	}
	client.link = newNetLink(conn)
	go client.link.receive()

	//Store link
	lnk := model.Link{
		TunnelID: client.service.ID,
		Created:  time.Now(),
	}
	err = database.Master.One("TunnelID", client.service.ID, &lnk)
	if err == storm.ErrNotFound {
		//保存一条新记录
		_ = database.Master.Save(&lnk)
	} else if err != nil {
		return err
	} else {
		//上线
	}
	client.link.id = lnk.ID

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

//Close 关闭
func (client *NetClient) Close() error {
	//记录启动
	client.Emit("close")

	if client.link != nil {
		return client.link.Close()
	}
	return nil //TODO return error
}

//GetLink 获取链接
func (client *NetClient) GetLink(id int) (Link, error) {
	return client.link, nil
}
