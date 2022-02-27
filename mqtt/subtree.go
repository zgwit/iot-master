package mqtt

import (
	"strings"
	"sync"
)

type SubNode struct {
	//Subscribed clients
	//clientId
	//clients map[string]packet.MsgQos
	clients sync.Map

	//Sub level
	//topic->children
	//children map[string]*SubNode
	children sync.Map

	//Multi Wildcard #
	mw *SubNode

	//Single Wildcard +
	sw *SubNode
}

func (sn *SubNode) Publish(topics []string, subs map[string]MsgQos) {
	if len(topics) == 0 {
		// Publish all matched clients
		sn.clients.Range(func(key, value interface{}) bool {
			clientId := key.(string)
			qos := value.(MsgQos)

			if sub, ok := subs[clientId]; ok {
				//rewrite by larger Qos
				if sub < qos {
					subs[clientId] = qos
				}
			} else {
				subs[clientId] = qos
			}

			return true
		})
	} else {
		name := topics[0]
		// Sub-Level

		if sub, ok := sn.children.Load(name); ok {
			sub.(*SubNode).Publish(topics[1:], subs)
		}
		// Multi wildcard
		if sn.mw != nil {
			sn.mw.Publish(topics[1:1], subs)
		}
		// Single wildcard
		if sn.sw != nil {
			sn.sw.Publish(topics[1:], subs)
		}
	}
}

func (sn *SubNode) Subscribe(topics []string, clientId string, qos MsgQos) {
	if len(topics) == 0 {
		sn.clients.Store(clientId, qos)
		return
	}

	name := topics[0]
	if name == "#" {
		if sn.mw == nil {
			sn.mw = &SubNode{}
		}
		sn.mw.Subscribe(topics[1:1], clientId, qos)
	} else if name == "+" {
		if sn.sw == nil {
			sn.sw = &SubNode{}
		}
		sn.sw.Subscribe(topics[1:], clientId, qos)
	} else {
		value, _ := sn.children.LoadOrStore(name, &SubNode{})
		value.(*SubNode).Subscribe(topics[1:], clientId, qos)
	}
}

func (sn *SubNode) UnSubscribe(topics []string, clientId string) {
	if len(topics) == 0 {
		sn.clients.Delete(clientId)
	} else {
		name := topics[0]
		if name == "#" {
			if sn.mw != nil {
				sn.mw.UnSubscribe(topics[1:1], clientId)
			}
		} else if name == "+" {
			if sn.sw != nil {
				sn.sw.UnSubscribe(topics[1:], clientId)
			}
		} else {
			sn.children.Range(func(key, value interface{}) bool {
				value.(*SubNode).UnSubscribe(topics[1:], clientId)
				return true
			})
		}
	}
}

func (sn *SubNode) ClearClient(clientId string) {
	sn.clients.Delete(clientId)

	if sn.mw != nil {
		sn.mw.ClearClient(clientId)
	}
	if sn.sw != nil {
		sn.sw.ClearClient(clientId)
	}

	sn.children.Range(func(key, value interface{}) bool {
		value.(*SubNode).ClearClient(clientId)
		return true
	})
}

type SubTree struct {
	//tree root
	root SubNode
}

func (st *SubTree) Publish(topic []byte, subs map[string]MsgQos) {
	topics := strings.Split(string(topic), "/")
	if topics[0] == "" {
		topics[0] = "/"
	}
	st.root.Publish(topics, subs)
}

func (st *SubTree) Subscribe(topic []byte, clientId string, qos MsgQos) {
	topics := strings.Split(string(topic), "/")
	if topics[0] == "" {
		topics[0] = "/"
	}
	st.root.Subscribe(topics, clientId, qos)
}

func (st *SubTree) UnSubscribe(topic []byte, clientId string) {
	topics := strings.Split(string(topic), "/")
	if topics[0] == "" {
		topics[0] = "/"
	}
	st.root.UnSubscribe(topics, clientId)
}

func (st *SubTree) ClearClient(clientId string) {
	st.root.ClearClient(clientId)
}
