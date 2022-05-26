package master

import (
	"errors"
	"fmt"
	"github.com/zgwit/iot-master/db"
	"github.com/zgwit/iot-master/log"
	"github.com/zgwit/iot-master/model"
	"math"
	"net"
	"sync"
)

var allPipes sync.Map

type Pipe struct {
	model.Pipe
	conn     *net.TCPConn
	listener *net.TCPListener

	running bool
}

func (p *Pipe) Open() error {
	if p.running {
		return errors.New("服务已经运行")
	}

	addr, err := net.ResolveTCPAddr("tcp", p.Addr)
	if err != nil {
		return err
	}
	p.listener, err = net.ListenTCP("tcp", addr)
	if err != nil {
		return err
	}

	link := GetLink(p.LinkId)
	if link == nil {
		return fmt.Errorf("连接 %d 找不到", p.LinkId)
	}

	go func() {
		p.running = true
		for {
			p.conn, err = p.listener.AcceptTCP()
			if err != nil {
				break
			}
			link.Instance.Pipe(p.conn)
		}
		p.running = false
	}()
	return nil
}

func (p *Pipe) Close() error {
	if !p.running {
		return errors.New("服务已经关闭")
	}
	if p.conn != nil {
		_ = p.conn.Close()
	}
	return p.listener.Close()
}

func (p *Pipe) Running() bool {
	return p.running
}

func LoadPipes() error {
	var pipes []*model.Pipe
	err := db.Engine.Limit(math.MaxInt).Find(&pipes)
	if err != nil {
		return err
	}
	for _, pipe := range pipes {
		if pipe.Disabled {
			continue
		}
		p := &Pipe{Pipe: *pipe}
		err = p.Open()
		if err != nil {
			log.Error(err)
		}
	}
	return nil
}

func LoadPipe(id int64) error {
	var pipe model.Pipe
	has, err := db.Engine.ID(id).Exist(&pipe)
	if err != nil {
		return err
	}
	if !has {
		return fmt.Errorf("找不到透传 %d", id)
	}

	if pipe.Disabled {
		return nil //TODO error ??
	}

	p := &Pipe{Pipe: pipe}
	err = p.Open()
	if err != nil {
		return err
	}
	allPipes.Store(pipe.Id, p)
	return nil
}

func GetPipe(id int64) *Pipe {
	d, ok := allPipes.Load(id)
	if ok {
		return d.(*Pipe)
	}
	return nil
}

func RemovePipe(id int64) error {
	d, ok := allPipes.LoadAndDelete(id)
	if ok {
		dev := d.(*Pipe)
		return dev.Close()
	}
	return nil //error
}
