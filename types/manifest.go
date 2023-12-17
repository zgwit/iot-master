package types

type ManifestBase struct {
	Id          string   `json:"id,omitempty"`          //ID
	Version     string   `json:"version,omitempty"`     //版本 semver.Version
	Icon        string   `json:"icon,omitempty"`        //图标
	Name        string   `json:"name,omitempty"`        //名称
	Url         string   `json:"url,omitempty"`         //链接
	Description string   `json:"description,omitempty"` //说明
	Keywords    []string `json:"keywords,omitempty"`    //关键字
}
