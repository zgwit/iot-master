package internal

import (
	paho "github.com/eclipse/paho.mqtt.golang"
	"github.com/zgwit/iot-master/v4/aggregator"
	"github.com/zgwit/iot-master/v4/broker"
	"github.com/zgwit/iot-master/v4/device"
	"github.com/zgwit/iot-master/v4/pkg/db"
	"github.com/zgwit/iot-master/v4/pkg/log"
	"github.com/zgwit/iot-master/v4/pkg/mqtt"
	"github.com/zgwit/iot-master/v4/pkg/pool"
	"github.com/zgwit/iot-master/v4/pkg/vconn"
	"github.com/zgwit/iot-master/v4/plugin"
	"github.com/zgwit/iot-master/v4/product"
	"github.com/zgwit/iot-master/v4/project"
	"github.com/zgwit/iot-master/v4/space"
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
	err = device.Open()
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
