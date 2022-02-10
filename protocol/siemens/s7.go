package siemens

import (
	"errors"
	"github.com/zgwit/iot-master/protocol/helper"
	"github.com/zgwit/iot-master/service"
	"strconv"
	"strings"
)

type s7command struct {
	Code byte
}

type s7address struct {
	s7command
	Block   int
	Address int
}

var commands = map[string]s7command{
	"I":  {0x81},
	"Q":  {0x82},
	"M":  {0x83},
	"D":  {0x84},
	"DB": {0x84},
	"T":  {0x1D},
	"C":  {0x1C},
	"V":  {0x84},
}

func parseSiemensAddress(address string) int {
	if strings.IndexByte(address, '.') < 0 {
		v, _ := strconv.Atoi(address)
		return v
	} else {
		strs := strings.Split(address, ".")
		v1, _ := strconv.Atoi(strs[0])
		v2, _ := strconv.Atoi(strs[1])
		return v1*8 + v2
	}
}

func parseAddress(address string) (addr s7address, err error) {
	//先检查两字节
	k := strings.ToUpper(address[:2])
	if cmd, ok := commands[k]; ok {
		addr.s7command = cmd
		if k == "DB" {
			i := strings.IndexByte(address, '.')
			addr.Block, _ = strconv.Atoi(address[2:i])
			addr.Address = parseSiemensAddress(address[i+1:])
			//addr.Address, _ = strconv.Atoi(address[i+1:])
		} else {
			addr.Address = parseSiemensAddress(address[2:])
		}
		return
	}

	//检测单字节
	k = strings.ToUpper(address[:1])
	if cmd, ok := commands[k]; ok {
		addr.s7command = cmd
		if k == "D" {
			i := strings.IndexByte(address, '.')
			addr.Block, _ = strconv.Atoi(address[1:i])
			addr.Address = parseSiemensAddress(address[i+1:])
			//addr.Address, _ = strconv.Atoi(address[i+1:])
		} else {
			addr.Address = parseSiemensAddress(address[1:])
		}
		return
	}

	err = errors.New("未知消息")
	return
}

var handshake1 = [22]byte{
	0x03, 0x00, 0x00, 0x16, 0x11, 0xE0, 0x00, 0x00, 0x00, 0x01, 0x00, 0xC0, 0x01, 0x0A, 0xC1, 0x02,
	0x01, 0x02, 0xC2, 0x02, 0x01, 0x00,
}

var handshake2 = [25]byte{
	0x03, 0x00, 0x00, 0x19, 0x02, 0xF0, 0x80, 0x32, 0x01, 0x00, 0x00, 0x04, 0x00, 0x00, 0x08, 0x00,
	0x00, 0xF0, 0x00, 0x00, 0x01, 0x00, 0x01, 0x01, 0xE0,
}

var handshake1_200smart = [22]byte{
	0x03, 0x00, 0x00, 0x16, 0x11, 0xE0, 0x00, 0x00, 0x00, 0x01, 0x00, 0xC1, 0x02, 0x10, 0x00, 0xC2,
	0x02, 0x03, 0x00, 0xC0, 0x01, 0x0A,
}

var handshake2_200smart = [25]byte{
	0x03, 0x00, 0x00, 0x19, 0x02, 0xF0, 0x80, 0x32, 0x01, 0x00, 0x00, 0xCC, 0xC1, 0x00, 0x08, 0x00,
	0x00, 0xF0, 0x00, 0x00, 0x01, 0x00, 0x01, 0x03, 0xC0,
}

var handshake1_200 = [22]byte{
	0x03, 0x00, 0x00, 0x16, 0x11, 0xE0, 0x00, 0x00, 0x00, 0x01, 0x00, 0xC1, 0x02, 0x4D, 0x57, 0xC2,
	0x02, 0x4D, 0x57, 0xC0, 0x01, 0x09,
}

