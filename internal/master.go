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

func RemoveDevice(id int) error {
	d, ok := devices.LoadAndDelete(id)
	if ok {
		dev := d.(*Device)
		return dev.Stop()
	}
	return nil //error
}


func GetProject(id int) *Project  {
	d, ok := projects.Load(id)
	if ok {
		return d.(*Project)
	}
	return nil
}

func RemoveProject(id int) error {
	d, ok := projects.LoadAndDelete(id)
	if ok {
		dev := d.(*Project)
		return dev.Stop()
	}
	return nil //error
}

