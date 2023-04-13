//go:build linux || darwin || freebsd || openbsd
// +build linux darwin freebsd openbsd

package mqtt

func Default() Options {
	return Options{
		Url:      "mqtt://iot-master.sock:1843",
		ClientId: lib.RandomString(8),
	}
}
