package protocol

type Adapter interface {
	Mount(device string) error
	Unmount(device string) error

	Get(device, point string) (any, error)
	Set(device, point string, value any) error
	Sync(device string) (map[string]any, error)
}
