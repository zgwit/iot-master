package mqtt

import (
	"encoding/binary"
	"fmt"
)

type SubTopic struct {
	topic []byte
	flag  byte
}

func (pkt *SubTopic) Topic() []byte {
	return pkt.topic
}

func (pkt *SubTopic) SetTopic(b []byte) {
	pkt.topic = b
}

func (pkt *SubTopic) Qos() MsgQos {
	return MsgQos(pkt.flag & 0x03) //0000 0011
}

func (pkt *SubTopic) SetQos(qos MsgQos) {
	pkt.flag &= 0xFC
	pkt.flag |= byte(qos)
}

type Subscribe struct {
	Header

	packetId uint16

	topics []*SubTopic
}

func (pkt *Subscribe) PacketId() uint16 {
	return pkt.packetId
}

func (pkt *Subscribe) SetPacketId(p uint16) {
	pkt.dirty = true
	pkt.packetId = p
}

func (pkt *Subscribe) Topics() []*SubTopic {
	return pkt.topics
}

func (pkt *Subscribe) AddTopic(topic []byte, qos MsgQos) {
	pkt.dirty = true
	st := &SubTopic{}
	st.SetTopic(topic)
	st.SetQos(qos)
	pkt.topics = append(pkt.topics, st)
}

func (pkt *Subscribe) ClearTopic() {
	pkt.dirty = true
	pkt.topics = pkt.topics[0:0]
}

func (pkt *Subscribe) Decode(buf []byte) error {
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
		st := &SubTopic{}
		//Topic
		if b, err := ReadBytes(buf[offset:]); err != nil {
			return err
		} else {
			st.SetTopic(b)
			offset += len(b) + 2
		}
		//Qos
		qos := buf[offset]
		if (qos & 0x03) != qos {
			return fmt.Errorf("Topic Qos %x", qos)
		}
		st.SetQos(MsgQos(qos))
		offset++
		pkt.topics = append(pkt.topics, st)
	}

	//Payload
	pkt.payload = buf[plo:offset]

	return nil
}

func (pkt *Subscribe) Encode() ([]byte, []byte, error) {
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
		pkt.remainLength += 2 + len(t.Topic())
		pkt.remainLength += 1
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
		if err := WriteBytes(pkt.payload[plo:], t.topic); err != nil {
			return pkt.head, nil, err
		} else {
			plo += len(t.topic) + 2
		}
		//Qos
		pkt.payload[plo] = t.flag
		plo++
	}

	return pkt.head, pkt.payload, nil
}
