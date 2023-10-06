package mod

import "github.com/blang/semver/v4"

type License struct {
	AppId    string         `json:"app_id"`
	User     string         `json:"user,omitempty"`
	Machine  string         `json:"machine"`
	Version  semver.Version `json:"version,omitempty"` //版本
	ExpireAt Time           `json:"expire_at"`
	Content  string         `json:"content"`
}
