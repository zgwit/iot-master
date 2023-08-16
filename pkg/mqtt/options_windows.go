package mqtt

import "github.com/zgwit/iot-master/v3/pkg/lib"

func Default() Options {
	return Options{
		Url:      "mqtt://localhost:1843",
		ClientId: lib.RandomString(8),
	}
}
