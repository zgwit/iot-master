package mqtt

import "github.com/zgwit/iot-master/v3/pkg/lib"

type Options struct {
	Url      string `json:"url,omitempty"`
	ClientId string `json:"clientId,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

func Default() Options {
	return Options{
		Url:      "mqtt://localhost:1843",
		ClientId: lib.RandomString(8),
	}
}
