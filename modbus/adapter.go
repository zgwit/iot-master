package modbus

import (
	"errors"
	"github.com/zgwit/iot-master/v4/device"
	"github.com/zgwit/iot-master/v4/pkg/db"
	"github.com/zgwit/iot-master/v4/pkg/log"
	"github.com/zgwit/iot-master/v4/types"
	"go.bug.st/serial"
	"io"
	"net"
	"time"
)

type Adapter struct {
	modbus  Modbus
	devices []*Device

	index map[string]*Device

	//index map[string]*device.Device
}

func (adapter *Adapter) start(tunnel string, opts types.Options) error {
	err := db.Engine.Where("tunnel_id=?", tunnel).And("disabled!=1").
		Asc("modbus_station").Find(&adapter.devices)

	if err != nil {
		return err
	}

	if len(adapter.devices) == 0 {
		return errors.New("无设备")
	}

	//索引
	//adapter.index = make(map[string]*device.Device)
	adapter.index = make(map[string]*Device)
	for _, d := range adapter.devices {
		adapter.index[d.Id] = d
		//adapter.index[d.Id], err = device.Ensure(d.Id)
	}

	//开始轮询
	go func() {

		//设备上线
		//!!! 不能这样做，不然启动服务器会产生大量的消息
		//for _, dev := range adapter.index {
		//	topic := fmt.Sprintf("device/online/%s", dev.Id)
		//	_ = mqtt.Publish(topic, nil)
		//}

	OUT:
		for {
			start := time.Now().Unix()
			for _, dev := range adapter.devices {
				//d := adapter.index[dev.Id]
				d := device.Get(dev.Id)
				values, err := adapter.Sync(d)
				if err != nil {
					log.Error(err)

					//连接断开错误
					if err == io.EOF {
						break OUT
					}

					//网络错误（读超时除外）
					var ne net.Error
					if errors.As(err, &ne) && !ne.Timeout() {
						break OUT
					}

					//串口错误（读超时除外）
					var se *serial.PortError
					if errors.As(err, &se) && (se.Code() != serial.InvalidTimeoutValue) {
						break OUT
					}
				}

				d.Push(values)
				//_ = pool.Insert(func() {
				//topic := fmt.Sprintf("device/%s/property", dev.Id)
				//mqtt.Publish(topic, values)
			}

			now := time.Now().Unix()
			interval := opts.Int64("poller_interval", 300) //默认5分钟轮询一次
			if now-start < interval {
				time.Sleep(time.Second * time.Duration(interval-(now-start)))
			}

			//避免空转，睡眠1分钟（可能有点长）
			if now-start < 1 {
				time.Sleep(time.Minute)
			}
		}

		//设备下线
		//for _, dev := range adapter.devices {
		//	topic := fmt.Sprintf("device/%s/offline", dev.Id)
		//	_ = mqtt.Publish(topic, nil)
		//}
	}()
	return nil
}

func (adapter *Adapter) Get(device *device.Device, name string) (any, error) {
	prod, err := GetProduct(device.ProductId, device.ProductVersion)
	if err != nil {
		return nil, err
	}

	mapper := prod.Lookup(name)
	if mapper == nil {
		return nil, errors.New("找不到数据点")
	}

	//此处全部读取了，有些冗余
	data, err := adapter.modbus.Read(adapter.index[device.Id].ModbusStation, mapper.Code, mapper.Address, 2)
	if err != nil {
		return nil, err
	}

	return mapper.Parse(mapper.Address, data)
}

func (adapter *Adapter) Set(device *device.Device, name string, value any) error {
	prod, err := GetProduct(device.ProductId, device.ProductVersion)
	if err != nil {
		return err
	}

	mapper := prod.Lookup(name)
	if mapper == nil {
		return errors.New("地址找不到")
	}

	data, err := mapper.Encode(value)
	if err != nil {
		return err
	}
	return adapter.modbus.Write(adapter.index[device.Id].ModbusStation, mapper.Code, mapper.Address, data)
}

func (adapter *Adapter) Sync(device *device.Device) (map[string]any, error) {
	values := make(map[string]any)

	prod, err := GetProduct(device.ProductId, device.ProductVersion)
	if err != nil {
		return nil, err
	}

	for _, poller := range prod.Pollers {
		data, err := adapter.modbus.Read(adapter.index[device.Id].ModbusStation, poller.Code, poller.Address, poller.Length)
		if err != nil {
			return nil, err
		}
		err = poller.Parse(prod.Mappers, data, values)
		if err != nil {
			return nil, err
		}
	}

	//TODO 计算器

	return values, nil
}
