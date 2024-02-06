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
	Pages map[string]*Page `json:"pages,omitempty"`

	//外部插件 进程
	Process *Process `json:"process,omitempty"`

	//更多
	Dependencies map[string]string `json:"dependencies,omitempty"` //应用和版本
}

type Menu struct {
	Name  string      `json:"name"`
	Items []*MenuItem `json:"items"`
}

type MenuItem struct {
	Name       string `json:"name,omitempty"`
	Url        string `json:"url,omitempty"`
	Standalone string `json:"standalone,omitempty"` //独立页面，弹窗显示
}

type Page struct {
	Name   string   `json:"name,omitempty"`
	Url    string   `json:"url,omitempty"`
	Select []string `json:"select,omitempty"` //页面选择器，比如：modbus, s7 ...
}

type Process struct {
	Main  string `json:"main"`
	Delay int    `json:"delay,omitempty"` //延迟启动 s
}
