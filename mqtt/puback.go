package mqtt

import (
	"encoding/binary"
	"fmt"
)

type PubAck struct {
	Header

	packetId uint16
}

func (pkt *PubAck) PacketId() uint16 {
	return pkt.packetId
}

func (pkt *PubAck) SetPacketId(p uint16) {
	pkt.dirty = true
	pkt.packetId = p
}

func (pkt *PubAck) Decode(buf []byte) error {
	pkt.dirty = false

	//Tips. remain length is fixed 2 & total is fixed 4
	total := len(buf)
	if total < 4 {
		return fmt.Errorf("Connack expect fixed 4 bytes (%d)", total)
	}

	offset := 0

	//Header
	pkt.header = buf[0]
	offset++

	//Remain Length
	if l, n, err := ReadRemainLength(buf[offset:]); err != nil {
		return err
	} else if l != 2 {
		return fmt.Errorf("Remain length must be 2, got %d", l)
	} else {
		pkt.remainLength = l
		offset += n
	}

	// PacketId
	pkt.packetId = binary.BigEndian.Uint16(buf[offset:])
	offset += 2

	// FixHead & VarHead
	pkt.head = buf[0:offset]

	return nil
}

func (pkt *PubAck) Encode() ([]byte, []byte, error) {
	if !pkt.dirty {
		return pkt.head, nil, nil
	}

	//Tips. remain length is fixed 2 & total is fixed 4
	//Remain Length
	pkt.remainLength = 0
	//Packet Id
	pkt.remainLength += 2

	//FixHead & VarHead
	hl := pkt.remainLength

	hl += 1 + LenLen(pkt.remainLength)
	//Alloc buffer
	pkt.head = make([]byte, hl)

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

	return pkt.head, nil, nil
}
