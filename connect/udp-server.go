package connect

import (
	"errors"
	"github.com/zgwit/iot-master/database"
	"github.com/zgwit/iot-master/events"
	"github.com/zgwit/iot-master/model"
	"github.com/zgwit/storm/v3"
	"github.com/zgwit/storm/v3/q"
	"net"
	"time"
)

//UdpServer UDP服务器
type UdpServer struct {
	events.EventEmitter

	tunnel *model.Tunnel

	children map[int]*UdpLink
	links    map[string]*UdpLink

	listener *net.UDPConn
	running  bool
}

func newUdpServer(tunnel *model.Tunnel) *UdpServer {
	svr := &UdpServer{
		tunnel: tunnel,
	}
	if tunnel.Register.Enable {
		svr.children = make(map[int]*UdpLink)
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

			if !server.tunnel.Register.Enable {
				//先结束其他链接
				for _, link := range server.links {
					_ = link.Close()
				}
				err = database.Master.One("TunnelId", server.tunnel.Id, &lnk)
			} else {
				if !server.tunnel.Register.Check(data) {
					_ = conn.Close()
					continue
				}
				sn := string(data)
				lnk.SN = sn
				err = database.Master.Select(q.Eq("TunnelId", server.tunnel.Id), q.Eq("SN", sn)).First(&lnk)
			}

			link = newUdpLink(conn, addr)

			if err == storm.ErrNotFound {
				//保存一条新记录
				_ = database.Master.Save(&lnk)
				link.first = true
			} else if err != nil {
				//return err
				continue
			} else {
				//上线
				_ = database.Master.UpdateField(&lnk, "Last", time.Now())
				_ = database.Master.UpdateField(&lnk, "Remote", conn.RemoteAddr().String())
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
func (server *UdpServer) GetLink(id int) Link {
	return server.children[id]
}

func (server *UdpServer) Running() bool {
	return server.running
}