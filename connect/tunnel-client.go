package connect

import (
	"errors"
	"github.com/zgwit/iot-master/db"
	"github.com/zgwit/iot-master/model"
	"net"
	"time"
)

//TunnelClient 网络链接
type TunnelClient struct {
	tunnelBase
	net string
}

func newTunnelClient(tunnel *model.Tunnel, net string) *TunnelClient {
	return &TunnelClient{
		tunnelBase: tunnelBase{tunnel: tunnel},
		net:        net,
	}
}

//Open 打开
func (client *TunnelClient) Open() error {
	if client.running {
		return errors.New("client is opened")
	}
	client.Emit("open")

	//发起连接
	conn, err := net.Dial(client.net, client.tunnel.Addr)
	if err != nil {
		return err
	}
	client.link = conn

	//开始接收数据
	go client.receive()

	//上线
	client.tunnel.Last = time.Now()
	client.tunnel.Remote = conn.RemoteAddr().String()
	_, _ = db.Engine.ID(client.tunnel.Id).Cols("last", "remote").Update(client.tunnel)

	return nil
}

func (client *TunnelClient) receive() {
	client.running = true
	buf := make([]byte, 1024)
	for {
		n, err := client.link.Read(buf)
		if err != nil {
			client.onClose()
			break
		}
		if n == 0 {
			continue
		}
		//透传转发
		if client.pipe != nil {
			_, err = client.pipe.Write(buf[:n])
			if err != nil {
				client.pipe = nil
			} else {
				continue
			}
		}
		client.Emit("data", buf[:n])
	}
	client.running = false

	//重连
	retry := &client.tunnel.Retry
	if retry.Enable && (retry.Maximum == 0 || client.retry < retry.Maximum) {
		client.retry++
		time.AfterFunc(time.Second*time.Duration(retry.Timeout), func() {
			_ = client.Open()
		})
	}
}

//Close 关闭
func (client *TunnelClient) Close() error {
	client.running = false

	//记录启动
	client.Emit("close")

	if client.link != nil {
		link := client.link
		client.link = nil
		return link.Close()
	}
	return errors.New("tunnel is closed")
}
