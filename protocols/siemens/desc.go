package siemens

import (
	"iot-master/connect"
	"iot-master/protocols/protocol"
)

var Codes = []protocol.Code{
	{"I", "I 输入"},
	{"Q", "Q 输出"},
	{"M", "M 标志"},
	{"V", "V 变量"},
	{"DB", "DB 数据块"},
	{"T", "T 定时器"},
	{"C", "C 计数器"},
}

var DescS7_200 = protocol.Desc{
	Name:    "Simatic-S7-200",
	Version: "1.0",
	Label:   "Simatic S7-200",
	Codes:   Codes,
	Factory: func(tunnel connect.Tunnel, opts protocol.Options) protocol.Protocol {
		return &Simatic{
			handshake1: parseHex(handshake1_200),
			handshake2: parseHex(handshake2_200),
			link:       tunnel,
		}
	},
}

var DescS7_200_Smart = protocol.Desc{
	Name:    "Simatic-S7-200-Smart",
	Version: "1.0",
	Label:   "Simatic S7-200 Smart",
	Codes:   Codes,
	Parser:  ParseAddress,
	Factory: new200Smart,
}

var DescS7_300 = protocol.Desc{
	Name:    "Simatic-S7-300",
	Version: "1.0",
	Label:   "Simatic S7-300",
	Codes:   Codes,
	Factory: func(tunnel connect.Tunnel, opts protocol.Options) protocol.Protocol {
		return &Simatic{
			handshake1: parseHex(handshake1_300),
			handshake2: parseHex(handshake2_300),
			link:       tunnel,
		}
	},
}

var DescS7_400 = protocol.Desc{
	Name:    "Simatic-S7-400",
	Version: "1.0",
	Label:   "Simatic S7-400",
	Codes:   Codes,
	Factory: func(tunnel connect.Tunnel, opts protocol.Options) protocol.Protocol {
		//TODO 设置机架和槽号
		//setRackSlot()
		return &Simatic{
			handshake1: parseHex(handshake1_400),
			handshake2: parseHex(handshake2_400),
			link:       tunnel,
		}
	},
}

var DescS7_1200 = protocol.Desc{
	Name:    "Simatic-S7-1200",
	Version: "1.0",
	Label:   "Simatic S7-1200",
	Codes:   Codes,
	Factory: func(tunnel connect.Tunnel, opts protocol.Options) protocol.Protocol {
		return &Simatic{
			handshake1: parseHex(handshake1_1200),
			handshake2: parseHex(handshake2_1200),
			link:       tunnel,
		}
	},
}

var DescS7_1500 = protocol.Desc{
	Name:    "Simatic-S7-1500",
	Version: "1.0",
	Label:   "Simatic S7-1500",
	Codes:   Codes,
	Factory: func(tunnel connect.Tunnel, opts protocol.Options) protocol.Protocol {
		return &Simatic{
			handshake1: parseHex(handshake1_1500),
			handshake2: parseHex(handshake2_1500),
			link:       tunnel,
		}
	},
}

func new200Smart(tunnel connect.Tunnel, opts protocol.Options) protocol.Protocol {
	s7 := &S7{
		handshake1: parseHex(handshake1_200_smart),
		handshake2: parseHex(handshake2_200_smart),
		link:       tunnel,
	}
	s7.Init()
	return s7
}
