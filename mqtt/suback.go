package mqtt

import (
	"encoding/binary"
)

type SubCode byte

const (
	SUB_CODE_QOS0 SubCode = iota
	SUB_CODE_QOS1
	SUB_CODE_QOS2
	SUB_CODE_ERR = 128
)

type SubAck struct {
	Header

	packetId uint16

	codes []byte
}

func (msg *SubAck) PacketId() uint16 {
	return msg.packetId
}

func (msg *SubAck) SetPacketId(p uint16) {
	msg.dirty = true
	msg.packetId = p
}

func (msg *SubAck) Codes() []byte {
	return msg.codes
}

func (msg *SubAck) AddCode(c SubCode) {
	msg.dirty = true
	msg.codes = append(msg.codes, byte(c))
}

func (msg *SubAck) ClearCode() {
	msg.dirty = true
	msg.codes = msg.codes[0:0]
}

func (msg *SubAck) Decode(buf []byte) error {
	msg.dirty = false

	//total := len(buf)
	offset := 0

	//Header
	msg.header = buf[0]
	offset++

	//Remain Length
	if l, n, err := ReadRemainLength(buf[offset:]); err != nil {
		return err
	} else {
		msg.remainLength = l
		offset += n
	}
	headerLen := offset

	// PacketId
	msg.packetId = binary.BigEndian.Uint16(buf[offset:])
	offset += 2

	// FixHead & VarHead
	msg.head = buf[0:offset]
	//plo := offset

	// Parse Codes
	l := msg.remainLength + headerLen
	msg.codes = buf[offset:l]

	//Payload
	msg.payload = msg.codes

	return nil
}

func (msg *SubAck) Encode() ([]byte, []byte, error) {
	if !msg.dirty {
		return msg.head, msg.payload, nil
	}

	//Remain Length
	msg.remainLength = 0
	//Packet Id
	msg.remainLength += 2

	//FixHead & VarHead
	hl := msg.remainLength

	//Codes
	msg.remainLength += len(msg.codes)

	//pl := msg.remainLength - hl
	hl += 1 + LenLen(msg.remainLength)

	//Alloc buffer
	msg.head = make([]byte, hl)
	//msg.payload = make([]byte, pl)

	//Header
	ho := 0
	msg.head[ho] = msg.header
	ho++

	//Remain Length
	if n, err := WriteRemainLength(msg.head[ho:], msg.remainLength); err != nil {
		return nil, nil, err
	} else {
		ho += n
	}

	//Packet Id
	binary.BigEndian.PutUint16(msg.head[ho:], msg.packetId)
	ho += 2

	//Codes
	msg.payload = msg.codes

	return msg.head, msg.payload, nil
}
