package core

import (
	paho "github.com/eclipse/paho.mqtt.golang"
	"github.com/zgwit/iot-master/v4/broker"
	"github.com/zgwit/iot-master/v4/pkg/aggregator"
	"github.com/zgwit/iot-master/v4/pkg/db"
	"github.com/zgwit/iot-master/v4/pkg/log"
	"github.com/zgwit/iot-master/v4/pkg/mqtt"
	"github.com/zgwit/iot-master/v4/pkg/pool"
	"github.com/zgwit/iot-master/v4/pkg/vconn"
	"github.com/zgwit/iot-master/v4/plugin"
	"github.com/zgwit/iot-master/v4/product"
	"github.com/zgwit/iot-master/v4/project"
	"github.com/zgwit/iot-master/v4/types"
	"net"
	"net/url"
)

func Open() error {

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
		new(types.User), new(types.Password),
		new(types.Broker), new(types.Gateway),
		new(types.Product), new(types.Device),
		new(types.Plugin),
		new(types.Project), new(types.ProjectUser),
		new(types.ProjectPlugin), new(types.ProjectDevice),
		new(types.History), new(types.ExternalAggregator),
		new(types.Alarm), new(types.ExternalValidator),
		new(types.Subscription), new(types.Notification),
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

	go func() {
		err = plugin.LoadAll()
		if err != nil {
			log.Error(err)
		}
	}()

	err = product.LoadAll()
	if err != nil {
		return err
	}

	err = project.LoadAll()
	if err != nil {
		return err
	}

	//err = LoadDevices()
	//if err != nil {
	//	return err
	//}

	//webServe(fmt.Sprintf(":%d", config.Config.Web))
	//err = SubscribeMaster()
	//if err != nil {
	//	return err
	//}

	//err = device.SubscribeEvent()
	//if err != nil {
	//	return err
	//}
	//
	//err = device.SubscribeProperty()
	//if err != nil {
	//	return err
	//}
	//
	////err = device.SubscribePropertyStrict()
	////if err != nil {
	////	return err
	////}
	//
	//err = device.SubscribeOnline()
	//if err != nil {
	//	return err
	//}

	return nil
}

func Close() {
	_ = db.Close()
	broker.Close()
	mqtt.Close()
	plugin.Close()
}
