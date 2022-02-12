package events

import "github.com/asaskevich/EventBus"

type EventEmitterInterface interface {
	Emit(event string, data ...interface{})
	On(event string, fn interface{})
	Once(event string, fn interface{})
	Off(event string, fn interface{})
}

type EventEmitter struct {
	events EventBus.Bus
}

func (e *EventEmitter) Emit(event string, data ...interface{}) {
	if e.events == nil {
		return
	}
	e.events.Publish(event, data...)
}

func (e *EventEmitter) On(event string, fn interface{}) {
	if e.events == nil {
		e.events = EventBus.New()
	}
	_ = e.events.Subscribe(event, fn)
}

func (e *EventEmitter) Once(event string, fn interface{}) {
	if e.events == nil {
		e.events = EventBus.New()
	}
	_ = e.events.SubscribeOnce(event, fn)
}

func (e *EventEmitter) Off(event string, fn interface{}) {
	if e.events == nil {
		return
	}
	_ = e.events.Unsubscribe(event, fn)
}
