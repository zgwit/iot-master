package tunnel

import (
	"errors"
	"fmt"
	"github.com/zgwit/iot-master/v4/pkg/db"
	"github.com/zgwit/iot-master/v4/pkg/log"
	"net"
)

func init() {
	db.Register(new(Server))
}

// Server TCP服务器
type Server struct {
	Base         `xorm:"extends"`
	RetryOptions `xorm:"extends"`

	Port       uint16 `json:"port,omitempty"`       //监听端口
	Standalone bool   `json:"standalone,omitempty"` //单例模式（不支持注册）

	children map[string]*Link

	listener *net.TCPListener
}

// Open 打开
func (s *Server) Open() error {
	if s.Running {
		return errors.New("s is opened")
	}

	addr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf(":%d", s.Port))
	if err != nil {
		return err
	}
	s.listener, err = net.ListenTCP("tcp", addr)
	if err != nil {
		return err
	}
	//defer s.listener.Close()

	s.Running = true
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
				const k = "internal"
				if cc, ok := s.children[k]; ok {
					_ = cc.Close()
				}

				lnk := &Link{
					Base:     s.Base, //TODO 删除lock
					ServerId: s.Id,
					Remote:   c.RemoteAddr().String(),
				}

				s.children[k] = lnk

				//启动轮询
				err = lnk.start()
				if err != nil {
					log.Error(err)
					continue
					//return
				}

				//以ServerID保存
				links.Store(s.Id, lnk)
				continue
			}

			//TODO 只有配置了注册包，才能正常通讯
			buf := make([]byte, 128)
			n, err := c.Read(buf)
			if err != nil {
				_ = c.Close()
				continue
			}
			data := buf[:n]
			sn := string(data)

			var link Link
			//get, err := db.Engine.Where("server_id=?", s.Id).And("sn=?", sn).Get(&conn)
			get, err := db.Engine.ID(sn).Get(&link)
			if err != nil {
				_, _ = c.Write([]byte(err.Error()))
				_ = c.Close()
				continue
			}
			if !get {
				link = Link{
					Base:     s.Base,
					ServerId: s.Id,
					Remote:   c.RemoteAddr().String(),
				}
				link.Id = sn //修改ID

				_, err := db.Engine.InsertOne(&link)
				if err != nil {
					_, _ = c.Write([]byte(err.Error()))
					_ = c.Close()
					continue
				}
			}

			s.children[sn] = &link

			//启动轮询
			err = link.start()
			if err != nil {
				log.Error(err)
				continue
			}

			links.Store(link.Id, &link)
		}

		s.Running = false
	}()

	return nil
}

// Close 关闭
func (s *Server) Close() (err error) {
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
