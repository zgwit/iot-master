package internal

import (
	paho "github.com/eclipse/paho.mqtt.golang"
	"github.com/zgwit/iot-master/v4/aggregator"
	"github.com/zgwit/iot-master/v4/broker"
	"github.com/zgwit/iot-master/v4/db"
	"github.com/zgwit/iot-master/v4/device"
	"github.com/zgwit/iot-master/v4/log"
	"github.com/zgwit/iot-master/v4/mqtt"
	"github.com/zgwit/iot-master/v4/plugin"
	"github.com/zgwit/iot-master/v4/pool"
	"github.com/zgwit/iot-master/v4/product"
	"github.com/zgwit/iot-master/v4/project"
	"github.com/zgwit/iot-master/v4/space"
	"github.com/zgwit/iot-master/v4/vconn"
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
	//err = db.Engine.Sync2(
	//	new(types.User), new(types.Password),
	//	new(types.Broker), new(types.Gateway),
	//	new(types.Product), new(types.Device),
	//	new(types.Plugin),
	//	new(types.Project), new(types.ProjectUser),
	//	new(types.ProjectPlugin), new(types.SpaceDevice),
	//	new(types.History), new(types.ExternalAggregator),
	//	new(alarm.Alarm), new(types.ExternalValidator),
	//	new(alarm.Subscription), new(alarm.Notification),
	//)
	//if err != nil {
	//	return err
	//}

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

	//加载所有插件
	err = plugin.Boot()
	if err != nil {
		return err
	}

	//加载产品库
	err = product.Boot()
	if err != nil {
		return err
	}

	//加载设备影子
	err = device.Boot()
	if err != nil {
		return err
	}

	//加载空间
	err = space.Boot()
	if err != nil {
		return err
	}

	//加载项目
	err = project.Boot()
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
