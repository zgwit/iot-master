package user

import (
	"encoding/gob"
)

func init() {
	gob.Register(User{})
}
