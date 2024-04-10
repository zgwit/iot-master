package notify

import (
	"github.com/zgwit/iot-master/v4/db"
)

func init() {
	db.Register(new(Subscription), new(Notification))
}
