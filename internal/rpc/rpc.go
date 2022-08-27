package rpc

import (
	"encoding/json"
	"github.com/zgwit/iot-master/internal/config"
	"github.com/zgwit/iot-master/plugin"
	"google.golang.org/grpc"
	"net"
	"os"
)

func jsonMarshalBuffer(value interface{}) (*plugin.Buffer, error) {
	buf, err := json.Marshal(value)
	if err != nil {
		return nil, err
	}

	return &plugin.Buffer{Value: buf}, nil
}

var server *grpc.Server

func Open(rpc config.RPC) error {
	server = grpc.NewServer()
	plugin.RegisterModuleServer(server, &moduleServer{})
	plugin.RegisterDeviceServer(server, &deviceServer{})
	plugin.RegisterProjectServer(server, &projectServer{})
	plugin.RegisterProductServer(server, &productServer{})
	plugin.RegisterTunnelServer(server, &tunnelServer{})
	plugin.RegisterServerServer(server, &serverServer{})
	plugin.RegisterUserServer(server, &userServer{})

	if rpc.Addr != "" {
		tcp, err := net.Listen("tcp", rpc.Addr)
		if err != nil {
			return err
		}
		go server.Serve(tcp)
	}

	if rpc.Sock != "" {
		_ = os.Remove(rpc.Sock)
		unix, err := net.Listen("unix", rpc.Sock)
		if err != nil {
			return err
		}
		go server.Serve(unix)
	}

	return nil
}

func Close() {
	server.Stop()
	server = nil
}
