package broker

import (
	"bytes"
	mqtt "github.com/mochi-mqtt/server/v2"
	"github.com/mochi-mqtt/server/v2/packets"
	"github.com/zgwit/iot-master/v4/db"
	"github.com/zgwit/iot-master/v4/log"
	"github.com/zgwit/iot-master/v4/pkg/vconn"
)

type Auth struct {
	mqtt.HookBase
}

// ID returns the ID of the hook.
func (h *Auth) ID() string {
	return "allow-gateway-auth"
}

// Provides indicates which hook methods this hook provides.
func (h *Auth) Provides(b byte) bool {
	return bytes.Contains([]byte{
		mqtt.OnConnectAuthenticate,
		mqtt.OnACLCheck,
	}, []byte{b})
}

// OnConnectAuthenticate returns true/allowed for all requests.
func (h *Auth) OnConnectAuthenticate(cl *mqtt.Client, pk packets.Packet) bool {

	if cl.Net.Inline {
		return true
	}

	//内部虚拟连接
	if _, ok := cl.Net.Conn.(*vconn.VConn); ok {
		return true
	}

	//Websocket连接
	if _, ok := cl.Net.Conn.(*wsConn); ok {
		return true
	}

	//根据网关ID，查密码
	var gw Gateway
	has, err := db.Engine.ID(cl.ID).Get(&gw)
	if err != nil {
		log.Error(err)
		return false
	}
	if !has {
		return false
	}
	if gw.Username == "" || gw.Password == "" {
		return true
	}

	return gw.Username == string(pk.Connect.Username) && gw.Password == string(pk.Connect.Password)
}

// OnACLCheck returns true/allowed for all checks.
func (h *Auth) OnACLCheck(cl *mqtt.Client, topic string, write bool) bool {
	return true
}
