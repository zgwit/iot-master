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

func (msg *Connack) SessionPresent() bool {
	return msg.confirm&0x01 == 0x01 //0000 0001
}

func (msg *Connack) SetSessionPresent(b bool) {
	msg.dirty = true
	if b {
		msg.confirm |= 0x01 //0000 0001
	} else {
		msg.confirm &= 0xFE //1111 1110
	}
}

func (msg *Connack) Code() ConnackCode {
	return ConnackCode(msg.code)
}

func (msg *Connack) SetCode(c ConnackCode) {
	msg.dirty = true
	msg.code = byte(c)
}

func (msg *Connack) Decode(buf []byte) error {
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

	//1 Confirm
	msg.confirm = buf[offset]
	offset++

	//2 Code
	msg.code = buf[offset]
	offset++

	// FixHead & VarHead
	msg.head = buf[0:offset]

	return nil
}

func (msg *Connack) Encode() ([]byte, []byte, error) {
	if !msg.dirty {
		return msg.head, nil, nil
	}

	//Tips. remain length is fixed 2 & total is fixed 4
	//Remain Length
	msg.remainLength = 0
	//Confirm
	msg.remainLength += 1
	//Code
	msg.remainLength += 1

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

	//1 Confirm
	msg.head[ho] = msg.confirm
	ho++

	//2 Code
	msg.head[ho] = msg.code
	ho++

	return msg.head, nil, nil
}
