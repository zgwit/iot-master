package broker

import (
	"errors"
	"github.com/gorilla/websocket"
	"github.com/mochi-co/mqtt/server/listeners"
	"github.com/mochi-co/mqtt/server/listeners/auth"
	"github.com/mochi-co/mqtt/server/system"
	"net"
	"net/http"
)

var (
	// ErrInvalidMessage indicates that a message payload was not valid.
	ErrInvalidMessage = errors.New("message type not binary")

	// wsUpgrader is used to upgrade the incoming http/tcp connection to a
	// websocket compliant connection.
	wsUpgrader = &websocket.Upgrader{
		Subprotocols: []string{"mqtt"},
		CheckOrigin:  func(r *http.Request) bool { return true },
	}
)

/* Copy from github.com/mochi-co/mqtt/server/listeners/websocket.go */

// wsConn is a websocket connection which satisfies the net.Conn interface.
// Inspired by
type wsConn struct {
	net.Conn
	c *websocket.Conn
}

// Read reads the next span of bytes from the websocket connection and returns
// the number of bytes read.
func (ws *wsConn) Read(p []byte) (n int, err error) {
	op, r, err := ws.c.NextReader()
	if err != nil {
		return
	}

	if op != websocket.BinaryMessage {
		err = ErrInvalidMessage
		return
	}

	return r.Read(p)
}

// Write writes bytes to the websocket connection.
func (ws *wsConn) Write(p []byte) (n int, err error) {
	err = ws.c.WriteMessage(websocket.BinaryMessage, p)
	if err != nil {
		return
	}

	return len(p), nil
}

// Close signals the underlying websocket conn to close.
func (ws *wsConn) Close() error {
	return ws.Conn.Close()
}

type WebSocket struct {
	id        string
	config    *listeners.Config
	establish listeners.EstablishFunc
	end       uint32
}

func (l *WebSocket) SetConfig(config *listeners.Config) {
	if config != nil {
		l.config = config
		if l.config.Auth == nil {
			l.config.Auth = new(auth.Disallow)
		}
	}
}

func (l *WebSocket) Listen(s *system.Info) error {
	return nil
}

func (l *WebSocket) Serve(establish listeners.EstablishFunc) {
	l.establish = establish
}

func (l *WebSocket) ID() string {
	return l.id
}

func (l *WebSocket) Close(closeClients listeners.CloseFunc) {
	closeClients(l.id)
}

func (l *WebSocket) Upgrade(w http.ResponseWriter, r *http.Request) {
	c, err := wsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	defer c.Close()

	_ = l.establish(l.id, &wsConn{c.UnderlyingConn(), c}, l.config.Auth)
}
