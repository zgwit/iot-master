package plugin

type Manifest struct {
	Id          string   `json:"id"`                    //ID
	Version     string   `json:"version,omitempty"`     //版本 semver.Version
	Icon        string   `json:"icon,omitempty"`        //图标
	Name        string   `json:"name,omitempty"`        //名称
	Url         string   `json:"url,omitempty"`         //链接
	Description string   `json:"description,omitempty"` //说明
	Keywords    []string `json:"keywords,omitempty"`    //关键字

	//菜单入口
	Menu *Menu `json:"menu,omitempty"`

	//子页面
	Pages *Pages `json:"pages,omitempty"`

	Process *Process `json:"process,omitempty"`

	//Pages *Pages   `json:"pages,omitempty"` //子页面入口

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

type Menu struct {
	Name  string  `json:"name"`
	Items []Entry `json:"items"`
	First bool    `json:"first,omitempty"`
}

type Pages struct {
	ProductEdit   *Entry `json:"product_edit,omitempty"`
	ProductDetail *Entry `json:"product_detail,omitempty"`
	DeviceEdit    *Entry `json:"device_edit,omitempty"`
	DeviceDetail  *Entry `json:"device_detail,omitempty"`
	ProjectEdit   *Entry `json:"project_edit,omitempty"`
	ProjectDetail *Entry `json:"project_detail,omitempty"`
	SpaceEdit     *Entry `json:"space_edit,omitempty"`
	SpaceDetail   *Entry `json:"space_detail,omitempty"`
}

type Process struct {
	Main  string `json:"main"`
	Delay int    `json:"delay,omitempty"` //延迟启动 s
}

type PageEntry struct {
	Detail *Entry `json:"detail,omitempty"`
	Edit   *Entry `json:"edit,omitempty"`
}
