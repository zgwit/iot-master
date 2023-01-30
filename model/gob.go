package model

import "encoding/gob"

func init() {
	gob.Register(User{})
	gob.Register(Server{})
	gob.Register(Device{})
	gob.Register(Plugin{})
}
