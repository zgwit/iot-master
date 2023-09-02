package internal

import (
	paho "github.com/eclipse/paho.mqtt.golang"
	"github.com/zgwit/iot-master/v4/aggregator"
	"github.com/zgwit/iot-master/v4/db"
	"github.com/zgwit/iot-master/v4/internal/broker"
	"github.com/zgwit/iot-master/v4/internal/cfg"
	device2 "github.com/zgwit/iot-master/v4/internal/device"
	"github.com/zgwit/iot-master/v4/internal/product"
	"github.com/zgwit/iot-master/v4/log"
	"github.com/zgwit/iot-master/v4/mqtt"
	"github.com/zgwit/iot-master/v4/plugin"
	"github.com/zgwit/iot-master/v4/pool"
	"github.com/zgwit/iot-master/v4/types"
	"github.com/zgwit/iot-master/v4/vconn"
	"net"
	"net/url"
)

func Open() error {

	cfg.Load()

	err := log.Open()
	if err != nil {
		return err
	}

	//加载数据库
	err = db.Open()
	if err != nil {
		return err
	}

	//线程池
	err = pool.Open()
	if err != nil {
		return err
	}

	//同步表结构
	err = db.Engine.Sync2(
		new(types.User), new(types.Password), new(types.Role),
		new(types.Broker), new(types.Gateway),
		new(types.Product), new(types.Device),
		new(types.History), new(types.Aggregator),
		new(types.Alarm), new(types.Validator),
		new(types.Subscription), new(types.Notification),
		new(types.App), new(types.Plugin),
	)
	if err != nil {
		return err
	}

	//db.Engine.SetLogLevel(0)
	//db.Engine.ShowSQL(true)

	//启动计划任务
	aggregator.Start()

	err = broker.Open()
	if err != nil {
		return err
	}

	if broker.Server != nil {
		token := mqtt.OpenBy(
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
		token.Wait()
		err = token.Error()
		if err != nil {
			return err
		}
	} else {
		//MQTT总线
		token := mqtt.Open()
		token.Wait()
		err = token.Error()
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

	err = device2.SubscribeEvent()
	if err != nil {
		return err
	}

	err = device2.SubscribeProperty()
	if err != nil {
		return err
	}

	//err = device.SubscribePropertyStrict()
	//if err != nil {
	//	return err
	//}

	err = device2.SubscribeOnline()
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
