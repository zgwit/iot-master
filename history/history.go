package history

import "errors"

var (
	errorUnregister = errors.New("历史数据库未注册")
)

type Point struct {
	Value any   `json:"value"`
	Time  int64 `json:"time"`
}

type Historian interface {
	Write(table, id string, timestamp int64, values map[string]any) error

	Query(table, id, name, start, end, window, method string) ([]Point, error)
}

var historian Historian

func Register(h Historian) {
	historian = h
}

func Registered() bool {
	return historian != nil
}

func Write(table, id string, timestamp int64, values map[string]any) error {
	if historian == nil {
		//return errorUnregister
		return nil
	}
	return historian.Write(table, id, timestamp, values)
}

func Query(table, id, name, start, end, window, method string) ([]Point, error) {
	if historian == nil {
		return nil, errorUnregister
	}
	return historian.Query(table, id, name, start, end, window, method)
}
