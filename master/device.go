package master

import (
	"errors"
	"fmt"
	"github.com/antonmedv/expr"
	"iot-master/model"
	"iot-master/pkg/convert"
	"iot-master/pkg/events"
	"iot-master/protocols/protocol"
	"strconv"
	"time"
)

//Device 设备
type Device struct {
	model.Device
	events.EventEmitter

	product *model.Product

	Context map[string]interface{}

	points  []*Point
	pollers []*Poller

	//命令索引
	commandIndex map[string]*model.Command

	protocol protocol.Protocol
	//mapper *Mapper

	running bool
}

func NewDevice(m *model.Device) (*Device, error) {
	dev := &Device{
		Device:       *m,
		Context:      make(map[string]interface{}),
		commandIndex: make(map[string]*model.Command, 0),
		pollers:      make([]*Poller, 0),
	}

	//加载产品
	var err error
	dev.product, err = LoadProduct(dev.ProductId)
	if err != nil {
		return nil, err
	}

	//索引命令
	for _, cmd := range dev.product.Commands {
		dev.commandIndex[cmd.Name] = cmd
	}

	//解析点位
	for _, v := range dev.product.Points {
		dev.points = append(dev.points, &Point{Point: *v, Addr: nil})

		//初始化默认值
		if v.Name != "" {
			if v.Precision > 0 {
				dev.Context[v.Name] = float64(0)
			} else {
				dev.Context[v.Name] = v.Type.Default()
			}
		}
	}

	//初始化
	for _, v := range dev.product.Pollers {
		dev.pollers = append(dev.pollers, &Poller{Poller: *v, Addr: nil, Device: dev})
	}

	return dev, nil
}

func (dev *Device) BindTunnel(tunnel *Tunnel) error {
	if tunnel.protocol == nil {
		return errors.New("通道未加载协议")
	}

	if dev.protocol == tunnel.protocol {
		return nil
	}
	dev.protocol = tunnel.protocol

	proto := dev.protocol.Desc()

	var err error
	//解析点位
	for _, v := range dev.points {
		v.Addr, err = proto.Parser(v.Code, v.Address)
		if err != nil {
			return err
		}
	}

	//初始化
	for _, v := range dev.pollers {
		v.Addr, err = proto.Parser(v.Code, v.Address)
		if err != nil {
			return err
		}
	}

	return err
}

func (dev *Device) onData(data map[string]interface{}) {

	//向上广播
	dev.Emit("data", data)
}

//Start 设备启动
func (dev *Device) Start() error {
	//if dev.running {
	//	return errors.New("已经启动")
	//}

	//找到链接，导入协议
	tunnel := GetTunnel(dev.TunnelId)
	if tunnel == nil {
		return errors.New("找不到链接")
	}

	//绑定链接
	err := dev.BindTunnel(tunnel)
	if err != nil {
		return err
	}

	//采集器
	for _, poller := range dev.pollers {
		err := poller.Start()
		if err != nil {
			return err
		}
	}

	dev.running = true

	return nil
}

//Stop 结束设备
func (dev *Device) Stop() error {
	dev.running = false

	for _, poller := range dev.pollers {
		poller.Stop()
	}
	return nil
}

func (dev *Device) Running() bool {
	return dev.running
}

func (dev *Device) read(addr protocol.Addr, size int) error {
	buf, err := dev.protocol.Read(dev.Station, addr, size)
	if err != nil {
		return err
	}
	values := make(map[string]interface{})
	for _, p := range dev.points {
		val, has := p.Addr.Lookup(buf, addr, p.Type, p.LittleEndian, p.Precision)
		if has {
			values[p.Name] = val
			dev.Context[p.Name] = val
		}
	}

	//处理数据
	dev.onData(values)

	return nil
}

func (dev *Device) Set(name string, value interface{}) error {
	for _, p := range dev.points {
		if p.Name == name {
			if p.Precision > 0 {
				value = convert.ToFloat64(value)
			} else {
				value = p.Type.Normalize(value)
			}
			dev.Context[name] = value

			buf := p.Type.Encode(value, p.LittleEndian, p.Precision)
			//dev.Context[name], _ = p.Type.Decode(buf, p.LittleEndian, p.Precision)
			return dev.protocol.Write(dev.Station, p.Addr, buf)
		}
	}

	//？？？不是数据点的是否要写入？？？
	dev.Context[name] = value
	return nil
}

func (dev *Device) Refresh() error {
	if !dev.running {
		return errors.New("设备未运行")
	}
	for _, poller := range dev.pollers {
		poller.Execute()
	}
	return nil
}

func (dev *Device) RefreshPoint(name string) (interface{}, error) {
	if !dev.running {
		return nil, errors.New("设备未运行")
	}
	for _, p := range dev.points {
		if p.Name == name {
			buf, err := dev.protocol.Read(dev.Station, p.Addr, p.Type.Size())
			if err != nil {
				return 0, err
			}
			val, has := p.Addr.Lookup(buf, p.Addr, p.Type, p.LittleEndian, p.Precision)
			if has {
				dev.Context[p.Name] = val
			}
			return val, nil
		}

	}
	return 0, fmt.Errorf("数据点不存在 %s", name)
}

//Execute 执行命令
func (dev *Device) Execute(command string, argv []interface{}) error {
	cmd, ok := dev.commandIndex[command]
	if !ok {
		return fmt.Errorf("找不到命令：%s", command)
	}

	//参数加入环境变量
	//优先级：参数 > 表达式 > 初始值
	env := make(map[string]interface{})
	for i, v := range argv {
		env["$"+strconv.Itoa(i+1)] = v
	}
	for k, v := range dev.Context {
		env[k] = v
	}

	//执行
	for _, directive := range cmd.Directives {
		var val interface{} = directive.Value
		if directive.Expression != "" {
			v, err := expr.Eval(directive.Expression, env)
			if err != nil {
				return err
			} else {
				val = v
				//val = helper.ToFloat64(v)
				//val = v.(float64)
			}
		}

		//延迟执行
		if directive.Delay > 0 {
			time.AfterFunc(time.Duration(directive.Delay)*time.Millisecond, func() {
				err := dev.Set(directive.Point, val)
				dev.Emit("error", err)
			})
		} else {
			err := dev.Set(directive.Point, val)
			//dev.events.Publish("error", err)
			return err
		}
	}

	return nil
}
