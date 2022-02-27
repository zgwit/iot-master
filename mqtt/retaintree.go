package mqtt

import (
	"strings"
	"sync"
)

type RetainNode struct {

	//Subscribed retains
	//clientId
	//retains map[string]*packet.Publish
	retains sync.Map

	//Sub level
	//topic->children
	//children map[string]*RetainNode
	children sync.Map
}

func (rn *RetainNode) Fetch(topics []string, cb func(clientId string, pub *Publish)) {
	if len(topics) == 0 {
		// Publish all matched retains
		rn.retains.Range(func(key, value interface{}) bool {
			cb(key.(string), value.(*Publish))
			return true
		})
	} else {
		name := topics[0]

		if name == "#" {
			//All retains
			rn.retains.Range(func(key, value interface{}) bool {
				cb(key.(string), value.(*Publish))
				return true
			})
			//And all children
			rn.children.Range(func(key, value interface{}) bool {
				value.(*RetainNode).Fetch(topics, cb)
				return true
			})
		} else if name == "+" {
			//Children
			rn.children.Range(func(key, value interface{}) bool {
				value.(*RetainNode).Fetch(topics[1:], cb)
				return true
			})
		} else {
			// Sub-Level
			if value, ok := rn.children.Load(name); ok {
				value.(*RetainNode).Fetch(topics[1:], cb)
			}
		}
	}
}

func (rn *RetainNode) Retain(topics []string, clientId string, pub *Publish) *RetainNode {
	if len(topics) == 0 {
		// Publish to specific client
		rn.retains.Store(clientId, pub)
		return rn
	} else {
		name := topics[0]

		// Sub-Level
		value, _ := rn.children.LoadOrStore(name, &RetainNode{})
		return value.(*RetainNode).Retain(topics[1:], clientId, pub)
	}
}

type RetainTree struct {
	//root
	root RetainNode

	//tree index
	//ClientId -> Node (hold Publish message)
	//retains map[string]*RetainNode
	retains sync.Map
}


func (rt *RetainTree) Fetch(topic []byte, cb func(clientId string, pub *Publish)) {
	topics := strings.Split(string(topic), "/")
	if topics[0] == "" {
		topics[0] = "/"
	}
	rt.root.Fetch(topics, cb)
}

func (rt *RetainTree) Retain(topic []byte, clientId string, pub *Publish) {
	// Remove last retain publish, firstly
	rt.UnRetain(clientId)

	topics := strings.Split(string(topic), "/")
	if topics[0] == "" {
		topics[0] = "/"
	}
	node := rt.root.Retain(topics, clientId, pub)

	//indexed node
	rt.retains.Store(clientId, node)
}

func (rt *RetainTree) UnRetain(clientId string) {
	if value, ok := rt.retains.Load(clientId); ok {
		node := value.(*RetainNode)
		node.retains.Delete(clientId)
		rt.retains.Delete(clientId)
	}
}
