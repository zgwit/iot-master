package mqtt

import (
	"encoding/binary"
)

type UnSubscribe struct {
	Header

	packetId uint16

	topics [][]byte
}

func (pkt *UnSubscribe) PacketId() uint16 {
	return pkt.packetId
}

func (pkt *UnSubscribe) SetPacketId(p uint16) {
	pkt.dirty = true
	pkt.packetId = p
}

func (pkt *UnSubscribe) Topics() [][]byte {
	return pkt.topics
}

func (pkt *UnSubscribe) AddTopic(topic []byte) {
	pkt.dirty = true
	pkt.topics = append(pkt.topics, topic)
}

func (pkt *UnSubscribe) ClearTopic() {
	pkt.dirty = true
	pkt.topics = pkt.topics[0:0]
}

func (pkt *UnSubscribe) Decode(buf []byte) error {
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
	plo := offset

	// Parse Topics
	for offset-headerLen < pkt.remainLength {
		//Topic
		if b, err := ReadBytes(buf[offset:]); err != nil {
			return err
		} else {
			pkt.AddTopic(b)
			offset += len(b) + 2
		}
	}

	//Payload
	pkt.payload = buf[plo:offset]

	return nil
}

func (pkt *UnSubscribe) Encode() ([]byte, []byte, error) {
	if !pkt.dirty {
		return pkt.head, pkt.payload, nil
	}

	//Remain Length
	pkt.remainLength = 0
	//Packet Id
	pkt.remainLength += 2

	//FixHead & VarHead
	hl := pkt.remainLength

	//Topics
	for _, t := range pkt.topics {
		pkt.remainLength += 2 + len(t)
	}

	pl := pkt.remainLength - hl
	hl += 1 + LenLen(pkt.remainLength)

	//Alloc buffer
	pkt.head = make([]byte, hl)
	pkt.payload = make([]byte, pl)

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

	plo := 0
	//Topics
	for _, t := range pkt.topics {
		//Topic
		if err := WriteBytes(pkt.payload[plo:], t); err != nil {
			return pkt.head, nil, err
		} else {
			plo += len(t) + 2
		}
	}

	return pkt.head, pkt.payload, nil
}
