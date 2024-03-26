package modbus

import (
	"errors"
	"github.com/zgwit/iot-master/v4/pkg/bin"
	"github.com/zgwit/iot-master/v4/pkg/convert"
)

type Poller struct {
	Code    uint8  `json:"code"`
	Address uint16 `json:"address"`
	Size    uint16 `json:"size"` //长度
}

type Mapper struct {
	Code      uint8   `json:"code"`
	Name      string  `json:"name"`              //名称
	Type      string  `json:"type"`              //类型
	Address   int     `json:"address"`           //偏移
	BigEndian bool    `json:"be,omitempty"`      //大端模式
	Rate      float64 `json:"rate,omitempty"`    //倍率
	Correct   float64 `json:"correct,omitempty"` //纠正
	Bits      []Bit   `json:"bits,omitempty"`    //位，1 2 3...
}

type Bit struct {
	Name string `json:"name"` //名称
	Bit  int    `json:"bit"`  //偏移
}

func (m *Mapper) Encode(data any) ([]byte, error) {
	//TODO 支持 布尔，数值，数组

	switch m.Code {
	case 1:
		//convert.ToBool(data) 太范了
		val := convert.ToBool(data)
		//val, ok := data.(bool)
		//if !ok {
		//	return nil, nil, errors.New("应该是布尔值")
		//}
		if val {
			return []byte{0xFF, 00}, nil
		} else {
			return []byte{0x00, 00}, nil
		}
	case 3:
		var ret []byte

		//纠正
		if m.Correct != 0 {
			data = convert.ToFloat64(data) - m.Correct
		}

		//倍率逆转换
		if m.Rate != 0 && m.Rate != 1 {
			data = convert.ToFloat64(data) / m.Rate
		}

		switch m.Type {
		case "short", "int16":
			ret = make([]byte, 2)
			val := convert.ToInt16(data)
			if m.BigEndian {
				bin.WriteUint16(ret, uint16(val))
			} else {
				bin.WriteUint16LittleEndian(ret, uint16(val))
			}
		case "word", "uint16":
			ret = make([]byte, 2)
			val := convert.ToUint16(data)
			if m.BigEndian {
				bin.WriteUint16(ret, val)
			} else {
				bin.WriteUint16LittleEndian(ret, val)
			}
		case "int32", "int":
			ret = make([]byte, 4)
			val := convert.ToInt32(data)
			if m.BigEndian {
				bin.WriteUint32(ret, uint32(val))
			} else {
				bin.WriteUint32LittleEndian(ret, uint32(val))
			}
		case "qword", "uint32", "uint":
			ret = make([]byte, 4)
			val := convert.ToUint32(data)
			if m.BigEndian {
				bin.WriteUint32(ret, val)
			} else {
				bin.WriteUint32LittleEndian(ret, val)
			}
		case "float", "float32":
			ret = make([]byte, 4)
			val := convert.ToFloat32(data)
			if m.BigEndian {
				bin.WriteFloat32(ret, val)
			} else {
				bin.WriteFloat32LittleEndian(ret, val)
			}
		case "double", "float64":
			ret = make([]byte, 8)
			val := convert.ToFloat64(data)
			if m.BigEndian {
				bin.WriteFloat64(ret, val)
			} else {
				bin.WriteFloat64LittleEndian(ret, val)
			}
		}

		return ret, nil
	}

	return nil, errors.New("找不到点位")
}

func (m *Mapper) Parse(base uint16, buf []byte) (any, error) {
	l := len(buf)

	//识别位
	if m.Code == 1 || m.Code == 2 {
		bytes := bin.ExpandBool(buf, int(m.Size))
		l = len(bytes)
		for _, p := range m.Points {
			offset := p.Offset
			if offset >= l {
				continue
			}
			ret[p.Name] = bytes[p.Offset] > 0
		}
		return nil, nil
	}

	//解析16位

		//offset := p.Offset * 2
		offset := p.Offset << 1
		if offset >= l {
			continue
		}

		switch p.Type {
		case "bit", "bool", "boolean":
			var v uint16
			if p.BigEndian {
				v = bin.ParseUint16(buf[offset:])
			} else {
				v = bin.ParseUint16LittleEndian(buf[offset:])
			}
			ret[p.Name] = 1<<(p.Bits-1)&v != 0
		case "short", "int16":
			if p.BigEndian {
				ret[p.Name] = int16(bin.ParseUint16(buf[offset:]))
			} else {
				ret[p.Name] = int16(bin.ParseUint16LittleEndian(buf[offset:]))
			}
			if p.Rate != 0 && p.Rate != 1 {
				ret[p.Name] = float64(ret[p.Name].(int16)) * p.Rate
			}
		case "word", "uint16":
			if p.BigEndian {
				ret[p.Name] = bin.ParseUint16(buf[offset:])
			} else {
				ret[p.Name] = bin.ParseUint16LittleEndian(buf[offset:])
			}
			if p.Rate != 0 && p.Rate != 1 {
				ret[p.Name] = float64(ret[p.Name].(uint16)) * p.Rate
			}
			if m.Bits != nil && len(m.Bits) > 0
		case "int32", "int":
			if p.BigEndian {
				ret[p.Name] = int32(bin.ParseUint32(buf[offset:]))
			} else {
				ret[p.Name] = int32(bin.ParseUint32LittleEndian(buf[offset:]))
			}
			if p.Rate != 0 && p.Rate != 1 {
				ret[p.Name] = float64(ret[p.Name].(int32)) * p.Rate
			}
		case "qword", "uint32", "uint":
			if p.BigEndian {
				ret[p.Name] = bin.ParseUint32(buf[offset:])
			} else {
				ret[p.Name] = bin.ParseUint32LittleEndian(buf[offset:])
			}
			if p.Rate != 0 && p.Rate != 1 {
				ret[p.Name] = float64(ret[p.Name].(uint32)) * p.Rate
			}
		case "float", "float32":
			if p.BigEndian {
				ret[p.Name] = bin.ParseFloat32(buf[offset:])
			} else {
				ret[p.Name] = bin.ParseFloat32LittleEndian(buf[offset:])
			}
			if p.Rate != 0 && p.Rate != 1 {
				ret[p.Name] = float64(ret[p.Name].(float32)) * p.Rate
			}
		case "double", "float64":
			if p.BigEndian {
				ret[p.Name] = bin.ParseFloat64(buf[offset:])
			} else {
				ret[p.Name] = bin.ParseFloat64LittleEndian(buf[offset:])
			}
			if p.Rate != 0 && p.Rate != 1 {
				ret[p.Name] = ret[p.Name].(float64) * p.Rate
			}
		}

}

func lookup(mappers []*Poller, name string) (*Poller, *Mapper) {
	for _, mapper := range mappers {
		for _, point := range mapper.Points {
			if point.Name == name {
				return mapper, point
			}
		}
	}
	return nil, nil
}
