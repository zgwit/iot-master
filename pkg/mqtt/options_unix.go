//go:build linux || darwin || freebsd || openbsd
// +build linux darwin freebsd openbsd

package mqtt

import "github.com/zgwit/iot-master/v3/pkg/lib"

func Default() Options {
	return Options{
		Url:      "mqtt://iot-master.sock:1843",
		ClientId: lib.RandomString(8),
	}
}
