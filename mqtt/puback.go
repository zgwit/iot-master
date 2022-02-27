package mqtt

import (
	"encoding/binary"
	"fmt"
)

type PubAck struct {
	Header

	packetId uint16
}

func (msg *PubAck) PacketId() uint16 {
	return msg.packetId
}

func (msg *PubAck) SetPacketId(p uint16) {
	msg.dirty = true
	msg.packetId = p
}

func (msg *PubAck) Decode(buf []byte) error {
	msg.dirty = false

	//Tips. remain length is fixed 2 & total is fixed 4
	total := len(buf)
	if total < 4 {
		return fmt.Errorf("Connack expect fixed 4 bytes (%d)", total)
	}

	offset := 0

	//Header
	msg.header = buf[0]
	offset++

	//Remain Length
	if l, n, err := ReadRemainLength(buf[offset:]); err != nil {
		return err
	} else if l != 2 {
		return fmt.Errorf("Remain length must be 2, got %d", l)
	} else {
		msg.remainLength = l
		offset += n
	}

	// PacketId
	msg.packetId = binary.BigEndian.Uint16(buf[offset:])
	offset += 2

	// FixHead & VarHead
	msg.head = buf[0:offset]

	return nil
}

func (msg *PubAck) Encode() ([]byte, []byte, error) {
	if !msg.dirty {
		return msg.head, nil, nil
	}

	//Tips. remain length is fixed 2 & total is fixed 4
	//Remain Length
	msg.remainLength = 0
	//Packet Id
	msg.remainLength += 2

	//FixHead & VarHead
	hl := msg.remainLength

	hl += 1 + LenLen(msg.remainLength)
	//Alloc buffer
	msg.head = make([]byte, hl)

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

	return msg.head, nil, nil
}
