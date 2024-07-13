package aggregate

import (
	"github.com/god-jason/bucket/table"
	"github.com/zgwit/iot-master/v5/base"
)

var _table = table.Table{
	Name: base.BucketAggregate,
	Fields: []*table.Field{
		base.ProductIdField,
		base.DeviceIdField,
		base.ProjectIdField,
		base.SpaceIdField,
		{Name: "date", Label: "日期", Type: "date"},
		//{Name: "created", Label: "创建日期", Type: "date"},
	},
}

func init() {
	table.Register(&_table)
}

func Table() *table.Table {
	return &_table
}
