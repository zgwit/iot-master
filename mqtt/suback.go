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

func (pkt *SubAck) PacketId() uint16 {
	return pkt.packetId
}

func (pkt *SubAck) SetPacketId(p uint16) {
	pkt.dirty = true
	pkt.packetId = p
}

func (pkt *SubAck) Codes() []byte {
	return pkt.codes
}

func (pkt *SubAck) AddCode(c SubCode) {
	pkt.dirty = true
	pkt.codes = append(pkt.codes, byte(c))
}

func (pkt *SubAck) ClearCode() {
	pkt.dirty = true
	pkt.codes = pkt.codes[0:0]
}

func (pkt *SubAck) Decode(buf []byte) error {
	pkt.dirty = false

	//total := len(buf)
	offset := 0

	//Header
	pkt.header = buf[0]
	offset++

	//Remain Length
	if l, n, err := ReadRemainLength(buf[offset:]); err != nil {
		return err
	} else {
		pkt.remainLength = l
		offset += n
	}
	headerLen := offset

	// PacketId
	pkt.packetId = binary.BigEndian.Uint16(buf[offset:])
	offset += 2

	// FixHead & VarHead
	pkt.head = buf[0:offset]
	//plo := offset

	// Parse Codes
	l := pkt.remainLength + headerLen
	pkt.codes = buf[offset:l]

	//Payload
	pkt.payload = pkt.codes

	return nil
}

func (pkt *SubAck) Encode() ([]byte, []byte, error) {
	if !pkt.dirty {
		return pkt.head, pkt.payload, nil
	}

	//Remain Length
	pkt.remainLength = 0
	//Packet Id
	pkt.remainLength += 2

	//FixHead & VarHead
	hl := pkt.remainLength

	//Codes
	pkt.remainLength += len(pkt.codes)

	//pl := pkt.remainLength - hl
	hl += 1 + LenLen(pkt.remainLength)

	//Alloc buffer
	pkt.head = make([]byte, hl)
	//pkt.payload = make([]byte, pl)

	//Header
	ho := 0
	pkt.head[ho] = pkt.header
	ho++

	//Remain Length
	if n, err := WriteRemainLength(pkt.head[ho:], pkt.remainLength); err != nil {
		return nil, nil, err
	} else {
		ho += n
	}

	//Packet Id
	binary.BigEndian.PutUint16(pkt.head[ho:], pkt.packetId)
	ho += 2

	//Codes
	pkt.payload = pkt.codes

	return pkt.head, pkt.payload, nil
}
