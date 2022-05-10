package omron

import "github.com/zgwit/iot-master/protocol"

var Codes = []protocol.Code{
	{"D", "DM Area"},
	{"C", "CIO Area"},
	{"W", "Work Area"},
	{"H", "Holding Bit Area"},
	{"A", "Auxiliary Bit Area"},
}

var DescHostlink = protocol.Protocol{
	Name:    "OmronHostlink",
	Version: "1.0",
	Label:   "Omron Hostlink",
	Codes:   Codes,
	//Factory: newHostlink,
}

var DescUDP = protocol.Protocol{
	Name:    "OmronFinsUDP",
	Version: "1.0",
	Label:   "Omron FINS UDP",
	Codes:   Codes,
	//Factory: newUDP,
}

var DescTCP = protocol.Protocol{
	Name:    "OmronFinsTCP",
	Version: "1.0",
	Label:   "Omron FINS TCP",
	Codes:   Codes,
	//Factory: newTCP,
}
