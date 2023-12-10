package app

import (
	"github.com/zgwit/iot-master/v4/mod"
	"net/http/httputil"
)

type App struct {
	*mod.Model

	runner *Runner
	proxy  *httputil.ReverseProxy
}
