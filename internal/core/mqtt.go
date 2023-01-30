package core

import (
	"fmt"
	"github.com/mochi-co/mqtt/v2"
	"github.com/mochi-co/mqtt/v2/hooks/auth"
	"github.com/mochi-co/mqtt/v2/listeners"

	//mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/zgwit/iot-master/v3/internal/db"
	"github.com/zgwit/iot-master/v3/model"
	"github.com/zgwit/iot-master/v3/pkg/log"
	"xorm.io/xorm"
)

var mqttServer *mqtt.Server

func Open() error {

	//创建内部Broker
	mqttServer = mqtt.New(nil)

	//TODO 鉴权
	_ = mqttServer.AddHook(new(auth.AllowHook), nil)

	err := loadListeners()
	if err != nil {
		return err
	}

	//TODO websocket

	err = mqttServer.Serve()
	if err != nil {
		return err
	}

	return nil
}

func loadListeners() error {
	//监听服务
	//加载数据库中 entrypoint
	var entries []model.Entrypoint
	err := db.Engine.Find(&entries)
	if err != nil && err != xorm.ErrNotExist {
		return err
	}

	for _, e := range entries {
		l := listeners.NewTCP("tcp", fmt.Sprintf(":%d", e.Port), nil)
		err = mqttServer.AddListener(l)
		if err != nil {
			//return err
			log.Error(err)
		}
	}

	return nil
}

func Close() {
	if mqttServer != nil {
		_ = mqttServer.Close()
	}
}

func Publish(topic string, payload []byte) error {
	//TODO 兼容struct类型
	return mqttServer.Publish(topic, payload, false, 0)
}
