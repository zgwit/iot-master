package internal

import "github.com/robfig/cron/v3"

func Open() error {

	err := LoadProducts()
	if err != nil {
		return err
	}

	//err = LoadDevices()
	//if err != nil {
	//	return err
	//}

	//webServe(fmt.Sprintf(":%d", config.Config.Web))
	err = subscribeMaster()
	if err != nil {
		return err
	}

	err = subscribeEvent()
	if err != nil {
		return err
	}

	err = subscribeProperty()
	if err != nil {
		return err
	}

	err = subscribePropertyStrict()
	if err != nil {
		return err
	}

	c := cron.New()
	_, err = c.AddFunc("@hourly", store)
	if err != nil {
		return err
	}

	return nil
}

func Close() {
	//TODO clear gateways devices data

}
