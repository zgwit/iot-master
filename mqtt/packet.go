package mqtt

import (
	"encoding/binary"
	"fmt"
)

type PacketType byte

const (
	RESERVED PacketType = iota
	CONNECT
	CONNACK
	PUBLISH
	PUBACK
	PUBREC
	PUBREL
	PUBCOMP
	SUBSCRIBE
	SUBACK
	UNSUBSCRIBE
	UNSUBACK
	PINGREQ
	PINGRESP
	DISCONNECT
	RESERVED2
)

var packetNames = []string{
	"RESERVED", "CONNECT", "CONNACK", "PUBLISH",
	"PUBACK", "PUBREC", "PUBREL", "PUBCOMP",
	"SUBSCRIBE", "SUBACK", "UNSUBSCRIBE", "UNSUBACK",
	"PINGREQ", "PINGRESP", "DISCONNECT", "RESERVED",
}

func (pkt PacketType) Name() string {
	return packetNames[pkt&0x0F]
}

func (pkt PacketType) NewPacket() Packet {
	var packet Packet
	switch pkt {
	case CONNECT:
		packet = new(Connect)
	case CONNACK:
		packet = new(Connack)
	case PUBLISH:
		packet = new(Publish)
	case PUBACK:
		packet = new(PubAck)
	case PUBREC:
		packet = new(PubRec)
	case PUBREL:
		packet = new(PubRel)
	case PUBCOMP:
		packet = new(PubComp)
	case SUBSCRIBE:
		packet = new(Subscribe)
	case SUBACK:
		packet = new(SubAck)
	case UNSUBSCRIBE:
		packet = new(UnSubscribe)
	case UNSUBACK:
		packet = new(UnSubAck)
	case PINGREQ:
		packet = new(PingReq)
	case PINGRESP:
		packet = new(PingResp)
	case DISCONNECT:
		packet = new(DisConnect)
	default:
		//error
		return nil
	}
	packet.SetType(pkt)
	return packet
}

type MsgQos byte

var qosNames = []string{
	"AtMostOnce", "AtLastOnce", "ExactlyOnce", "QosError",
}

func (qos MsgQos) Name() string {
	// 0000 0011
	return qosNames[qos&0x03]
}

func (qos MsgQos) Level() uint8 {
	return uint8(qos & 0x03)
}

const (
	// Qos0 At most once
	Qos0 MsgQos = iota
	// Qos1 At least once
	Qos1
	// Qos2 Exactly once
	Qos2
)

type Packet interface {
	Type() PacketType
	SetType(t PacketType)
	Dup() bool
	SetDup(b bool)
	Qos() MsgQos
	SetQos(qos MsgQos)
	Retain() bool
	SetRetain(b bool)
	RemainLength() int

	Decode([]byte) error
	Encode() ([]byte, []byte, error)
}

func LenLen(rl int) int {
	if rl <= 127 { //0x7F
		return 1
	} else if rl <= 16383 { //0x7F 7F
		return 2
	} else if rl <= 2097151 { //0x7F 7F 7F
		return 3
	} else {
		return 4
	}
}

func ReadRemainLength(b []byte) (int, int, error) {
	length := len(b)
	size := 1
	for {
		if length < size {
			return 0, size, fmt.Errorf("[ReadRemainLength] Expect at leat %d bytes", 1)
		}

		if b[size-1] > 0x80 {
			size += 1
		} else {
			break
		}

		if size > 4 {
			return 0, size, fmt.Errorf("[ReadRemainLength] Expect at most 4 bytes, got %d", size)
		}
	}
	rl, size := binary.Uvarint(b)
	return int(rl), size, nil
}

func WriteRemainLength(b []byte, rl int) (int, error) {
	length := len(b)
	ll := LenLen(rl)
	if ll > length {
		return 0, fmt.Errorf("[ReadRemainLength] Expect at most %d bytes for remain length", ll)
	}
	return binary.PutUvarint(b, uint64(rl)), nil
}

func ReadBytes(buf []byte) ([]byte, error) {
	if len(buf) < 2 {
		return nil, fmt.Errorf("[readLPBytes] Expect at least %d bytes for prefix", 2)
	}
	length := int(binary.BigEndian.Uint16(buf))
	total := length + 2
	if len(buf) < total {
		return nil, fmt.Errorf("[readLPBytes] Expect at least %d bytes", length+2)
	}
	b := buf[2 : total]
	return b, nil
}

func WriteBytes(buf []byte, b []byte) error {
	length, size := len(b), len(buf)

	if length > 65535 {
		return fmt.Errorf("[writeLPBytes] Too much bytes(%d) to write", length)
	}

	total := length + 2
	if size < total {
		return fmt.Errorf("[writeLPBytes] Expect at least %d bytes", total)
	}

	binary.BigEndian.PutUint16(buf, uint16(length))

	copy(buf[2:], b)

	return nil
}

func BytesDup(buf []byte) []byte {
	b := make([]byte, len(buf))
	copy(b, buf)
	return b
}

func DecodePacket(buf []byte) (Packet, error) {
	mt := PacketType(buf[0] >> 4)
	pkt := mt.NewPacket()
	if pkt != nil {
		err := pkt.Decode(buf)
		return pkt, err
	} else {
		return nil, fmt.Errorf("unknown messege type %d", mt)
	}
}

func EncodePacket(pkt Packet) ([]byte, []byte, error) {
	return pkt.Encode()
}