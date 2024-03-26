package protocol

type Adapter interface {
	Get(device, point string) (any, error)
	Set(device, point string, value any) error
	Sync(device string) (map[string]any, error)
}
