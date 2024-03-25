package tunnel

import (
	"github.com/zgwit/iot-master/v4/lib"
	"github.com/zgwit/iot-master/v4/pkg/db"
	"github.com/zgwit/iot-master/v4/pkg/log"
	"xorm.io/xorm"
)

var serials lib.Map[Serial]

var clients lib.Map[Client]

var servers lib.Map[Server]

var links lib.Map[Link]

func LoadSerials() error {
	var serials []*Serial
	err := db.Engine.Find(&serials)
	if err != nil {
		if err == xorm.ErrNotExist {
			return nil
		}
		return err
	}
	for _, m := range serials {
		if m.Disabled {
			continue
		}
		go func(m *Serial) {
			err := LoadSerial(m)
			if err != nil {
				log.Error(err)
			}
		}(m)
	}
	return nil
}

func LoadSerial(m *Serial) error {
	serials.Store(m.Id, m)
	return m.Open()
}

func GetSerial(id string) *Serial {
	return serials.Load(id)
}

func LoadClients() error {
	var clients []*Client
	err := db.Engine.Find(&clients)
	if err != nil {
		if err == xorm.ErrNotExist {
			return nil
		}
		return err
	}
	for _, m := range clients {
		if m.Disabled {
			continue
		}
		go func(m *Client) {
			err := LoadClient(m)
			if err != nil {
				log.Error(err)
			}
		}(m)
	}
	return nil
}

func LoadClient(m *Client) error {
	clients.Store(m.Id, m)
	return m.Open()
}

func GetClient(id string) *Client {
	return clients.Load(id)
}

func LoadServers() error {
	var servers []*Server
	err := db.Engine.Find(&servers)
	if err != nil {
		if err == xorm.ErrNotExist {
			return nil
		}
		return err
	}
	for _, m := range servers {
		if m.Disabled {
			continue
		}
		go func(m *Server) {
			err := LoadServer(m)
			if err != nil {
				log.Error(err)
			}
		}(m)
	}
	return nil
}

func LoadServer(m *Server) error {
	servers.Store(m.Id, m)
	return m.Open()
}

func GetServer(id string) *Server {
	return servers.Load(id)
}

func GetLink(id string) *Link {
	return links.Load(id)
}
