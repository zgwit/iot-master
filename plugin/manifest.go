package plugin

type Manifest struct {
	Id          string `json:"id"`                    //ID
	Version     string `json:"version,omitempty"`     //版本 semver.Version
	Icon        string `json:"icon,omitempty"`        //图标
	Name        string `json:"name,omitempty"`        //名称
	Description string `json:"description,omitempty"` //说明

	//菜单入口
	Menus map[string]*Menu `json:"menus,omitempty"` //admin, project

	//子页面
	Pages map[string]*Entry `json:"pages,omitempty"`

	//外部插件 进程
	Process *Process `json:"process,omitempty"`

	//更多
	Dependencies map[string]string `json:"dependencies,omitempty"` //应用和版本
}

type Entry struct {
	Name string `json:"name,omitempty"`
	Url  string `json:"url,omitempty"`
}

type Menu struct {
	Name  string   `json:"name"`
	Items []*Entry `json:"items"`
	First bool     `json:"first,omitempty"`
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
