package tunnel

import (
	"errors"
	"github.com/zgwit/iot-master/v4/pkg/log"
	"go.bug.st/serial"
	"time"
)

// Serial 串口
type Serial struct {
	Base         `xorm:"extends"`
	RetryOptions `xorm:"extends"`

	PortName   string `json:"port_name,omitempty"`   //port, e.g. COM1 "/dev/ttySerial1".
	BaudRate   uint   `json:"baud_rate,omitempty"`   //9600 115200
	DataBits   uint   `json:"data_bits,omitempty"`   //5 6 7 8
	StopBits   uint   `json:"stop_bits,omitempty"`   //1 2
	ParityMode int    `json:"parity_mode,omitempty"` //0 1 2 NONE ODD EVEN
}

// Open 打开
func (s *Serial) Open() error {
	if s.Running {
		return errors.New("serial is opened")
	}
	s.closed = false

	opts := serial.Mode{
		BaudRate: int(s.BaudRate),
		DataBits: int(s.DataBits),
		StopBits: serial.StopBits(s.StopBits),
		Parity:   serial.Parity(s.ParityMode),
	}

	port, err := serial.Open(s.PortName, &opts)
	if err != nil {
		//TODO 串口重试
		s.Retry()
		return err
	}

	//读超时
	err = port.SetReadTimeout(time.Second * 5)
	if err != nil {
		return err
	}

	s.Running = true
	s.conn = port

	//清空重连计数
	//s.retry = 0

	//守护协程
	go func() {
		timeout := s.RetryOptions.RetryTimeout
		if timeout == 0 {
			timeout = 10
		}
		for {
			time.Sleep(time.Second * time.Duration(timeout))
			if s.Running {
				continue
			}
			if s.closed {
				return
			}

			//如果掉线了，就重新打开
			err := s.Open()
			if err != nil {
				log.Error(err)
			}
			break //Open中，会重新启动协程
		}
	}()

	//启动轮询
	return s.start()
}

func (s *Serial) Retry() {
	retry := &s.RetryOptions
	if retry.RetryMaximum == 0 || s.retry < retry.RetryMaximum {
		s.retry++
		timeout := retry.RetryTimeout
		if timeout == 0 {
			timeout = 10
		}
		s.retryTimer = time.AfterFunc(time.Second*time.Duration(timeout), func() {
			s.retryTimer = nil
			err := s.Open()
			if err != nil {
				log.Error(err)
			}
		})
	}
}
