package events

import "testing"

func TestOnce(t *testing.T) {
	e := EventEmitter{}
	e.On("test", func(arg ...interface{}) {
		t.Log("On", arg)
	})
	e.Once("test", func(arg ...interface{}) {
		t.Log("Once", arg)
	})
	e.Emit("test", "msg1")
	e.Emit("test", "msg2")
}
