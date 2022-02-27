package mqtt

import (
	"fmt"
)

type DisConnect struct {
	Header
}

func (pkt *DisConnect) Decode(buf []byte) error {
	pkt.dirty = false

	//Tips. remain length is fixed 0 & total is fixed 2
	total := len(buf)
	if total < 2 {
		return fmt.Errorf("DisConnect expect fixed 2 bytes, got %d", total)
	}

	offset := 0

	//Header
	pkt.header = buf[0]
	offset++

	//Remain Length
	if l, n, err := ReadRemainLength(buf[offset:]); err != nil {
		return err
	} else if l != 0 {
		return fmt.Errorf("Remain length must be 0, got %d", l)
	} else {
		pkt.remainLength = l
		offset += n
	}

	// FixHead & VarHead
	pkt.head = buf[0:offset]

	return nil
}

func (pkt *DisConnect) Encode() ([]byte, []byte, error) {
	if !pkt.dirty {
		return pkt.head, nil, nil
	}

	//Tips. remain length is fixed 0 & total is fixed 2
	//Remain Length
	pkt.remainLength = 0

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

	return pkt.head, nil, nil
}
