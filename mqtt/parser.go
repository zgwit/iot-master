package mqtt

import (
	"encoding/binary"
	"log"
)

type Parser struct {
	buf []byte
}

func (p *Parser) Parse(buf []byte) []Packet {
	var b []byte

	//上次剩余
	if p.buf != nil {
		b = append(p.buf, buf...)
		p.buf = nil
	} else {
		//复制内存，避免覆盖
		b = make([]byte, len(buf))
		copy(b, buf)
	}

	//解析
	return p.parse(b)
}


func (p *Parser) parse(buf []byte) []Packet {

	messages := make([]Packet, 0)

	for {
		remain := len(buf)

		if remain < 2 {
			//包头都不够，等待剩余内容
			//可能需要 超时处理
			break
		}


		//读取Remain Length
		rl, rll := binary.Uvarint(buf[1:])
		//TODO 判断是否够

		remainLength := int(rl)
		packLen := remainLength + rll + 1
		if packLen > remain {
			//等待包体
			break
		}

		msg, err := Decode(buf[:packLen])
		if err != nil {
			log.Println(err)
			buf = buf[packLen:]
			continue
		}

		messages = append(messages, msg)

		//切片，继续解析
		buf = buf[packLen:]
	}

	if len(buf) > 0 {
		//p.buf = buf[:]
		p.buf = buf
	}

	return messages
}
