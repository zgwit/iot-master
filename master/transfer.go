package master

import (
	"errors"
	"fmt"
	"iot-master/db"
	"iot-master/log"
	"iot-master/model"
	"math"
	"net"
	"sync"
)

var allTransfers sync.Map

type Transfer struct {
	model.Transfer
	conn     *net.TCPConn
	listener *net.TCPListener

	running bool
}

func (t *Transfer) Open() error {
	if t.running {
		return errors.New("服务已经运行")
	}

	addr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf(":%d", t.Port))
	if err != nil {
		return err
	}
	t.listener, err = net.ListenTCP("tcp", addr)
	if err != nil {
		return err
	}

	tunnel := GetTunnel(t.TunnelId)
	if tunnel == nil {
		return fmt.Errorf("连接 %d 找不到", t.TunnelId)
	}

	go func() {
		t.running = true
		for {
			t.conn, err = t.listener.AcceptTCP()
			if err != nil {
				break
			}
			tunnel.Instance.Pipe(t.conn)
		}
		t.running = false
	}()
	return nil
}

func (t *Transfer) Close() error {
	if !t.running {
		return errors.New("服务已经关闭")
	}
	if t.conn != nil {
		_ = t.conn.Close()
	}
	return t.listener.Close()
}

func (t *Transfer) Running() bool {
	return t.running
}

func LoadTransfers() error {
	var transfers []*model.Transfer
	err := db.Engine.Limit(math.MaxInt).Find(&transfers)
	if err != nil {
		return err
	}
	for _, transfer := range transfers {
		if transfer.Disabled {
			continue
		}
		p := &Transfer{Transfer: *transfer}
		err = p.Open()
		if err != nil {
			log.Error(err)
		}
	}
	return nil
}

func LoadTransfer(id int64) error {
	var transfer model.Transfer
	has, err := db.Engine.ID(id).Get(&transfer)
	if err != nil {
		return err
	}
	if !has {
		return fmt.Errorf("找不到透传 %d", id)
	}

	if transfer.Disabled {
		return nil //TODO error ??
	}

	p := &Transfer{Transfer: transfer}
	err = p.Open()
	if err != nil {
		return err
	}
	allTransfers.Store(transfer.Id, p)
	return nil
}

func GetTransfer(id int64) *Transfer {
	d, ok := allTransfers.Load(id)
	if ok {
		return d.(*Transfer)
	}
	return nil
}

func RemoveTransfer(id int64) error {
	d, ok := allTransfers.LoadAndDelete(id)
	if ok {
		dev := d.(*Transfer)
		return dev.Close()
	}
	return nil //error
}
