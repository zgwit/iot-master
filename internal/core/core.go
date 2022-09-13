package core

import "sync"

var Tunnels sync.Map //status
var Servers sync.Map //status

var Devices sync.Map  //{values, status}
var Projects sync.Map //{values, }
