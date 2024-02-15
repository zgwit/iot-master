package alarm

import (
	"github.com/zgwit/iot-master/v4/pkg/db"
)

func init() {
	db.Register(new(Alarm), new(Subscription), new(Notification))
}
