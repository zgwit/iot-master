package app

import (
	"github.com/blang/semver/v4"
	"time"
)

type Model struct {
	Id       string    `json:"id" xorm:"pk"`
	Name     string    `json:"name"`
	Version  string    `json:"version"`
	Command  string    `json:"command,omitempty"`
	Running  bool      `json:"running,omitempty" xorm:"-"`
	Disabled bool      `json:"disabled,omitempty"`
	Created  time.Time `json:"created,omitempty" xorm:"created"`
}

type Manifest struct {
	Id           string            `json:"id"`                     //ID
	Icon         string            `json:"icon,omitempty"`         //图标
	Name         string            `json:"name"`                   //名称
	Version      semver.Version    `json:"version,omitempty"`      //SEMVER
	Type         string            `json:"type"`                   //类型：服务、应用、静态页面
	Main         string            `json:"main,omitempty"`         //入口：程序文件
	Keywords     []string          `json:"keywords,omitempty"`     //关键字
	Description  string            `json:"description,omitempty"`  //说明
	Dependencies map[string]string `json:"dependencies,omitempty"` //应用和版本
	Os           []string          `json:"os,omitempty"`           //操作系统支持：linux windows darwin
	Arch         []string          `json:"arch,omitempty"`         //CPU架构：x64 ia32 aarch64
	Author       string            `json:"author,omitempty"`       //作者
	Url          string            `json:"url,omitempty"`          //主页
	Bugs         string            `json:"bugs,omitempty"`         //Bug
	License      string            `json:"license,omitempty"`      //软件协议：GPL MIT Apache 。。。
}
