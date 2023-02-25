package internal

import "github.com/zgwit/iot-master/v3/mqtt"

func Open() error {
	err := mqtt.Open()
	if err != nil {
		return err
	}

	err = subscribeProperty()
	if err != nil {
		return err
	}

	err = subscribeMaster()
	if err != nil {
		return err
	}

	//webServe(fmt.Sprintf(":%d", config.Config.Web))
	return nil
}

func Close() {
	//TODO clear gateways devices data

}
