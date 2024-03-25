package device

type Adapter interface {
	Get(device *Device, point string) (any, error)
	Set(device *Device, point string, value any) error
	Sync(device *Device) (map[string]any, error)
}
