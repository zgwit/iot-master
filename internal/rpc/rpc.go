package rpc

import (
	"encoding/json"
	"github.com/zgwit/iot-master/plugin"
	"google.golang.org/grpc"
	"net"
)

func jsonMarshalBuffer(value interface{}) (*plugin.Buffer, error) {
	buf, err := json.Marshal(value)
	if err != nil {
		return nil, err
	}

	return &plugin.Buffer{Value: buf}, nil
}

var server *grpc.Server

func Open() error {
	server = grpc.NewServer()
	plugin.RegisterModuleServer(server, &moduleServer{})
	plugin.RegisterDeviceServer(server, &deviceServer{})
	plugin.RegisterProjectServer(server, &projectServer{})
	plugin.RegisterProductServer(server, &productServer{})
	plugin.RegisterTunnelServer(server, &tunnelServer{})
	plugin.RegisterServerServer(server, &serverServer{})
	plugin.RegisterUserServer(server, &userServer{})

	tcp, err := net.Listen("tcp", ":1843")
	if err != nil {
		return err
	}
	unix, err := net.Listen("unix", "/iot-master.sock")
	if err != nil {
		return err
	}

	go server.Serve(tcp)
	go server.Serve(unix)

	return nil
}

func Close() {
	server.Stop()
	server = nil
}
