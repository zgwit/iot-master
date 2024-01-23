package web

import (
	"github.com/zgwit/iot-master/v4/config"
	"github.com/zgwit/iot-master/v4/log"
	"golang.org/x/crypto/acme/autocert"
	"net/http"
)

func ServeTLS() {
	cert := config.GetString(MODULE, "cert")
	key := config.GetString(MODULE, "key")

	log.Info("Web ServeTLS", cert, key)
	err := Engine.RunTLS(":443", cert, key)
	if err != nil {
		log.Fatal(err)
	}
}

func ServeLetsEncrypt() {
	hosts := config.GetStringSlice(MODULE, "hosts")
	log.Info("Web ServeLetsEncrypt", hosts)

	//初始化autocert
	manager := &autocert.Manager{
		Cache:      autocert.DirCache("certs"),
		Email:      config.GetString(MODULE, "email"),
		HostPolicy: autocert.HostWhitelist(hosts...),
		Prompt:     autocert.AcceptTOS,
	}

	//创建server
	svr := &http.Server{
		Addr:      ":443",
		TLSConfig: manager.TLSConfig(),
		Handler:   Engine,
	}

	//监听https
	err := svr.ListenAndServeTLS("", "")
	if err != nil {
		log.Fatal(err)
	}
}
