package modbus

import (
	"errors"
	"fmt"
	"github.com/zgwit/iot-master/v4/pkg/bin"
	"github.com/zgwit/iot-master/v4/pkg/convert"
)

type Mapper struct {
	Code      uint8   `json:"code"`
	Name      string  `json:"name"`              //名称
	Type      string  `json:"type"`              //类型
	Address   uint16  `json:"address"`           //偏移
	BigEndian bool    `json:"be,omitempty"`      //大端模式
	Rate      float64 `json:"rate,omitempty"`    //倍率
	Correct   float64 `json:"correct,omitempty"` //纠正
	Bits      []*Bit  `json:"bits,omitempty"`    //位，1 2 3...
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

func (m *Mapper) Parse(address uint16, buf []byte) (any, error) {
	l := len(buf)

	//defer recover()

	//识别位
	//if m.Code == 1 || m.Code == 2 {
	//	bytes := bin.ExpandBool(buf, int(m.Length))
	//	l = len(bytes)
	//	for _, p := range m.Points {
	//		offset := m.Offset
	//		if offset >= l {
	//			continue
	//		}
	//		ret = bytes[m.Offset] > 0
	//	}
	//	return nil, nil
	//}

	//解析16位

	offset := int((m.Address - address) * 2)
	//offset := m.Offset << 1
	if offset >= l {
		return nil, errors.New("长度不够")
	}

	var ret any
	switch m.Type {
	case "short", "int16":
		if m.BigEndian {
			ret = int16(bin.ParseUint16(buf[offset:]))
		} else {
			ret = int16(bin.ParseUint16LittleEndian(buf[offset:]))
		}
		if m.Rate != 0 && m.Rate != 1 {
			ret = float64(ret.(int16)) * m.Rate
		}
		if m.Correct != 0 {
			ret = float64(ret.(int16)) + m.Correct
		}
	case "word", "uint16":
		if m.BigEndian {
			ret = bin.ParseUint16(buf[offset:])
		} else {
			ret = bin.ParseUint16LittleEndian(buf[offset:])
		}
		//取位
		if m.Bits != nil && len(m.Bits) > 0 {
			rets := make(map[string]bool)
			for _, b := range m.Bits {
				rets[b.Name] = (ret.(uint16))&(1<<b.Bit) > 0
			}
			return rets, nil
		}

		if m.Rate != 0 && m.Rate != 1 {
			ret = float64(ret.(uint16)) * m.Rate
		}
		if m.Correct != 0 {
			ret = float64(ret.(int16)) + m.Correct
		}
	case "int32", "int":
		if m.BigEndian {
			ret = int32(bin.ParseUint32(buf[offset:]))
		} else {
			ret = int32(bin.ParseUint32LittleEndian(buf[offset:]))
		}
		if m.Rate != 0 && m.Rate != 1 {
			ret = float64(ret.(int32)) * m.Rate
		}
		if m.Correct != 0 {
			ret = float64(ret.(int16)) + m.Correct
		}
	case "qword", "uint32", "uint":
		if m.BigEndian {
			ret = bin.ParseUint32(buf[offset:])
		} else {
			ret = bin.ParseUint32LittleEndian(buf[offset:])
		}
		if m.Rate != 0 && m.Rate != 1 {
			ret = float64(ret.(uint32)) * m.Rate
		}
		if m.Correct != 0 {
			ret = float64(ret.(int16)) + m.Correct
		}
	case "float", "float32":
		if m.BigEndian {
			ret = bin.ParseFloat32(buf[offset:])
		} else {
			ret = bin.ParseFloat32LittleEndian(buf[offset:])
		}
		if m.Rate != 0 && m.Rate != 1 {
			ret = float64(ret.(float32)) * m.Rate
		}
		if m.Correct != 0 {
			ret = float64(ret.(int16)) + m.Correct
		}
	case "double", "float64":
		if m.BigEndian {
			ret = bin.ParseFloat64(buf[offset:])
		} else {
			ret = bin.ParseFloat64LittleEndian(buf[offset:])
		}
		if m.Rate != 0 && m.Rate != 1 {
			ret = ret.(float64) * m.Rate
		}
		if m.Correct != 0 {
			ret = float64(ret.(int16)) + m.Correct
		}
	default:
		return nil, fmt.Errorf("不支持的数据类型 %s", m.Type)
	}

	return ret, nil
}
