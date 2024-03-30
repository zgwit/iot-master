package tunnel

import (
	"errors"
	"fmt"
	"github.com/zgwit/iot-master/v4/pkg/db"
	"github.com/zgwit/iot-master/v4/pkg/log"
	"github.com/zgwit/iot-master/v4/protocol"
	"net"
)

func init() {
	db.Register(new(Server))
}

// Server TCP服务器
type Server struct {
	Base `xorm:"extends"`

	Port       uint16 `json:"port,omitempty"`       //监听端口
	Standalone bool   `json:"standalone,omitempty"` //单例模式（不支持注册）

	children map[string]*Link

	listener *net.TCPListener
}

func (s *Server) handleStandalone(c *net.TCPConn) (err error) {
	const k = "internal"
	if cc, ok := s.children[k]; ok {
		_ = cc.Close()
	}

	l := &Link{
		Base:     s.Base, //TODO 删除lock
		ServerId: s.Id,
		Remote:   c.RemoteAddr().String(),
	}

	s.children[k] = l
	//以ServerID保存
	links.Store(s.Id, l)

	//启动轮询
	l.adapter, err = protocol.Create(l.Id, l, l.ProtocolName, l.ProtocolOptions.ProtocolOptions)
	return err
}

func (s *Server) handleIncoming(c *net.TCPConn) error {
	//TODO 只有配置了注册包，才能正常通讯
	buf := make([]byte, 128)
	n, err := c.Read(buf)
	if err != nil {
		_ = c.Close()
		return err
	}

	data := buf[:n]
	sn := string(data)

	var l Link
	//get, err := db.Engine.Where("server_id=?", s.Id).And("sn=?", sn).Get(&conn)
	get, err := db.Engine.ID(sn).Get(&l)
	if err != nil {
		_, _ = c.Write([]byte(err.Error()))
		_ = c.Close()
		return err
	}

	if !get {
		l = Link{
			Base:     s.Base,
			ServerId: s.Id,
			Remote:   c.RemoteAddr().String(),
		}
		l.Id = sn //修改ID

		_, err = db.Engine.InsertOne(&l)
		if err != nil {
			_, _ = c.Write([]byte(err.Error()))
			_ = c.Close()
			return err
		}
	}

	s.children[sn] = &l
	links.Store(l.Id, &l)

	//启动轮询
	l.adapter, err = protocol.Create(l.Id, &l, l.ProtocolName, l.ProtocolOptions.ProtocolOptions)
	return err
}

// Open 打开
func (s *Server) Open() error {
	if s.running {
		return errors.New("server is opened")
	}
	s.closed = false

	log.Trace("create server", s.Port)
	addr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf(":%d", s.Port))
	if err != nil {
		return err
	}

	//守护
	if s.keeper == nil {
		s.keeper = Keep(s)
	}

	s.listener, err = net.ListenTCP("tcp", addr)
	if err != nil {
		return err
	}
	s.running = true

	go func() {
		for {
			c, err := s.listener.AcceptTCP()
			if err != nil {
				//TODO 需要正确处理接收错误
				log.Error(err)
				break
			}

			//单例模式，关闭之前的连接
			if s.Standalone {
				err = s.handleStandalone(c)
				if err != nil {
					log.Error(err)
				}
				continue
			}

		}

		s.running = false
	}()

	return nil
}

// Close 关闭
func (s *Server) Close() error {
	s.running = false
	s.closed = true

	//close tunnels
	if s.children != nil {
		for _, l := range s.children {
			_ = l.Close()
		}
	}

	return s.listener.Close()
}

// GetTunnel 获取连接
func (s *Server) GetTunnel(id string) *Link {
	return s.children[id]
}
