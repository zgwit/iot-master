//go:build linux || darwin || freebsd || openbsd
// +build linux darwin freebsd openbsd

package broker

func Default() Options {
	return Options{
		Enable: true,
		Addr:   ":1843",
		Unix:   true,
	}
}
