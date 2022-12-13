package model

import "encoding/gob"

func init() {
	gob.Register(User{})
	gob.Register(Tunnel{})
	gob.Register(Server{})
	gob.Register(Device{})
	gob.Register(Product{})
	gob.Register(Plugin{})
}
