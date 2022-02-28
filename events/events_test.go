package events

import "testing"

func TestOnce(t *testing.T) {
	e := EventEmitter{}
	e.Once("test", func(arg ...interface{}) {
		t.Log(arg)
	})
	e.Emit("test", "publish1")
	e.Emit("test", "publish2")
}
