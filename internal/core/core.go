package core

import (
	"github.com/zgwit/iot-master/model"
	"sync"
)

var TunnelStatus Map[model.Status] //status
var ServerStatus Map[model.Status] //status

var Devices Map[*Device]
var Projects Map[*Project]

type Map[T any] struct {
	container map[string]T
	lock      sync.RWMutex
}

func (c *Map[T]) Load(name string) T {
	c.lock.RLock()
	defer c.lock.RUnlock()
	if c.container == nil {
		return nil
	}
	return c.container[name]
}

func (c *Map[T]) Store(name string, value T) {
	c.lock.Lock()
	defer c.lock.Unlock()
	if c.container == nil {
		c.container = make(map[string]T)
	}
	c.container[name] = value
}

func (c *Map[T]) Delete(name string) {
	c.lock.Lock()
	defer c.lock.Unlock()
	if c.container == nil {
		return
	}
	delete(c.container, name)
}
