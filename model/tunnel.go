package model

import (
	"bytes"
	"encoding/hex"
	"regexp"
	"time"
)

//Tunnel 通道模型
type Tunnel struct {
	ID        int              `json:"id" storm:"id,increment"`
	Name      string           `json:"name"`
	Type      string           `json:"type"` //serial tcp-client tcp-server udp-client udp-server
	Addr      string           `json:"addr"`
	Retry     int              `json:"retry"` //重试
	Register  *RegisterPacket  `json:"register,omitempty"`
	Heartbeat *HeartBeatPacket `json:"heartbeat,omitempty"`
	Serial    *SerialOptions   `json:"serial,omitempty"`
	Disabled  bool             `json:"disabled"`
	Created   time.Time        `json:"created"`
}

//SerialOptions 串口参数
type SerialOptions struct {
	//PortName   string `json:"port_name"`   // /dev/tty.usb.. COM1
	BaudRate   uint `json:"baud_rate"`   //9600 ... 115200 ...
	DataBits   uint `json:"data_bits"`   //5 6 7 8
	StopBits   uint `json:"stop_bits"`   // 1 2
	ParityMode uint `json:"parity_mode"` // 0:NONE 1:ODD 2:EVEN
}

//RegisterPacket 注册包
type RegisterPacket struct {
	Regex  string `json:"regex,omitempty"`
	Length int    `json:"length,omitempty"`

	regex *regexp.Regexp
}

//Check 检查
func (p *RegisterPacket) Check(buf []byte) bool {
	if p.Regex != "" {
		if p.regex == nil {
			p.regex = regexp.MustCompile(p.Regex)
		}
		return p.regex.MatchString(string(buf))
	}
	if p.Length > 0 {
		if len(buf) != p.Length {
			return false
		}
	}
	return true
}

//HeartBeatPacket 心跳包
type HeartBeatPacket struct {
	Timeout int64  `json:"timeout"`
	Regex   string `json:"regex,omitempty"`
	Length  int    `json:"length,omitempty"`
	Hex     string `json:"hex,omitempty"`
	Text    string `json:"text,omitempty"`

	hex   []byte
	regex *regexp.Regexp
	last  int64
}

//Check 检查
func (p *HeartBeatPacket) Check(buf []byte) bool {

	now := time.Now().Unix()
	if p.last == 0 {
		p.last = now
	}
	if p.last+p.Timeout > now {
		p.last = now
		return false
	}
	p.last = now

	if p.Regex != "" {
		if p.regex == nil {
			p.regex = regexp.MustCompile(p.Regex)
		}
		return p.regex.Match(buf)
	}

	if p.Length > 0 {
		if len(buf) != p.Length {
			return false
		}
	}

	if p.Hex != "" {
		if p.hex == nil {
			//var err error
			p.hex, _ = hex.DecodeString(p.Hex)
		}
		return bytes.Equal(p.hex, buf)
	}

	if p.Text != "" {
		return p.Text == string(buf)
	}

	return true
}

//TunnelHistory 通道历史
type TunnelHistory struct {
	ID       int       `json:"id" storm:"id,increment"`
	TunnelID int       `json:"tunnel_id"`
	History  string    `json:"history"`
	Created  time.Time `json:"created"`
}