var handshake2_200 = [25]byte{
	0x03, 0x00, 0x00, 0x19, 0x02, 0xF0, 0x80, 0x32, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x08, 0x00,
	0x00, 0xF0, 0x00, 0x00, 0x01, 0x00, 0x01, 0x03, 0xC0,
}

type S7 struct {
	link service.Link
}

func (t *S7) PackCommand(cmd []byte) []byte {
	length := len(cmd)

	buf := make([]byte, length+17)
	//TPKT
	buf[0] = 0x03
	buf[1] = 0x00
	helper.WriteUint16(buf[2:], uint16(length+17)) // 长度
	//ISO-COTP
	buf[4] = 0x02 // 固定
	buf[5] = 0xF0
	buf[6] = 0x80

	buf[7] = 0x32 //Protocol ID 协议ID，固定为32
	buf[8] = 0x01 //Message Type(ROSCTR) 1 请求 2 ACK 3 ACK-Data 7 Userdata
	buf[9] = 0x0  //Reserved
	buf[10] = 0x0
	helper.WriteUint16(buf[9:], 0)               // PDU ref 标识序列号
	helper.WriteUint16(buf[13:], uint16(length)) // Param length
	helper.WriteUint16(buf[15:], 0)              // Data length

	//仅出现在Ack-Data消息中
	//buf[17] Error class
	//buf[18] Error Code

	copy(buf[17:], cmd)

	return buf
}

func (t *S7) BuildReadCommand(addr s7address, length uint16) []byte {
	buf := make([]byte, 14)
	buf[0] = 0x04 // 4读 5写
	buf[1] = 1    // 读取块数
	buf[2] = 0x12 //specification type 指定有效值类型
	buf[3] = 0x0A //length 接下来本次地址访问长度
	buf[4] = 0x10 //syntax id 语法标记，ANY
	buf[5] = 0x02 //variable type 1 bit 2 word 3 dint 4 real 5 counter???
	// (byte)(address[ii].Content1 == 0x1D ? 0x1D : address[ii].Content1 == 0x1C ? 0x1C : 0x02);
	helper.WriteUint16(buf[6:], length)                // 访问数据的个数
	helper.WriteUint16(buf[8:], uint16(addr.Block))    //db number DB块编号，如果访问的是DB块的话
	buf[10] = addr.Code                                //area 访问数据类型
	helper.WriteUint24(buf[11:], uint32(addr.Address)) //address 偏移位置

	return t.PackCommand(buf)
}

func (t *S7) BuildWriteCommand(addr s7address, values []byte) []byte {
	length := len(values)

	buf := make([]byte, 14)
	buf[0] = 0x05 // 4读 5写
	buf[1] = 1    // 读取块数
	buf[2] = 0x12 // 指定有效值类型
	buf[3] = 0x0A // 接下来本次地址访问长度
	buf[4] = 0x10 // 语法标记，ANY
	buf[5] = 0x02 // 按字为单位，1位 2字
	// (byte)(address[ii].Content1 == 0x1D ? 0x1D : address[ii].Content1 == 0x1C ? 0x1C : 0x02);
	helper.WriteUint16(buf[6:], uint16(length))        // 访问数据的个数
	helper.WriteUint16(buf[8:], uint16(addr.Block))    // DB块编号，如果访问的是DB块的话
	buf[10] = addr.Code                                // 访问数据类型
	helper.WriteUint24(buf[11:], uint32(addr.Address)) // 偏移位置
	// 按字写入
	buf[14] = 0x00
	buf[15] = 0x04
	helper.WriteUint16(buf[16:], uint16(length*8)) // 按位计算的长度

	//添加数据
	copy(buf[18:], values)

	return t.PackCommand(buf)
}

//Read 读到数据
func (t *S7) Read(address string, length int) ([]byte, error) {

}

//Write 写入数据
func (t *S7) Write(address string, values []byte) error {

}
