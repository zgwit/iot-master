package mod

type App struct {
	Id           string            `json:"id" xorm:"pk"`           //ID
	Icon         string            `json:"icon,omitempty"`         //图标
	Name         string            `json:"name"`                   //名称
	Version      string            `json:"version"`                //SEMVER
	Type         string            `json:"type"`                   //类型：服务、应用、静态页面
	Main         string            `json:"main,omitempty"`         //入口：程序文件
	Keywords     []string          `json:"keywords,omitempty"`     //关键字
	Description  string            `json:"description,omitempty"`  //说明
	Dependencies map[string]string `json:"dependencies,omitempty"` //应用和版本
	Homepage     string            `json:"homepage"`               //主页
	Bugs         string            `json:"bugs,omitempty"`         //Bug
	License      string            `json:"license,omitempty"`      //软件协议
	Author       string            `json:"author,omitempty"`       //作者
	Os           []string          `json:"os,omitempty"`           //操作系统支持：linux windows darwin
	Arch         []string          `json:"arch,omitempty"`         //CPU架构：x64 ia32 aarch64
}
