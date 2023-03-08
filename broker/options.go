package broker

type Options struct {
	Enable    bool       `json:"enable"`
	Listeners []Listener `json:"listeners"`
}

type Listener struct {
	Type string `json:"type"`
	Addr string `json:"addr"`
}

func Default() Options {
	return Options{
		Enable: true,
		Listeners: []Listener{
			{Type: "tcp", Addr: ":1843"}, //后期可以关掉，可限制内部使用
			{Type: "unix", Addr: "iot-master.sock"},
		},
	}
}
