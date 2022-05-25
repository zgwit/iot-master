package connect

import (
	"errors"
	"github.com/jacobsa/go-serial/serial"
	"github.com/zgwit/iot-master/database"
	"github.com/zgwit/iot-master/events"
	"github.com/zgwit/iot-master/model"
	"github.com/zgwit/storm/v3"
	"time"
)

//Serial 串口
type Serial struct {
	events.EventEmitter

	tunnel *model.Tunnel

	link  *SerialLink

	retry int
	running bool
}

func newSerial(tunnel *model.Tunnel) *Serial {
	return &Serial{
		tunnel: tunnel,
	}
}

//Open 打开
func (s *Serial) Open() error {
	if s.running {
		return errors.New("serial is opened")
	}
	s.Emit("open")

	options := serial.OpenOptions{
		PortName:   s.tunnel.Addr,
		BaudRate:   s.tunnel.Serial.BaudRate,
		DataBits:   s.tunnel.Serial.DataBits,
		StopBits:   s.tunnel.Serial.StopBits,
		ParityMode: serial.ParityMode(s.tunnel.Serial.ParityMode),
		Rs485Enable: s.tunnel.Serial.RS485,
		MinimumReadSize: 4, //避免只读了一个字节就返回
	}
	port, err := serial.Open(options)
	if err != nil {
		//TODO 串口重试
		return err
	}

	s.running = true

	//清空重连计数
	s.retry = 0

	s.link = newSerialLink(port)
	go s.link.receive()

	//Store link
	lnk := model.Link{TunnelId: s.tunnel.Id, Last: time.Now(), Remote: s.tunnel.Addr}
	err = database.Master.One("TunnelId", s.tunnel.Id, &lnk)
	if err == storm.ErrNotFound {
		//保存一条新记录
		_ = database.Master.Save(&lnk)
		s.link.first = true
	} else if err != nil {
		return err
	} else {
		//上线
		_= database.Master.UpdateField(&lnk, "Last", time.Now())
		_= database.Master.UpdateField(&lnk, "Remote", s.tunnel.Addr)
	}
	s.link.id = lnk.Id

	s.Emit("link", s.link)

	//断线后，要重连
	s.link.Once("close", func() {
		//已经关闭，则不再重连
		if !s.running {
			return
		}
		s.running = false

		retry := s.tunnel.Retry
		if retry.Enable && (retry.Maximum == 0 || s.retry < retry.Maximum) {
			s.retry++
			time.AfterFunc(time.Second*time.Duration(retry.Timeout), func() {
				_ = s.Open()
			})
		}
	})

	return nil
}

//Close 关闭
func (s *Serial) Close() error {
	s.Emit("close")
	if s.link != nil {
		link := s.link
		s.link = nil
		return link.Close()
	}
	s.running = false
	return errors.New("link is closed")
}

func (s *Serial) GetLink(id int64) Link {
	return s.link
}

func (s *Serial) Running() bool {
	return s.running
}
