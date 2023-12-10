package app

import (
	"github.com/blang/semver/v4"
	"github.com/zgwit/iot-master/v4/types"
)

type License struct {
	AppId    string         `json:"app_id"`
	Issuer   string         `json:"issuer,omitempty"`
	User     string         `json:"user,omitempty"`
	Machine  string         `json:"machine"`
	Version  semver.Version `json:"version,omitempty"` //版本
	ExpireAt types.Time     `json:"expire_at"`
	Content  string         `json:"content"`
}
