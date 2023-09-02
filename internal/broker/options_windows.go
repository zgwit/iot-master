package broker

func Default() Options {
	return Options{
		Enable: true,
		Addr:   ":1843",
	}
}
