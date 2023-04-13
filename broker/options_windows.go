package broker

func Default() Options {
	return Options{
		Enable: true,
		Type:   "tcp",
		Addr:   ":1843",
	}
}
