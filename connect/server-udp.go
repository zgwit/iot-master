package connect

import (
	"errors"
	"github.com/zgwit/iot-master/db"
	"github.com/zgwit/iot-master/events"
	"github.com/zgwit/iot-master/model"
	"net"
	"time"
)

//ServerUDP UDP服务器
type ServerUDP struct {
	events.EventEmitter

	server *model.Server

	children map[int64]*ServerUdpTunnel
	tunnels  map[string]*ServerUdpTunnel

	listener *net.UDPConn
	running  bool
}

func newServerUDP(server *model.Server) *ServerUDP {
	svr := &ServerUDP{
		server:   server,
		children: make(map[int64]*ServerUdpTunnel),
		tunnels:  make(map[string]*ServerUdpTunnel),
	}
	return svr
}

//Open 打开
func (server *ServerUDP) Open() error {
	if server.running {
		return errors.New("server is opened")
	}
	server.Emit("open")

	addr, err := net.ResolveUDPAddr("udp", resolvePort(server.server.Addr))
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
			tunnel, ok := server.tunnels[addr.String()]
			if ok {
				tunnel.onData(data)
				continue
			}

			var tnl model.Tunnel

			if !server.server.Register.Check(data) {
				_ = conn.Close()
				continue
			}
			sn := string(data)
			tnl.SN = sn
			has, err := db.Engine.Where("server_id=?", server.server.Id).And("sn", sn).Get(&tnl)
			if err != nil {
				//return err
				//TODO 日志，关闭连接
				continue
			}

			if !has {
				//保存一条新记录
				tnl = model.Tunnel{ServerId: server.server.Id, Last: time.Now(), Remote: conn.RemoteAddr().String()}
				_, _ = db.Engine.InsertOne(&tnl)
			} else {
				//上线
				tnl.Last = time.Now()
				tnl.Remote = conn.RemoteAddr().String()
				_, _ = db.Engine.ID(tnl.Id).Cols("last", "remote").Update(tnl)
			}

			tunnel = newUdpLink(&tnl, conn, addr)
			tunnel.first = !has
			server.children[tnl.Id] = tunnel

			//启动对应的设备 发消息
			server.Emit("tunnel", tunnel)

			tunnel.Once("close", func() {
				delete(server.children, tnl.Id)
				delete(server.tunnels, tunnel.addr.String())
			})
		}

		server.running = false
	}()

	return nil
}

//Close 关闭
func (server *ServerUDP) Close() (err error) {
	server.Emit("close")
	//close tunnels
	if server.children != nil {
		for _, l := range server.children {
			_ = l.Close()
		}
	}
	return server.listener.Close()
}

//GetTunnel 获取链接
func (server *ServerUDP) GetTunnel(id int64) Tunnel {
	return server.children[id]
}

func (server *ServerUDP) Running() bool {
	return server.running
}
