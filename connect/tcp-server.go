package connect

import (
	"errors"
	"github.com/zgwit/iot-master/db"
	"github.com/zgwit/iot-master/events"
	"github.com/zgwit/iot-master/model"
	"net"
	"time"
)

//TcpServer TCP服务器
type TcpServer struct {
	events.EventEmitter

	tunnel *model.Tunnel

	children map[int64]*NetLink

	listener *net.TCPListener

	running bool
}

func newTcpServer(tunnel *model.Tunnel) *TcpServer {
	svr := &TcpServer{
		tunnel: tunnel,
	}
	if tunnel.Register.Enable {
		svr.children = make(map[int64]*NetLink)
	}
	return svr
}

//Open 打开
func (server *TcpServer) Open() error {
	if server.running {
		return errors.New("server is opened")
	}
	server.Emit("open")

	addr, err := net.ResolveTCPAddr("tcp", resolvePort(server.tunnel.Addr))
	if err != nil {
		return err
	}
	server.listener, err = net.ListenTCP("tcp", addr)
	if err != nil {
		return err
	}

	server.running = true
	go func() {
		for {
			conn, err := server.listener.AcceptTCP()
			if err != nil {
				//TODO 需要正确处理接收错误
				break
			}

			lnk := model.Link{TunnelId: server.tunnel.Id, Last: time.Now(), Remote: conn.RemoteAddr().String()}

			has := false
			if !server.tunnel.Register.Enable {
				//先结束历史链接
				for _, link := range server.children {
					_ = link.Close()
				}
				has, err = db.Engine.Where("tunnel_id=?", server.tunnel.Id).Exist(&lnk)
			} else {
				buf := make([]byte, 128)
				n := 0
				n, err = conn.Read(buf)
				if err != nil {
					_ = conn.Close()
					continue
				}
				data := buf[:n]
				if !server.tunnel.Register.Check(data) {
					_ = conn.Close()
					continue
				}
				sn := string(data)
				lnk.SN = sn
				has, err = db.Engine.Where("tunnel_id=?", server.tunnel.Id).And("sn", sn).Exist(&lnk)
			}

			if err != nil {
				//return err
				//TODO 日志，关闭连接
				continue
			}

			if !has {
				//保存一条新记录
				_, _ = db.Engine.InsertOne(&lnk)
			} else {
				//上线
				lnk.Last = time.Now()
				lnk.Remote = conn.RemoteAddr().String()
				_, _ = db.Engine.ID(lnk.Id).Cols("last", "remote").Update(lnk)
			}

			link := newNetLink(conn)
			go link.receive()

			link.id = lnk.Id
			server.children[lnk.Id] = link

			//启动对应的设备 发消息
			server.Emit("link", link)

			link.Once("close", func() {
				delete(server.children, link.id)
			})
		}

		server.running = false
	}()

	return nil
}

//Close 关闭
func (server *TcpServer) Close() (err error) {
	server.Emit("close")

	//close links
	if server.children != nil {
		for _, l := range server.children {
			_ = l.Close()
		}
	}
	return server.listener.Close()
}

//GetLink 获取连接
func (server *TcpServer) GetLink(id int64) Link {
	return server.children[id]
}

func (server *TcpServer) Running() bool {
	return server.running
}