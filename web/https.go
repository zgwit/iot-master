package web

import (
	"github.com/gin-gonic/autotls"
	"github.com/zgwit/iot-master/v4/pkg/config"
	"github.com/zgwit/iot-master/v4/pkg/log"
	"golang.org/x/crypto/acme/autocert"
)

func ServeTLS() error {
	cert := config.GetString(MODULE, "cert")
	key := config.GetString(MODULE, "key")

	log.Info("Web server tls", cert, key)
	return Engine.RunTLS(":443", cert, key)
}

func ServeLetsEncrypt() error {
	hosts := config.GetStringSlice(MODULE, "hosts")
	log.Info("Web server with LetsEncrypt", hosts)

	//初始化autocert
	manager := &autocert.Manager{
		Cache:      autocert.DirCache("certs"),
		Email:      config.GetString(MODULE, "email"),
		HostPolicy: autocert.HostWhitelist(hosts...),
		Prompt:     autocert.AcceptTOS,
	}

	return autotls.RunWithManager(Engine, manager)
}
