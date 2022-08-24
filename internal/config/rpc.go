package config

//RPC 参数
type RPC struct {
	Addr string
	Sock string
}

var RPCDefault = RPC{
	Addr: ":1843",
	Sock: "/iot-master-rpc.sock",
}
