package connect

import (
	"errors"
	"github.com/zgwit/iot-master/database"
	"github.com/zgwit/iot-master/events"
	"github.com/zgwit/iot-master/model"
	"github.com/zgwit/storm/v3"
	"net"
	"time"
)

//NetClient 网络链接
type NetClient struct {
	events.EventEmitter

	tunnel  *model.Tunnel
	link    *NetLink
	net     string
	retry   int
	running bool
}

func newNetClient(tunnel *model.Tunnel, net string) *NetClient {
	return &NetClient{
		tunnel: tunnel,
		net:    net,
	}
}

//Open 打开
func (client *NetClient) Open() error {
	if client.running {
		return errors.New("client is opened")
	}
	client.Emit("open")

	//发起连接
	conn, err := net.Dial(client.net, client.tunnel.Addr)
	if err != nil {
		return err
	}
	client.running = true

	client.link = newNetLink(conn)
	go client.link.receive()

	//Store link
	lnk := model.Link{TunnelId: client.tunnel.Id, Last: time.Now(), Remote: client.tunnel.Addr}
	err = database.Master.One("TunnelId", client.tunnel.Id, &lnk)
	if err == storm.ErrNotFound {
		//保存一条新记录
		_ = database.Master.Save(&lnk)
		client.link.first = true
	} else if err != nil {
		return err
	} else {
		//上线
		_ = database.Master.UpdateField(&lnk, "Last", time.Now())
		_ = database.Master.UpdateField(&lnk, "Remote", client.tunnel.Addr)
	}
	client.link.id = lnk.Id

	client.Emit("link", client.link)

	client.link.Once("close", func() {
		if !client.running {
			return
		}
		client.running = false

		retry := client.tunnel.Retry
		if retry.Enable && (retry.Maximum == 0 || client.retry < retry.Maximum) {
			client.retry++
			time.AfterFunc(time.Second*time.Duration(retry.Timeout), func() {
				_ = client.Open()
			})
		}
	})
	return nil
}

//Close 关闭
func (client *NetClient) Close() error {
	client.running = false

	//记录启动
	client.Emit("close")

	if client.link != nil {
		link := client.link
		client.link = nil
		return link.Close()
	}
	return errors.New("link is closed")
}

//GetLink 获取链接
func (client *NetClient) GetLink(id int) Link {
	return client.link
}

func (client *NetClient) Running() bool {
	return client.running
}
