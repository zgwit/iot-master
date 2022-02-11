package main

import (
	"github.com/zgwit/iot-master/internal"
	"github.com/zgwit/iot-master/web"
)

func main() {

	_ = internal.Start()

	web.Serve()
}