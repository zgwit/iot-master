package tunnel

import (
	"errors"
	"github.com/zgwit/iot-master/v4/connect"
	"github.com/zgwit/iot-master/v4/device"
	"github.com/zgwit/iot-master/v4/protocol"
	"io"
	"sync"
	"time"
)

type PollerOptions struct {
	PollerPeriod   uint `json:"poller_period,omitempty"`   //采集周期
	PollerInterval uint `json:"poller_interval,omitempty"` //采集间隔
}

type ProtocolOptions struct {
	ProtocolName    string         `json:"protocol_name,omitempty"`    //协议 rtu tcp parallel-tcp
	ProtocolOptions map[string]any `json:"protocol_options,omitempty"` //协议参数
}

type RetryOptions struct {
	RetryTimeout uint `json:"retry_timeout,omitempty"` //重试时间
	RetryMaximum uint `json:"retry_maximum,omitempty"` //最大次数
}

type Base struct {
	Id          string `json:"id,omitempty" xorm:"pk"` //ID
	Name        string `json:"name,omitempty"`         //名称
	Description string `json:"description,omitempty"`  //说明

	Heartbeat string `json:"heartbeat,omitempty"` //心跳包

	ProtocolOptions `xorm:"extends"`
	PollerOptions   `xorm:"extends"`

	Disabled bool      `json:"disabled"`
	Created  time.Time `json:"created" xorm:"created"` //创建时间

	Running bool `json:"Running,omitempty" xorm:"-"`

	connect.Conn

	adapter device.Adapter

	lock sync.Mutex

	closed bool

	retry      uint
	retryTimer *time.Timer

	//透传
	pipe io.ReadWriteCloser
}

// Close 关闭
func (l *Base) Close() error {
	if l.retryTimer != nil {
		l.retryTimer.Stop()
	}
	if !l.Running {
		return errors.New("Tunnel closed")
	}

	l.closed = true

	l.onClose()
	return l.Conn.Close()
}

func (l *Base) onClose() {
	l.Running = false
	if l.pipe != nil {
		_ = l.pipe.Close()
	}
}

// Write 写
func (l *Base) Write(data []byte) (int, error) {
	if !l.Running {
		return 0, errors.New("model closed")
	}
	if l.pipe != nil {
		return 0, nil //透传模式下，直接抛弃
	}
	return l.Conn.Write(data)
}

// Read 读
func (l *Base) Read(data []byte) (int, error) {
	if !l.Running {
		return 0, errors.New("model closed")
	}

	if l.pipe != nil {
		//TODO 先read，然后透传
		return 0, nil //透传模式下，直接抛弃
	}
	return l.Conn.Read(data)
}

func (l *Base) start() (err error) {
	l.adapter, err = protocol.Create(l.Id, l.Conn, l.ProtocolName, l.ProtocolOptions.ProtocolOptions)
	return
}

func (l *Base) Pipe(pipe io.ReadWriteCloser) {
	//关闭之前的透传
	if l.pipe != nil {
		_ = l.pipe.Close()
	}

	l.pipe = pipe
	//传入空，则关闭
	if pipe == nil {
		return
	}

	buf := make([]byte, 1024)
	for {
		n, err := pipe.Read(buf)
		if err != nil {
			//if err == io.EOF {
			//	continue
			//}
			//pipe关闭，则不再透传
			break
		}
		//将收到的数据转发出去
		n, err = l.Conn.Write(buf[:n])
		if err != nil {
			//发送失败，说明连接失效
			_ = pipe.Close()
			break
		}
	}
	l.pipe = nil

	//TODO 使用io.copy
	//go io.Copy(pipe, l.conn)
	//go io.Copy(l.conn, pipe)
}
