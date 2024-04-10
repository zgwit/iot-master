package internal

import (
	paho "github.com/eclipse/paho.mqtt.golang"
	"github.com/zgwit/iot-master/v4/broker"
	"github.com/zgwit/iot-master/v4/log"
	"github.com/zgwit/iot-master/v4/mqtt"
	"github.com/zgwit/iot-master/v4/pkg/vconn"
	"net"
	"net/url"
)

func Open() error {

	//内部桥接mqtt
	mqtt.OpenBy(
		func(uri *url.URL, options paho.ClientOptions) (net.Conn, error) {
			c1, c2 := vconn.New()
			//EstablishConnection会读取connect，导致拥堵
			go func() {
				//broker.Server.Clients.GetByListener("internal")
				err := broker.Server.EstablishConnection("internal", c1)
				if err != nil {
					log.Error(err)
				}
			}()
			return c2, nil
		})

	return nil
}
