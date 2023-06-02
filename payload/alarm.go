package payload

type Alarm struct {
	Product string `json:"product,omitempty"`
	Device  string `json:"device,omitempty"`
	Type    string `json:"type,omitempty"`
	Title   string `json:"title,omitempty"`
	Message string `json:"message,omitempty"`
	Level   uint   `json:"level,omitempty"`
}
