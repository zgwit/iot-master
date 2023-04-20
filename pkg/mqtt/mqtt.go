package mqtt

import (
	"encoding/json"
	paho "github.com/eclipse/paho.mqtt.golang"
)

var Client paho.Client

func Close() {
	Client.Disconnect(0)
}

func Open(cfg Options) error {
	opts := paho.NewClientOptions()
	opts.AddBroker(cfg.Url)
	opts.SetClientID(cfg.ClientId)
	opts.SetUsername(cfg.Username)
	opts.SetPassword(cfg.Password)
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

func Publish(topic string, payload []byte, retain bool, qos byte) error {
	//if broker.Server != nil {
	//	return broker.Server.Publish(topic, payload, retain, qos)
	//}
	//是否需要等待？
	token := Client.Publish(topic, qos, retain, payload)
	_ = token.Wait()
	return token.Error()
}

func PublishJson(topic string, payload any) error {
	bytes, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	return Publish(topic, bytes, false, 0)
}
