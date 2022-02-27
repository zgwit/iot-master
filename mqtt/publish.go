package mqtt

import (
	"encoding/binary"
	"fmt"
)

type Publish struct {
	Header

	topic []byte

	packetId uint16

	//payload []byte
}

func (pkt *Publish) Topic() []byte {
	return pkt.topic
}

func (pkt *Publish) SetTopic(b []byte) {
	pkt.dirty = true
	pkt.topic = b
}

func (pkt *Publish) PacketId() uint16 {
	return pkt.packetId
}

func (pkt *Publish) SetPacketId(p uint16) {
	pkt.dirty = true
	pkt.packetId = p
}

func (pkt *Publish) Payload() []byte {
	return pkt.payload
}

func (pkt *Publish) SetPayload(p []byte) {
	pkt.dirty = true
	pkt.payload = p
}

func (pkt *Publish) Decode(buf []byte) error {
	pkt.dirty = false

	total := len(buf)
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

	if total < pkt.remainLength+headerLen {
		return fmt.Errorf("payload is not enough expect %d, got %d", pkt.remainLength+headerLen, total)
	}

	//1 Topic
	if b, err := ReadBytes(buf[offset:]); err != nil {
		return err
	} else {
		pkt.topic = b
		offset += len(b) + 2
	}

	//2 PacketId //Only Qos1 Qos2 has packet id
	if pkt.Qos() > Qos0 {
		pkt.packetId = binary.BigEndian.Uint16(buf[offset:])
		offset += 2
	}

	// FixHead & VarHead
	pkt.head = buf[0:offset]
	//plo := offset

	//3 Payload
	l := pkt.remainLength + headerLen
	b := buf[offset:l]
	pkt.payload = b

	offset += len(b)

	return nil
}

func (pkt *Publish) Encode() ([]byte, []byte, error) {
	if !pkt.dirty {
		return pkt.head, pkt.payload, nil
	}

	//Remain Length
	pkt.remainLength = 0
	//Topic
	pkt.remainLength += 2 + len(pkt.topic)
	//PacketId
	if pkt.Qos() > Qos0 {
		pkt.remainLength += 2
	}

	//FixHead & VarHead
	hl := pkt.remainLength

	//Payload
	pkt.remainLength += len(pkt.payload)

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

	//1 Topic
	if err := WriteBytes(pkt.head[ho:], pkt.topic); err != nil {
		return nil, nil, err
	} else {
		ho += len(pkt.topic) + 2
	}

	//2 PacketId
	if pkt.Qos() > Qos0 {
		binary.BigEndian.PutUint16(pkt.head[ho:], pkt.packetId)
		ho += 2
	}

	//3 Payload
	//pkt.payload = payload

	return pkt.head, pkt.payload, nil
}
