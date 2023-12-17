package plugin

import (
	"github.com/zgwit/iot-master/v4/types"
)

type Manifest struct {
	types.ManifestBase

	Type         string            `json:"type"`                   //类型：服务、应用、静态页面
	Main         string            `json:"main,omitempty"`         //入口：程序文件
	Dependencies map[string]string `json:"dependencies,omitempty"` //应用和版本
	Os           []string          `json:"os,omitempty"`           //操作系统支持：linux windows darwin
	Arch         []string          `json:"arch,omitempty"`         //CPU架构：x64 ia32 aarch64
	Author       string            `json:"author,omitempty"`       //作者
	Bugs         string            `json:"bugs,omitempty"`         //Bug
	License      string            `json:"license,omitempty"`      //软件协议：GPL MIT Apache 。。。

	Entries []*Entry `json:"entries" xorm:"json"`
}

type Entry struct {
	Name        string `json:"name"`
	Icon        string `json:"icon,omitempty"`
	Path        string `json:"path"`
	Description string `json:"description,omitempty"`
}
