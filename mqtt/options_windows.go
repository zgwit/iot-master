package mqtt

import (
	"github.com/zgwit/iot-master/v4/lib"
)

func Default() Options {
	return Options{
		Url:      "mqtt://localhost:1843",
		ClientId: lib.RandomString(8),
	}
}
