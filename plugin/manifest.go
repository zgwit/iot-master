package plugin

type Manifest struct {
	Id          string   `json:"id"`                    //ID
	Version     string   `json:"version,omitempty"`     //版本 semver.Version
	Icon        string   `json:"icon,omitempty"`        //图标
	Name        string   `json:"name,omitempty"`        //名称
	Url         string   `json:"url,omitempty"`         //链接
	Description string   `json:"description,omitempty"` //说明
	Keywords    []string `json:"keywords,omitempty"`    //关键字

	//前端入口
	Menus *[]Entry `json:"menus,omitempty"`
	Pages *Pages   `json:"pages,omitempty"` //子页面入口

	//外部插件
	Type string `json:"type,omitempty"` //类型：内部、应用、[静态页面]
	Main string `json:"main,omitempty"` //入口：程序文件

	//更多
	Dependencies map[string]string `json:"dependencies,omitempty"` //应用和版本

	//Os           []string          `json:"os,omitempty"`           //操作系统支持：linux windows darwin
	//Arch         []string          `json:"arch,omitempty"`         //CPU架构：x64 ia32 aarch64
	//Author       string            `json:"author,omitempty"`       //作者
	//Bugs         string            `json:"bugs,omitempty"`         //Bug
	//License      string            `json:"license,omitempty"`      //软件协议：GPL MIT Apache 。。。
}

type Entry struct {
	Name string `json:"name,omitempty"`
	Url  string `json:"url,omitempty"`
}

type PageEntry struct {
	Detail *Entry `json:"detail,omitempty"`
	Edit   *Entry `json:"edit,omitempty"`
}

type Pages struct {
	Product *PageEntry `json:"product,omitempty"`
	Device  *PageEntry `json:"device,omitempty"`
	Project *PageEntry `json:"project,omitempty"`
	Space   *PageEntry `json:"space,omitempty"`
}
