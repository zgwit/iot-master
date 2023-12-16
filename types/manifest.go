package types

import "github.com/blang/semver/v4"

type ManifestBase struct {
	Id          string         `json:"id,omitempty"`          //ID
	Version     semver.Version `json:"version,omitempty"`     //版本
	Icon        string         `json:"icon,omitempty"`        //图标
	Name        string         `json:"name,omitempty"`        //名称
	Url         string         `json:"url,omitempty"`         //链接
	Description string         `json:"description,omitempty"` //说明
	Keywords    []string       `json:"keywords,omitempty"`    //关键字
}
