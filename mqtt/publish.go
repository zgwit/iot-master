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

func (msg *Publish) Topic() []byte {
	return msg.topic
}

func (msg *Publish) SetTopic(b []byte) {
	msg.dirty = true
	msg.topic = b
}

func (msg *Publish) PacketId() uint16 {
	return msg.packetId
}

func (msg *Publish) SetPacketId(p uint16) {
	msg.dirty = true
	msg.packetId = p
}

func (msg *Publish) Payload() []byte {
	return msg.payload
}

func (msg *Publish) SetPayload(p []byte) {
	msg.dirty = true
	msg.payload = p
}

func (msg *Publish) Decode(buf []byte) error {
	msg.dirty = false

	total := len(buf)
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

	if total < msg.remainLength+headerLen {
		fmt.Errorf("Payload is not enough expect %d, got %d", msg.remainLength+headerLen, total)
	}

	//1 Topic
	if b, err := ReadBytes(buf[offset:]); err != nil {
		return err
	} else {
		msg.topic = b
		offset += len(b) + 2
	}

	//2 PacketId //Only Qos1 Qos2 has packet id
	if msg.Qos() > Qos0 {
		msg.packetId = binary.BigEndian.Uint16(buf[offset:])
		offset += 2
	}

	// FixHead & VarHead
	msg.head = buf[0:offset]
	//plo := offset

	//3 Payload
	l := msg.remainLength + headerLen
	b := buf[offset:l]
	msg.payload = b

	offset += len(b)

	return nil
}

func (msg *Publish) Encode() ([]byte, []byte, error) {
	if !msg.dirty {
		return msg.head, msg.payload, nil
	}

	//Remain Length
	msg.remainLength = 0
	//Topic
	msg.remainLength += 2 + len(msg.topic)
	//PacketId
	if msg.Qos() > Qos0 {
		msg.remainLength += 2
	}

	//FixHead & VarHead
	hl := msg.remainLength

	//Payload
	msg.remainLength += len(msg.payload)

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

	//1 Topic
	if err := WriteBytes(msg.head[ho:], msg.topic); err != nil {
		return nil, nil, err
	} else {
		ho += len(msg.topic) + 2
	}

	//2 PacketId
	if msg.Qos() > Qos0 {
		binary.BigEndian.PutUint16(msg.head[ho:], msg.packetId)
		ho += 2
	}

	//3 Payload
	//msg.payload = payload

	return msg.head, msg.payload, nil
}
