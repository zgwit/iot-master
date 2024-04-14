package gateway

import (
	"errors"
	"fmt"
	"github.com/zgwit/iot-master/v4/db"
	"github.com/zgwit/iot-master/v4/log"
	"github.com/zgwit/iot-master/v4/protocol"
	"net"
)

func init() {
	db.Register(new(Client))
}

// Client 网络链接
type Client struct {
	Base `xorm:"extends"`

	Net  string `json:"net,omitempty"`  //类型 tcp udp
	Addr string `json:"addr,omitempty"` //地址，主机名或IP
	Port uint16 `json:"port,omitempty"` //端口号
}

// Open 打开
func (c *Client) Open() error {
	if c.running {
		return errors.New("client is opened")
	}
	c.closed = false

	//守护
	if c.keeper == nil {
		c.keeper = Keep(c)
	}

	//发起连接
	addr := fmt.Sprintf("%s:%d", c.Addr, c.Port)
	log.Trace("create client ", addr)
	conn, err := net.Dial(c.Net, addr)
	if err != nil {
		//time.AfterFunc(time.Minute, client.RetryOptions)
		c.Status = err.Error()
		return err
	}
	c.running = true
	c.Status = "正常"

	c.conn = &netConn{conn}

	//启动轮询
	c.adapter, err = protocol.Create(c, c.ProtocolName, c.ProtocolOptions)

	return err
}
