package core

import "sync"

var TunnelStatus sync.Map //status
var ServerStatus sync.Map //status

//var DeviceStatus sync.Map  //{values, status}
//var ProjectStatus sync.Map //{values, }

var Devices sync.Map //map[]
var Projects sync.Map
