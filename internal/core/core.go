package core

import (
	"github.com/zgwit/iot-master/v2/model"
)

var GatewayStatus Map[model.Status] //status
var TunnelStatus Map[model.Status]  //status
var ServerStatus Map[model.Status]  //status

var Devices Map[Device]
var Projects Map[Project]
