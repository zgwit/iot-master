package mqtt

import (
	"encoding/json"
	paho "github.com/eclipse/paho.mqtt.golang"
	"github.com/zgwit/iot-master/v4/config"
	"github.com/zgwit/iot-master/v4/pkg/pool"
)

var Client paho.Client

func Close() {
	Client.Disconnect(0)
}

func Open() paho.Token {
	opts := paho.NewClientOptions()
	opts.AddBroker(config.GetString(MODULE, "url"))
	opts.SetClientID(config.GetString(MODULE, "clientId"))
	opts.SetUsername(config.GetString(MODULE, "username"))
	opts.SetPassword(config.GetString(MODULE, "password"))
	opts.SetConnectRetry(true) //重试

	opts.SetKeepAlive(20)

	//重连时，恢复订阅
	opts.SetCleanSession(false)
	opts.SetResumeSubs(true)

	//加上订阅处理
	opts.SetOnConnectHandler(func(client paho.Client) {
		//for topic, _ := range subs {
		//	Client.Subscribe(topic, 0, func(client paho.Client, message paho.Message) {
		//
		//		go func() {
		//			//依次处理回调
		//			if cbs, ok := subs[topic]; ok {
		//				for _, cb := range cbs {
		//					cb(message.Topic(), message.Payload())
		//				}
		//			}
		//		}()
		//	})
		//}
	})

	Client = paho.NewClient(opts)
	return Client.Connect()
}

func OpenBy(fn paho.OpenConnectionFunc) paho.Token {
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
	return Client.Connect()
}

func Publish(topic string, payload any) paho.Token {
	bytes, _ := json.Marshal(payload)
	return Client.Publish(topic, 0, false, bytes)
}

var subscribes = map[string]any{}

func Subscribe[T any](filter string, cb func(topic string, value *T)) {
	callbacks, ok := subscribes[filter]
	cbs := callbacks.([]func(topic string, value *T))
	if ok {
		subscribes[filter] = append(cbs, cb)
		return
	}

	//初次订阅
	Client.Subscribe(filter, 0, func(client paho.Client, message paho.Message) {
		subs := subscribes[filter]
		cs := subs.([]func(topic string, value *T))

		//解析JSON
		var value T
		if len(message.Payload()) > 0 {
			err := json.Unmarshal(message.Payload(), &value)
			if err != nil {
				return
			}
		}

		//回调
		for _, c := range cs {
			//放入线程池处理
			_ = pool.Insert(func() {
				c(message.Topic(), &value)
			})
		}
	})
}
