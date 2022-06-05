package siemens

import (
	"github.com/zgwit/iot-master/connect"
	"github.com/zgwit/iot-master/protocols/protocol"
)

type S7 struct {
	handshake1 []byte
	handshake2 []byte

	link connect.Tunnel
	desc *protocol.Desc
}

func (s *S7) Init() {
	s.link.On("online", func() {
		_ = s.HandShake()
	})
	return
}

func (s *S7) Desc() *protocol.Desc {
	return &DescS7_200_Smart
}

func (s *S7) HandShake() error {
	_, err := s.link.Ask(s.handshake1, 5)
	if err != nil {
		return err
	}
	//TODO 检查结果
	_, err = s.link.Ask(s.handshake2, 5)
	if err != nil {
		return err
	}
	//TODO 检查结果
	return nil
}

func (s *S7) Read(station int, addr protocol.Addr, size int) ([]byte, error) {
	address := addr.(*Address)

	var vt uint8 = VariableTypeWord
	offset := address.Offset
	if address.HasBit {
		vt = VariableTypeBit
		offset = offset*8 + uint32(address.Bits)
	}

	pack := S7Package{
		Type:      MessageTypeJobRequest,
		Reference: 0,
		param: S7Parameter{
			Code:  ParameterTypeRead,
			Count: 1,
			Type:  vt,
			Areas: []S7ParameterArea{
				{
					Code:   address.Code,
					DB:     address.DB,
					Size:   uint16(size),
					Offset: offset,
				},
			},
		},
	}

	cmd := pack.encode()

	buf, err := s.link.Ask(cmd, 5)
	if err != nil {
		return nil, err
	}

	//解析数据
	var resp S7Package
	err = resp.decode(buf, false)
	if err != nil {
		return nil, err
	}

	return resp.data[0].Data, nil
}

func (s *S7) Poll(station int, addr protocol.Addr, size int) ([]byte, error) {
	//TODO 使用Reference
	return s.Read(station, addr, size)
}

func (s *S7) Write(station int, addr protocol.Addr, data []byte) error {
	address := addr.(*Address)
	length := len(data)

	var vt uint8 = VariableTypeWord
	offset := address.Offset
	if address.HasBit {
		vt = VariableTypeBit
		offset = offset*8 + uint32(address.Bits)
	}

	pack := S7Package{
		Type:      MessageTypeJobRequest,
		Reference: 0,
		param: S7Parameter{
			Code:  ParameterTypeWrite,
			Count: 1,
			Type:  vt,
			Areas: []S7ParameterArea{
				{
					Code:   address.Code,
					DB:     address.DB,
					Size:   uint16(length),
					Offset: offset,
				},
			},
		},
		data: []S7Data{{
			Type:  vt,
			Count: uint16(length),
			Data:  data,
		}},
	}

	cmd := pack.encode()

	buf, err := s.link.Ask(cmd, 5)
	if err != nil {
		return err
	}

	//解析结果
	var resp S7Package
	err = resp.decode(buf, true)
	if err != nil {
		return err
	}

	return nil
}
