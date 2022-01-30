package types

import "time"

type Checker interface {
	Reset()
	Check(now *time.Time) bool
}

type Delay struct {
	Delay int

	start time.Time
}

func (d *Delay) Reset() {

}

func (d *Delay) Check(now *time.Time) bool {

	return false
}

type Repeater struct {
	Interval int
	Total    int

	last time.Time

	raised     bool
	resetTimes int
}


func (d *Repeater) Reset() {

}

func (d *Repeater) Check(now *time.Time) bool {

	return false
}
