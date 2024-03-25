package tunnel

import "github.com/zgwit/iot-master/v4/pkg/log"

func Boot() error {
	err := LoadSerials()
	if err != nil {
		return err
	}
	err = LoadClients()
	if err != nil {
		return err
	}
	err = LoadServers()
	if err != nil {
		return err
	}
	return nil
}

func Close() {
	links.Range(func(name string, link *Link) bool {
		err := link.Close()
		if err != nil {
			log.Error(err)
		}
		return true
	})
}
