package core

import (
	"github.com/zgwit/iot-master/v3/model"
)

var GatewayStatus Map[model.Status] //status
var TunnelStatus Map[model.Status]  //status
var ServerStatus Map[model.Status]  //status

var Devices Map[Device]

var Services Map[model.Service]
