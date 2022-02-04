package tsdb

import "github.com/nakabonne/tstorage"


func Test()  {
	store, _ := tstorage.NewStorage(
		tstorage.WithTimestampPrecision(tstorage.Seconds),
		tstorage.WithDataPath("history"),
		)
	defer store.Close()

}
