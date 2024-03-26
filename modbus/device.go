package modbus

import (
	"github.com/zgwit/iot-master/v4/pkg/db"
)

func init() {
	db.Register(new(Device))
}

type Device struct {
	Id string `json:"id" xorm:"pk"`

	//modbus站号
	ModbusStation uint8 `json:"modbus_station,omitempty"`
}
