package events

import (
	"reflect"
	"sync"
)

type Handler func(args ...interface{})

//EventInterface Events接口
type EventInterface interface {
	Emit(event string, data ...interface{})
	On(event string, fn interface{})
	Once(event string, fn interface{})
	Off(event string, fn interface{})
}

type EventEmitter struct {
	handlers sync.Map
}

type eventHandler struct {
	callback reflect.Value
	once     bool
}

//Emit 发送消息
func (e *EventEmitter) Emit(event string, data ...interface{}) {
	val, ok := e.handlers.Load(event)
	if !ok {
		return
	}
	handlers := val.(*sync.Map)
	args := make([]reflect.Value, 0)
	for _, v := range data {
		args = append(args, reflect.ValueOf(v))
	}
	handlers.Range(func(key, value interface{}) bool {
		handler := value.(*eventHandler)
		handler.callback.Call(args)
		//处理仅订阅一次
		if handler.once {
			handlers.Delete(key)
		}
		return true
	})
}

//On 监听
func (e *EventEmitter) On(event string, fn interface{}) {
	callback := reflect.ValueOf(fn)
	val, ok := e.handlers.Load(event)
	if !ok {
		val = new(sync.Map)
		e.handlers.Store(event, val)
	}
	handlers := val.(*sync.Map)
	handlers.Store(callback.Pointer(), &eventHandler{
		callback: callback,
		once:     false,
	})
}

//Once 监听一次
func (e *EventEmitter) Once(event string, fn interface{}) {
	callback := reflect.ValueOf(fn)
	val, ok := e.handlers.Load(event)
	if !ok {
		val = new(sync.Map)
		e.handlers.Store(event, val)
	}
	handlers := val.(*sync.Map)
	handlers.Store(callback.Pointer(), &eventHandler{
		callback: callback,
		once:     true,
	})
}

//Off 取消监听
func (e *EventEmitter) Off(event string, fn interface{}) {
	callback := reflect.ValueOf(fn)
	val, ok := e.handlers.Load(event)
	if !ok {
		return
	}
	handlers := val.(*sync.Map)
	handlers.Delete(callback.Pointer())
}
