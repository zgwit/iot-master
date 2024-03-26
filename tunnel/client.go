package tunnel

import (
	"errors"
	"fmt"
	"github.com/zgwit/iot-master/v4/pkg/log"
	"net"
	"time"
)

// Client 网络链接
type Client struct {
	Base         `xorm:"extends"`
	RetryOptions `xorm:"extends"`

	Net  string `json:"net,omitempty"`  //类型 tcp udp
	Addr string `json:"addr,omitempty"` //地址，主机名或IP
	Port uint16 `json:"port,omitempty"` //端口号

}

// Open 打开
func (client *Client) Open() error {
	if client.Running {
		return errors.New("client is opened")
	}
	client.closed = false

	//发起连接
	addr := fmt.Sprintf("%s:%d", client.Addr, client.Port)
	conn, err := net.Dial(client.Net, addr)
	if err != nil {
		client.Retry()
		//time.AfterFunc(time.Minute, client.RetryOptions)
		return err
	}
	client.retry = 0
	client.conn = &netConn{conn}

	//守护协程
	go func() {
		timeout := client.RetryTimeout
		if timeout == 0 {
			timeout = 10
		}
		for {
			time.Sleep(time.Second * time.Duration(timeout))
			if client.Running {
				continue
			}
			if client.closed {
				return
			}
			//如果掉线了，就重新打开
			err := client.Open()
			if err != nil {
				log.Error(err)
			}
			break //Open中，会重新启动协程
		}
	}()

	//启动轮询
	return client.start()
}

func (client *Client) Retry() {
	//重连
	retry := &client.RetryOptions
	if retry.RetryMaximum == 0 || client.retry < retry.RetryMaximum {
		client.retry++
		timeout := retry.RetryTimeout
		if timeout == 0 {
			timeout = 10
		}
		client.retryTimer = time.AfterFunc(time.Second*time.Duration(timeout), func() {
			client.retryTimer = nil
			err := client.Open()
			if err != nil {
				log.Error(err)
			}
		})
	}
}

// Close 关闭
func (client *Client) Close() error {
	client.Running = false

	if client.conn != nil {
		link := client.conn
		client.conn = nil
		return link.Close()
	}
	return errors.New("model is closed")
}
