package config

import (
	"os"
	"path"
)

// RPC 参数
type RPC struct {
	Addr string
	Sock string
}

var RPCDefault = RPC{
	Addr: ":1843",
	Sock: path.Join(os.TempDir(), "iot-master-rpc.sock"),
}
