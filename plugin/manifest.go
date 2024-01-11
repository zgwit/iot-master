package plugin

type Manifest struct {
	Type         string            `json:"type"`                   //类型：应用、外部、静态页面
	Main         string            `json:"main,omitempty"`         //入口：程序文件
	Pages        *Pages            `json:"pages,omitempty"`        //子页面入口
	Dependencies map[string]string `json:"dependencies,omitempty"` //应用和版本
	Os           []string          `json:"os,omitempty"`           //操作系统支持：linux windows darwin
	Arch         []string          `json:"arch,omitempty"`         //CPU架构：x64 ia32 aarch64
	Author       string            `json:"author,omitempty"`       //作者
	Bugs         string            `json:"bugs,omitempty"`         //Bug
	License      string            `json:"license,omitempty"`      //软件协议：GPL MIT Apache 。。。
}

type PageEntry struct {
	Detail string `json:"detail,omitempty"`
	Edit   string `json:"edit,omitempty"`
}

type Pages struct {
	Project *PageEntry `json:"project,omitempty"`
	Product *PageEntry `json:"product,omitempty"`
	Device  *PageEntry `json:"device,omitempty"`
}
