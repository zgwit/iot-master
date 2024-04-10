package types

import (
	"github.com/zgwit/iot-master/v4/db"
)

func init() {
	db.Register(new(User), new(Password))
}
