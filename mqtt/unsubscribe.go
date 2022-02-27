package mqtt

import (
	"encoding/binary"
)

type UnSubscribe struct {
	Header

	packetId uint16

	topics [][]byte
}

func (msg *UnSubscribe) PacketId() uint16 {
	return msg.packetId
}

func (msg *UnSubscribe) SetPacketId(p uint16) {
	msg.dirty = true
	msg.packetId = p
}

func (msg *UnSubscribe) Topics() [][]byte {
	return msg.topics
}

func (msg *UnSubscribe) AddTopic(topic []byte) {
	msg.dirty = true
	msg.topics = append(msg.topics, topic)
}

func (msg *UnSubscribe) ClearTopic() {
	msg.dirty = true
	msg.topics = msg.topics[0:0]
}

func (msg *UnSubscribe) Decode(buf []byte) error {
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
	plo := offset

	// Parse Topics
	for offset-headerLen < msg.remainLength {
		//Topic
		if b, err := ReadBytes(buf[offset:]); err != nil {
			return err
		} else {
			msg.AddTopic(b)
			offset += len(b) + 2
		}
	}

	//Payload
	msg.payload = buf[plo:offset]

	return nil
}

func (msg *UnSubscribe) Encode() ([]byte, []byte, error) {
	if !msg.dirty {
		return msg.head, msg.payload, nil
	}

	//Remain Length
	msg.remainLength = 0
	//Packet Id
	msg.remainLength += 2

	//FixHead & VarHead
	hl := msg.remainLength

	//Topics
	for _, t := range msg.topics {
		msg.remainLength += 2 + len(t)
	}

	pl := msg.remainLength - hl
	hl += 1 + LenLen(msg.remainLength)

	//Alloc buffer
	msg.head = make([]byte, hl)
	msg.payload = make([]byte, pl)

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

	plo := 0
	//Topics
	for _, t := range msg.topics {
		//Topic
		if err := WriteBytes(msg.payload[plo:], t); err != nil {
			return msg.head, nil, err
		} else {
			plo += len(t) + 2
		}
	}

	return msg.head, msg.payload, nil
}
