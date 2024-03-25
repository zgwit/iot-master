package broker

import (
	"github.com/zgwit/iot-master/v4/pkg/db"
	"time"
)

func init() {
	db.Register(new(Broker), new(Gateway))
}

type Broker struct {
	Id          string    `json:"id" xorm:"pk"`
	Name        string    `json:"name"`
	Description string    `json:"description,omitempty"`
	Port        int       `json:"port,omitempty"` //TODO 添加TLS证书
	Disabled    bool      `json:"disabled,omitempty"`
	Created     time.Time `json:"created,omitempty" xorm:"created"`
}
