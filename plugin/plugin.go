package plugin

type Plugin struct {
	Id           string   `json:"id"`
	Name         string   `json:"name"`
	Version      string   `json:"version,omitempty"`
	EntryPoint   string   `json:"entrypoint,omitempty"`
	Dependencies []string `json:"dependencies,omitempty"`
}
