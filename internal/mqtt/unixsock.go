package mqtt

import (
	"github.com/mochi-co/mqtt/server/listeners"
	"github.com/mochi-co/mqtt/server/listeners/auth"
	"github.com/mochi-co/mqtt/server/system"
	"net"
	"sync"
	"sync/atomic"
)

// UnixSock is a listener for establishing client connections on basic UnixSock protocol.
type UnixSock struct {
	sync.RWMutex
	id       string            // the internal id of the listener.
	protocol string            // the UnixSock protocol to use.
	address  string            // the network address to bind to.
	listen   net.Listener      // a net.Listener which will listen for new clients.
	config   *listeners.Config // configuration values for the listener.
	end      uint32            // ensure the close methods are only called once.
}

// NewUnixSock initialises and returns a new UnixSock listener, listening on an address.
func NewUnixSock(id, address string) *UnixSock {
	return &UnixSock{
		id:       id,
		protocol: "unix",
		address:  address,
		config: &listeners.Config{ // default configuration.
			Auth: new(auth.Allow),
		},
	}
}

// SetConfig sets the configuration values for the listener config.
func (l *UnixSock) SetConfig(config *listeners.Config) {
	l.Lock()
	if config != nil {
		l.config = config

		// If a config has been passed without an auth controller,
		// it may be a mistake, so disallow all traffic.
		if l.config.Auth == nil {
			l.config.Auth = new(auth.Disallow)
		}
	}

	l.Unlock()
}

// ID returns the id of the listener.
func (l *UnixSock) ID() string {
	l.RLock()
	id := l.id
	l.RUnlock()
	return id
}

// Listen starts listening on the listener's network address.
func (l *UnixSock) Listen(s *system.Info) error {
	var err error
	l.listen, err = net.Listen(l.protocol, l.address)
	return err
}

// Serve starts waiting for new UnixSock connections, and calls the establish
// connection callback for any received.
func (l *UnixSock) Serve(establish listeners.EstablishFunc) {
	for {
		if atomic.LoadUint32(&l.end) == 1 {
			return
		}

		conn, err := l.listen.Accept()
		if err != nil {
			return
		}

		if atomic.LoadUint32(&l.end) == 0 {
			go func() {
				_ = establish(l.id, conn, l.config.Auth)
			}()
		}
	}
}

// Close closes the listener and any client connections.
func (l *UnixSock) Close(closeClients listeners.CloseFunc) {
	l.Lock()
	defer l.Unlock()

	if atomic.CompareAndSwapUint32(&l.end, 0, 1) {
		closeClients(l.id)
	}

	if l.listen != nil {
		err := l.listen.Close()
		if err != nil {
			return
		}
	}
}
