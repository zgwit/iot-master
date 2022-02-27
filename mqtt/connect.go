package mqtt

import (
	"encoding/binary"
	"fmt"
	"regexp"
)

var clientIdRegex *regexp.Regexp

func init() {
	clientIdRegex = regexp.MustCompile("^[0-9a-zA-Z _]*$")
}

var SupportedVersions map[byte]string = map[byte]string{
	0x3: "MQIsdp",
	0x4: "MQTT",
}

type Connect struct {
	Header

	protoName []byte

	protoLevel byte

	flag byte

	keepAlive uint16

	clientId    []byte

	willTopic   []byte
	willMessage []byte

	userName []byte
	password []byte
}

func (pkt *Connect) ProtoName() []byte {
	return pkt.protoName
}

func (pkt *Connect) SetProtoName(b []byte) {
	pkt.dirty = true
	pkt.protoName = b
}

func (pkt *Connect) ProtoLevel() byte {
	return pkt.protoLevel
}

func (pkt *Connect) SetProtoLevel(b byte) {
	pkt.dirty = true
	pkt.protoLevel = b
	pkt.protoName = []byte(SupportedVersions[b])
}

func (pkt *Connect) UserNameFlag() bool {
	return pkt.flag&0x80 == 0x80 //1000 0000
}

func (pkt *Connect) SetUserNameFlag(b bool) {
	pkt.dirty = true
	if b {
		pkt.flag |= 0x80 //1000 0000
	} else {
		pkt.flag &= 0x7F //0111 1111
	}
}

func (pkt *Connect) PasswordFlag() bool {
	return pkt.flag&0x40 == 0x40 //0100 0000
}

func (pkt *Connect) SetPasswordFlag(b bool) {
	pkt.dirty = true
	if b {
		pkt.flag |= 0x40 //0100 0000
	} else {
		pkt.flag &= 0xBF //1011 1111
	}
}

func (pkt *Connect) WillRetain() bool {
	return pkt.flag&0x20 == 0x40 //0010 0000
}

func (pkt *Connect) SetWillRetain(b bool) {
	pkt.dirty = true
	if b {
		pkt.flag |= 0x20 //0010 0000
	} else {
		pkt.flag &= 0xDF //1101 1111
	}
	pkt.SetWillFlag(true)
}

func (pkt *Connect) WillQos() MsgQos {
	return MsgQos((pkt.flag & 0x18) >> 2) // 0001 1000
}

func (pkt *Connect) SetWillQos(qos MsgQos) {
	pkt.dirty = true
	pkt.flag &= 0xE7 // 1110 0111
	pkt.flag |= byte(qos << 2)
	pkt.SetWillFlag(true)
}

func (pkt *Connect) WillFlag() bool {
	return pkt.flag&0x04 == 0x04 //0000 0100
}

func (pkt *Connect) SetWillFlag(b bool) {
	pkt.dirty = true
	if b {
		pkt.flag |= 0x04 //0000 0100
	} else {
		pkt.flag &= 0xFB //1111 1011

		pkt.SetWillQos(Qos0)
		pkt.SetWillRetain(false)
	}
}

func (pkt *Connect) KeepAlive() uint16 {
	return pkt.keepAlive
}

func (pkt *Connect) SetKeepAlive(k uint16) {
	pkt.dirty = true
	pkt.keepAlive = k
}

func (pkt *Connect) ClientId() []byte {
	return pkt.clientId
}

func (pkt *Connect) SetClientId(b []byte) {
	pkt.dirty = true
	pkt.clientId = b
	//pkt.ValidClientId()
}

func (pkt *Connect) WillTopic() []byte {
	return pkt.willTopic
}

func (pkt *Connect) SetWillTopic(b []byte) {
	pkt.dirty = true
	pkt.willTopic = b
	pkt.SetWillFlag(true)
}

func (pkt *Connect) WillMessage() []byte {
	return pkt.willMessage
}

func (pkt *Connect) SetWillMessage(b []byte) {
	pkt.dirty = true
	pkt.willMessage = b
	pkt.SetWillFlag(true)
}

func (pkt *Connect) UserName() []byte {
	return pkt.userName
}

func (pkt *Connect) SetUserName(b []byte) {
	pkt.dirty = true
	pkt.userName = b
	pkt.SetUserNameFlag(true)
}

func (pkt *Connect) Password() []byte {
	return pkt.password
}

func (pkt *Connect) SetPassword(b []byte) {
	pkt.dirty = true
	pkt.password = b
	pkt.SetPasswordFlag(true)
}

func (pkt *Connect) CleanSession() bool {
	return pkt.flag&0x02 == 0x02 //0000 0010
}

func (pkt *Connect) SetCleanSession(b bool) {
	pkt.dirty = true
	if b {
		pkt.flag |= 0x02 //0000 0010
	} else {
		pkt.flag &= 0xFD //1111 1101
	}
}

func (pkt *Connect) ValidClientId() bool {

	if pkt.ProtoLevel() == 0x3 {
		return true
	}

	return clientIdRegex.Match(pkt.clientId)
}

