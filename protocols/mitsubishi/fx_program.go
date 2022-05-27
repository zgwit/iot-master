package mitsubishi

import (
	"errors"
	"github.com/zgwit/iot-master/connect"
	"github.com/zgwit/iot-master/protocol/helper"
	"strconv"
	"strings"
)

type fxProgramCommand struct {
	Offset uint16
	IsBit  bool
	Base   int
}

type fxProgramAddr struct {
	fxProgramCommand
	Address uint16
}

var fxProgramCommands = map[string]fxProgramCommand{
	"X":  {0x0080, true, 8},   //X输入继电器
	"Y":  {0x00A0, true, 8},   //Y输出继电器
	"M":  {0x0100, true, 10},  //M中间继电器
	"D":  {0x1000, false, 10}, //D数据寄存器
	"S":  {0x0000, true, 10},  //S步进继电器
	"TS": {0x00C0, true, 10},  //定时器的触点
	"TC": {0x02C0, true, 10},  //定时器的线圈
	"TN": {0x0800, false, 10}, //定时器的当前值 ?
	"CS": {0x01C0, true, 10},  //计数器的触点
	"CC": {0x03C0, true, 10},  //计数器的线圈
	"CN": {0x0A00, false, 10}, //计数器的当前值 ?
}

func parseFxProgramAddress(address string) (addr fxProgramAddr, err error) {
	var v uint64

	//先检查两字节
	k := strings.ToUpper(address[:2])
	if cmd, ok := fxProgramCommands[k]; ok {
		addr.fxProgramCommand = cmd
		v, err = strconv.ParseUint(address[2:], cmd.Base, 16)
		if k == "CN" && v >= 200 {
			addr.Address = uint16((v-200)*4 + 0x0C00)
		} else if addr.IsBit {
			addr.Address = uint16(int(v)/8 + addr.Base)
		} else {
			addr.Address = uint16(int(v)*2 + addr.Base)
		}

		return
	}

	//检测单字节
	k = strings.ToUpper(address[:1])
	if cmd, ok := fxProgramCommands[k]; ok {
		addr.fxProgramCommand = cmd
		v, err = strconv.ParseUint(address[1:], cmd.Base, 16)
		if k == "D" && v >= 8000 {
			addr.Address = uint16((v-8000)*2 + 0x0E00)
		} else if k == "M" && v >= 8000 {
			addr.Address = uint16((v-8000)/8 + 0x01E0)
		} else if addr.IsBit {
			addr.Address = uint16(int(v)/8 + addr.Base)
		} else {
			addr.Address = uint16(int(v)*2 + addr.Base)
		}

		return
	}

	err = errors.New("未知消息")
	return
}

//FxProgram FX协议
type FxProgram struct {
	link connect.Tunnel
}

//NewFxSerial 新建
func NewFxSerial() *FxProgram {
	return &FxProgram{}
}

//Read 解析
func (t *FxProgram) Read(address string, length int) ([]byte, error) {
	addr, err := parseFxProgramAddress(address)
	if err != nil {
		return nil, err
	}

	buf := make([]byte, 11)
	buf[0] = 0x02                                // STX
	buf[1] = 0x30                                // 命令 Read
	helper.WriteUint16Hex(buf[2:], addr.Address) // 偏移地址
	helper.WriteUint8Hex(buf[6:], uint8(length)) // 读取长度
	buf[8] = 0x03                                // ETX

	// 计算和校验
	var sum uint8 = 0
	for i := 1; i < len(buf)-2; i++ {
		sum += buf[i]
	}

	//最后两位是校验
	helper.WriteUint8Hex(buf[len(buf)-2:], sum)

	//发送请求
	if e := t.link.Write(buf); e != nil {
		return nil, e
	}

	recvLength := length
	if !addr.IsBit {
		recvLength = recvLength << 1
	}

	recv := make([]byte, recvLength+4)
	//length, err = t.link.Read(recv)
	//if err != nil {
	//	return nil, err
	//}

	//NAK ...
	//STX ... ETX 和检验

	if recv[0] == 0x15 {
		return nil, errors.New("返回错误")
	}

	ret := helper.FromHex(recv[1 : length-3])

	return ret, nil
}

//Write 写
func (t *FxProgram) Write(address string, values []byte) error {

	addr, err := parseFxProgramAddress(address)
	if err != nil {
		return err
	}

	//先转成十六进制
	values = helper.ToHex(values)

	length := len(values)

	buf := make([]byte, 11+length)
	buf[0] = 0x02                                // STX
	buf[1] = 0x31                                // 命令 Write
	helper.WriteUint16Hex(buf[2:], addr.Address) // 偏移地址
	helper.WriteUint8Hex(buf[6:], uint8(length)) // 写入长度
	copy(buf[8:], values)                        // 写入内容
	buf[len(buf)-3] = 0x03                       // ETX

	// 计算和校验
	var sum uint8 = 0
	for i := 1; i < len(buf)-2; i++ {
		sum += buf[i]
	}
	//最后两位是校验
	helper.WriteUint8Hex(buf[len(buf)-2:], sum)

	//发送请求
	if e := t.link.Write(buf); e != nil {
		return e
	}

	recv := make([]byte, 1)
	//length, err = t.link.Read(recv)
	//if err != nil {
	//	return err
	//}
	//ACK 0x06
	//NAK 0x15
	if recv[0] == 0x15 {
		return errors.New("错误")
	} else {
		return nil
	}
}
