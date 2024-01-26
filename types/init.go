package types

import (
	"github.com/zgwit/iot-master/v4/db"
)

func init() {
	db.Register(new(User), new(Password),
		new(Broker), new(Gateway),
		new(Product), new(Device),
		new(Project), new(ProjectUser),
		new(Space), new(SpaceDevice),
		new(History))
}
