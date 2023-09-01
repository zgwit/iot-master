//go:build linux || darwin || freebsd || openbsd
// +build linux darwin freebsd openbsd

package mqtt

import "github.com/zgwit/iot-master/v4/pkg/lib"

func Default() Options {
	return Options{
		Url:      "unix://iot-master.sock",
		ClientId: lib.RandomString(8),
	}
}
