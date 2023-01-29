package model

import "encoding/gob"

func init() {
	gob.Register(User{})
	gob.Register(Entrypoint{})
	gob.Register(Device{})
	gob.Register(Plugin{})
}
