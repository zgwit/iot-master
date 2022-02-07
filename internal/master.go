package interval

import "sync"

var devices sync.Map
var projects sync.Map

func GetDevice(id int) *Device  {
	d, ok := devices.Load(id)
	if ok {
		return d.(*Device)
	}
	return nil
}

