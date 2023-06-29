package internal

import (
	paho "github.com/eclipse/paho.mqtt.golang"
	"github.com/zgwit/iot-master/v3/internal/aggregator"
	"github.com/zgwit/iot-master/v3/internal/broker"
	"github.com/zgwit/iot-master/v3/internal/config"
	"github.com/zgwit/iot-master/v3/internal/device"
	"github.com/zgwit/iot-master/v3/internal/plugin"
	"github.com/zgwit/iot-master/v3/internal/product"
	"github.com/zgwit/iot-master/v3/model"
	"github.com/zgwit/iot-master/v3/pkg/db"
	"github.com/zgwit/iot-master/v3/pkg/log"
	"github.com/zgwit/iot-master/v3/pkg/mqtt"
	"github.com/zgwit/iot-master/v3/pkg/vconn"
	"net"
	"net/url"
)

func Open() error {

	config.Load()

	err := log.Open()
	if err != nil {
		return err
	}

	//加载数据库
	err = db.Open()
	if err != nil {
		return err
	}

	//同步表结构
	err = db.Engine.Sync2(
		new(model.User), new(model.Password), new(model.Role),
		new(model.Broker), new(model.Gateway),
		new(model.Product), new(model.Device),
		new(model.History), new(model.Aggregator),
		new(model.Alarm), new(model.Validator),
		new(model.Subscription), new(model.Notification),
		new(model.App), new(model.Plugin),
	)
	if err != nil {
		return err
	}

	//启动计划任务
	aggregator.Start()

	err = broker.Open()
	if err != nil {
		return err
	}

	if broker.Server != nil {
		err = mqtt.OpenBy(
			func(uri *url.URL, options paho.ClientOptions) (net.Conn, error) {
				c1, c2 := vconn.New()
				//EstablishConnection会读取connect，导致拥堵
				go func() {
					err := broker.Server.EstablishConnection("internal", c1)
					if err != nil {
						log.Error(err)
					}
				}()
				return c2, nil
			})
		if err != nil {
			return err
		}
	} else {
		//MQTT总线
		err = mqtt.Open()
		if err != nil {
			return err
		}
	}

	err = product.LoadAll()
	if err != nil {
		return err
	}

	go func() {
		err = plugin.LoadAll()
		if err != nil {
			log.Error(err)
		}
	}()

	//err = LoadDevices()
	//if err != nil {
	//	return err
	//}

	//webServe(fmt.Sprintf(":%d", config.Config.Web))
	//err = SubscribeMaster()
	//if err != nil {
	//	return err
	//}

	err = device.SubscribeEvent()
	if err != nil {
		return err
	}

	err = device.SubscribeProperty()
	if err != nil {
		return err
	}

	err = device.SubscribePropertyStrict()
	if err != nil {
		return err
	}

	err = device.SubscribeOnline()
	if err != nil {
		return err
	}

	return nil
}

func Close() {
	_ = db.Close()
	broker.Close()
	mqtt.Close()
	plugin.Close()
}
