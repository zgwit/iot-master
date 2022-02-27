package mqtt

type Header struct {
	header       byte
	remainLength int

	dirty bool

	head    []byte
	payload []byte
}

func (hdr *Header) Type() PacketType {
	return PacketType((hdr.header & 0xF0) >> 4)
}

func (hdr *Header) SetType(t PacketType) {
	hdr.dirty = true
	hdr.header &= 0x0F // 0000 1111
	hdr.header |= byte(t << 4)
}

func (hdr *Header) Dup() bool {
	return hdr.header&0x08 == 0x08 //0000 1000
}

func (hdr *Header) SetDup(b bool) {
	hdr.dirty = true
	if b {
		hdr.header |= 0x08 //0000 1000
	} else {
		hdr.header &= 0xF7 //1111 0111
	}
}

func (hdr *Header) Qos() MsgQos {
	return MsgQos((hdr.header & 0x06) >> 1) //0000 0110
}

func (hdr *Header) SetQos(qos MsgQos) {
	hdr.dirty = true
	hdr.header &= 0xF9           //1111 1001
	hdr.header |= byte(qos << 1) //0000 0110
}

func (hdr *Header) Retain() bool {
	return hdr.header&0x01 == 0x01
}

func (hdr *Header) SetRetain(b bool) {
	hdr.dirty = true
	if b {
		hdr.header |= 0x01 //0000 0001
	} else {
		hdr.header &= 0xFE //1111 1110
	}
}

func (hdr *Header) RemainLength() int {
	return hdr.remainLength
}
