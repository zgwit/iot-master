package mqtt

import (
	"fmt"
)

type ConnackCode byte

const (
	CONNACK_ACCEPTED ConnackCode = iota
	CONNACK_ERROR_VERSION
	CONNACK_INVALID_CLIENT_ID
	CONNACK_UNAVAILABLE
	CONNACK_INVALID_USERNAME_PASSWORD
	CONNACK_NOT_AUTHORIZED
)

type Connack struct {
	Header

	confirm byte
	code    byte
}

func (pkt *Connack) SessionPresent() bool {
	return pkt.confirm&0x01 == 0x01 //0000 0001
}

func (pkt *Connack) SetSessionPresent(b bool) {
	pkt.dirty = true
	if b {
		pkt.confirm |= 0x01 //0000 0001
	} else {
		pkt.confirm &= 0xFE //1111 1110
	}
}

func (pkt *Connack) Code() ConnackCode {
	return ConnackCode(pkt.code)
}

func (pkt *Connack) SetCode(c ConnackCode) {
	pkt.dirty = true
	pkt.code = byte(c)
}

func (pkt *Connack) Decode(buf []byte) error {
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

	//1 Confirm
	pkt.confirm = buf[offset]
	offset++

	//2 Code
	pkt.code = buf[offset]
	offset++

	// FixHead & VarHead
	pkt.head = buf[0:offset]

	return nil
}

func (pkt *Connack) Encode() ([]byte, []byte, error) {
	if !pkt.dirty {
		return pkt.head, nil, nil
	}

	//Tips. remain length is fixed 2 & total is fixed 4
	//Remain Length
	pkt.remainLength = 0
	//Confirm
	pkt.remainLength += 1
	//Code
	pkt.remainLength += 1

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

	//1 Confirm
	pkt.head[ho] = pkt.confirm
	ho++

	//2 Code
	pkt.head[ho] = pkt.code
	ho++

	return pkt.head, nil, nil
}
