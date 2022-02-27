package mqtt

import (
	"encoding/binary"
	"fmt"
)

type SubTopic struct {
	topic []byte
	flag  byte
}

func (msg *SubTopic) Topic() []byte {
	return msg.topic
}

func (msg *SubTopic) SetTopic(b []byte) {
	msg.topic = b
}

func (msg *SubTopic) Qos() MsgQos {
	return MsgQos(msg.flag & 0x03) //0000 0011
}

func (msg *SubTopic) SetQos(qos MsgQos) {
	msg.flag &= 0xFC
	msg.flag |= byte(qos)
}

type Subscribe struct {
	Header

	packetId uint16

	topics []*SubTopic
}

func (msg *Subscribe) PacketId() uint16 {
	return msg.packetId
}

func (msg *Subscribe) SetPacketId(p uint16) {
	msg.dirty = true
	msg.packetId = p
}

func (msg *Subscribe) Topics() []*SubTopic {
	return msg.topics
}

func (msg *Subscribe) AddTopic(topic []byte, qos MsgQos) {
	msg.dirty = true
	st := &SubTopic{}
	st.SetTopic(topic)
	st.SetQos(qos)
	msg.topics = append(msg.topics, st)
}

func (msg *Subscribe) ClearTopic() {
	msg.dirty = true
	msg.topics = msg.topics[0:0]
}

func (msg *Subscribe) Decode(buf []byte) error {
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
		msg.topics = append(msg.topics, st)
	}

	//Payload
	msg.payload = buf[plo:offset]

	return nil
}

func (msg *Subscribe) Encode() ([]byte, []byte, error) {
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
		msg.remainLength += 2 + len(t.Topic())
		msg.remainLength += 1
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
		if err := WriteBytes(msg.payload[plo:], t.topic); err != nil {
			return msg.head, nil, err
		} else {
			plo += len(t.topic) + 2
		}
		//Qos
		msg.payload[plo] = t.flag
		plo++
	}

	return msg.head, msg.payload, nil
}
