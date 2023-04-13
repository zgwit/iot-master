package mqtt

type Options struct {
	Url      string `json:"url,omitempty"`
	ClientId string `json:"clientId,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}
