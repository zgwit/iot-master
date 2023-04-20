package payload

type Event struct {
	Id      string         `json:"id"`
	Name    string         `json:"name"`
	Title   string         `json:"title"`
	Message string         `json:"message,omitempty"`
	Output  map[string]any `json:"output"`
}