func (pkt *Connect) Decode(buf []byte) error {
	pkt.dirty = false

	//total := len(buf)
	//TODO 判断buf长度
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

	//1 Protocol Name
	if b, err := ReadBytes(buf[offset:]); err != nil {
		return err
	} else {
		pkt.protoName = b
		offset += len(b) + 2
	}

	//2 Protocol Level
	pkt.protoLevel = buf[offset]
	if version, ok := SupportedVersions[pkt.ProtoLevel()]; !ok {
		return fmt.Errorf("Protocol level (%d) is not support", pkt.ProtoLevel())
	} else if ver := string(pkt.ProtoName()); ver != version {
		return fmt.Errorf("Protocol name (%s) invalid", ver)
	}
	offset++

	//3 Connect flag
	pkt.flag = buf[offset]
	offset++

	if pkt.flag&0x1 != 0 {
		return fmt.Errorf("Connect Flags (%x) reserved bit 0", pkt.flag)
	}

	if pkt.WillQos() > Qos2 {
		return fmt.Errorf("Invalid WillQoS (%d)", pkt.WillQos())
	}

	if !pkt.WillFlag() && (pkt.WillRetain() || pkt.WillQos() != Qos0) {
		return fmt.Errorf("Invalid WillFlag (%x)", pkt.flag)
	}

	if pkt.UserNameFlag() != pkt.PasswordFlag() {
		return fmt.Errorf("UserName Password must be both exists or not (%x)", pkt.flag)
	}

	//4 Keep Alive
	pkt.keepAlive = binary.BigEndian.Uint16(buf[offset:])
	offset += 2

	// FixHead & VarHead
	pkt.head = buf[0:offset]
	plo := offset

	//5 ClientId
	if b, err := ReadBytes(buf[offset:]); err != nil {
		return err
	} else {
		pkt.clientId = b
		ln := len(b)
		offset += ln + 2

		// None ClientId, Must Clean Session
		if ln == 2 && !pkt.CleanSession() {
			return fmt.Errorf("None ClientId, Must Clean Session (%x)", pkt.flag)
		}

		// ClientId at most 23 characters
		if ln > 128+2 {
			return fmt.Errorf("Too long ClientId (%s)", string(pkt.ClientId()))
		}

		// ClientId 0-9, a-z, A-Z
		if ln > 0 && !pkt.ValidClientId() {
			return fmt.Errorf("Invalid ClientId (%s)", string(pkt.ClientId()))
		}
	}

	//6 Will Topic & Packet
	if pkt.WillFlag() {
		if b, err := ReadBytes(buf[offset:]); err != nil {
			return err
		} else {
			pkt.willTopic = b
			offset += len(b) + 2
		}

		if b, err := ReadBytes(buf[offset:]); err != nil {
			return err
		} else {
			pkt.willMessage = b
			offset += len(b) + 2
		}
	}

	//7 UserName & Password
	if pkt.UserNameFlag() {
		if b, err := ReadBytes(buf[offset:]); err != nil {
			return err
		} else {
			pkt.userName = b
			offset += len(b) + 2
		}

		if b, err := ReadBytes(buf[offset:]); err != nil {
			return err
		} else {
			pkt.password = b
			offset += len(b) + 2
		}
	}

	//Payload
	pkt.payload = buf[plo:offset]

	return nil
}

func (pkt *Connect) Encode() ([]byte, []byte, error) {
	if !pkt.dirty {
		return pkt.head, pkt.payload, nil
	}

	//Remain Length
	pkt.remainLength = 0
	//Protocol Name
	pkt.remainLength += 2 + len(pkt.protoName)
	//Protocol Level
	pkt.remainLength += 1
	//Connect Flags
	pkt.remainLength += 1
	//Keep Alive
	pkt.remainLength += 1

	//FixHead & VarHead
	hl := pkt.remainLength

	//ClientId
	pkt.remainLength += 2 + len(pkt.clientId)
	//Will Topic & Packet
	if pkt.WillFlag() {
		pkt.remainLength += 2 + len(pkt.willTopic)
		pkt.remainLength += 2 + len(pkt.willMessage)
	}
	//UserName & Password
	if pkt.UserNameFlag() {
		pkt.remainLength += 2 + len(pkt.userName)
		pkt.remainLength += 2 + len(pkt.password)
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

	//1 Protocol Name
	if err := WriteBytes(pkt.head[ho:], pkt.protoName); err != nil {
		return nil, nil, err
	} else {
		ho += len(pkt.protoName) + 2
	}

	//2 Protocol Level
	pkt.head[ho] = pkt.protoLevel
	ho++

	//3 Connect Flags
	pkt.head[ho] = pkt.flag
	ho++

	//4 Keep Alive
	binary.BigEndian.PutUint16(pkt.head[ho:], pkt.keepAlive)
	ho += 2

	plo := 0
	//5 ClientId
	if err := WriteBytes(pkt.payload[plo:], pkt.clientId); err != nil {
		return pkt.head, nil, err
	} else {
		plo += len(pkt.clientId) + 2
	}

	//6 Will Topic & Packet
	if pkt.WillFlag() {
		if err := WriteBytes(pkt.payload[plo:], pkt.willTopic); err != nil {
			return pkt.head, nil, err
		} else {
			plo += len(pkt.willTopic) + 2
		}

		if err := WriteBytes(pkt.payload[plo:], pkt.willMessage); err != nil {
			return pkt.head, nil, err
		} else {
			plo += len(pkt.willMessage) + 2
		}
	}

	//7 UserName & Password
	if pkt.UserNameFlag() {
		if err := WriteBytes(pkt.payload[plo:], pkt.userName); err != nil {
			return pkt.head, nil, err
		} else {
			plo += len(pkt.userName) + 2
		}

		if err := WriteBytes(pkt.payload[plo:], pkt.password); err != nil {
			return pkt.head, nil, err
		} else {
			plo += len(pkt.password) + 2
		}
	}

	return pkt.head, pkt.payload, nil
}
