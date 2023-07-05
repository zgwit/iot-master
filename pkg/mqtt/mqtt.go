package mqtt

import (
	"encoding/json"
	paho "github.com/eclipse/paho.mqtt.golang"
)

var Client paho.Client

func Close() {
	Client.Disconnect(0)
}

func Open() error {
	opts := paho.NewClientOptions()
	opts.AddBroker(options.Url)
	opts.SetClientID(options.ClientId)
	opts.SetUsername(options.Username)
	opts.SetPassword(options.Password)
	opts.SetConnectRetry(true) //重试

	//重连时，恢复订阅
	opts.SetCleanSession(false)
	opts.SetResumeSubs(true)

	Client = paho.NewClient(opts)
	token := Client.Connect()
	token.Wait()
	return token.Error()
}

func OpenBy(fn paho.OpenConnectionFunc) error {
	//client := Server.NewClient(nil, "internal", "internal", true)
	opts := paho.NewClientOptions()
	opts.AddBroker(":1883")
	opts.SetClientID("internal")

	//重连时，恢复订阅
	opts.SetCleanSession(false)
	opts.SetResumeSubs(true)

	//使用虚拟连接
	opts.SetCustomOpenConnectionFn(fn)

	Client = paho.NewClient(opts)
	token := Client.Connect()
	token.Wait()
	return token.Error()
}

func PublishRaw(topic string, payload []byte, retain bool, qos byte) error {
	//是否需要等待？
	token := Client.Publish(topic, qos, retain, payload)
	_ = token.Wait()
	return token.Error()
}

func Publish(topic string, payload any) paho.Token {
	bytes, _ := json.Marshal(payload)
	return Client.Publish(topic, 0, false, bytes)
}

type Handler func(topic string, payload []byte)

var subs = map[string][]Handler{}

func _subscribeHandlers(client paho.Client, message paho.Message) {
	//依次处理回调
	if cbs, ok := subs[message.Topic()]; ok {
		for _, cb := range cbs {
			cb(message.Topic(), message.Payload())
		}
	}
}

func Subscribe(topic string, cb Handler) paho.Token {
	if cbs, ok := subs[topic]; !ok {
		subs[topic] = []Handler{cb}
	} else {
		subs[topic] = append(cbs, cb)
	}
	return Client.Subscribe(topic, 0, _subscribeHandlers)
}

func SubscribeJson(topic string, cb func(topic string, data map[string]any)) paho.Token {
	return Subscribe(topic, func(topic string, payload []byte) {
		var data map[string]any
		err := json.Unmarshal(payload, &data)
		if err != nil {
			return
		}
		cb(topic, data)
	})
}

func SubscribeStruct[T any](topic string, cb func(topic string, data *T)) paho.Token {
	return Subscribe(topic, func(topic string, payload []byte) {
		var data T
		err := json.Unmarshal(payload, &data)
		if err != nil {
			return
		}
		cb(topic, &data)
	})
}
