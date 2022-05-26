package connect

import (
	"errors"
	"github.com/zgwit/iot-master/db"
	"github.com/zgwit/iot-master/events"
	"github.com/zgwit/iot-master/model"
	"net"
	"time"
)

//UdpServer UDP服务器
type UdpServer struct {
	events.EventEmitter

	tunnel *model.Tunnel

	children map[int64]*UdpLink
	links    map[string]*UdpLink

	listener *net.UDPConn
	running  bool
}

func newUdpServer(tunnel *model.Tunnel) *UdpServer {
	svr := &UdpServer{
		tunnel: tunnel,
	}
	if tunnel.Register.Enable {
		svr.children = make(map[int64]*UdpLink)
		svr.links = make(map[string]*UdpLink)
	}
	return svr
}

//Open 打开
func (server *UdpServer) Open() error {
	if server.running {
		return errors.New("server is opened")
	}
	server.Emit("open")

	addr, err := net.ResolveUDPAddr("udp", resolvePort(server.tunnel.Addr))
	if err != nil {
		return err
	}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		//TODO 需要正确处理接收错误
		return err
	}
	server.listener = conn //共用连接

	server.running = true
	go func() {
		for {
			buf := make([]byte, 1024)
			n, addr, err := conn.ReadFromUDP(buf)
			if err != nil {
				_ = conn.Close()
				//continue
				break
			}
			data := buf[:n]

			//如果已经保存了链接 TODO 要有超时处理
			link, ok := server.links[addr.String()]
			if ok {
				link.onData(data)
				continue
			}

			lnk := model.Link{TunnelId: server.tunnel.Id, Last: time.Now(), Remote: conn.RemoteAddr().String()}

			has := false
			if !server.tunnel.Register.Enable {
				//先结束其他链接
				for _, link := range server.links {
					_ = link.Close()
				}
				has, err = db.Engine.Where("tunnel_id=?", server.tunnel.Id).Get(&lnk)
			} else {
				if !server.tunnel.Register.Check(data) {
					_ = conn.Close()
					continue
				}
				sn := string(data)
				lnk.SN = sn
				has, err = db.Engine.Where("tunnel_id=?", server.tunnel.Id).And("sn", sn).Get(&lnk)
			}

			if err != nil {
				//return err
				//TODO 日志，关闭连接
				continue
			}

			link = newUdpLink(conn, addr)

			if !has {
				//保存一条新记录
				_, _ = db.Engine.InsertOne(&lnk)
				link.first = true
			} else {
				//上线
				lnk.Last = time.Now()
				lnk.Remote = conn.RemoteAddr().String()
				_, _ = db.Engine.ID(lnk.Id).Cols("last", "remote").Update(lnk)
			}

			link.id = lnk.Id
			server.children[lnk.Id] = link

			//启动对应的设备 发消息
			server.Emit("link", link)

			link.Once("close", func() {
				delete(server.children, link.id)
				delete(server.links, link.addr.String())
			})
		}

		server.running = false
	}()

	return nil
}

//Close 关闭
func (server *UdpServer) Close() (err error) {
	server.Emit("close")
	//close links
	if server.children != nil {
		for _, l := range server.children {
			_ = l.Close()
		}
	}
	return server.listener.Close()
}

//GetLink 获取链接
func (server *UdpServer) GetLink(id int64) Link {
	return server.children[id]
}

func (server *UdpServer) Running() bool {
	return server.running
}