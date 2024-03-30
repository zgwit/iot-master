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
	"slices"
	"time"
)

type Adapter struct {
	modbus  Modbus
	devices []*Device

	index map[string]*Device

	//index map[string]*device.Device
}

func (adapter *Adapter) start(tunnel string, opts types.Options) error {
	err := db.Engine.Where("tunnel_id=?", tunnel).And("disabled!=1").Find(&adapter.devices)
	if err != nil {
		return err
	}

	//if len(adapter.devices) == 0 {
	//	return errors.New("无设备")
	//}

	//索引
	for _, d := range adapter.devices {
		adapter.index[d.Id] = d
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
				d, err := device.Ensure(dev.Id)
				if err != nil {
					log.Error(err)
					continue
				}

				values, err := adapter.Sync(dev.Id)
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

				//d := device.Get(dev.Id)
				if values != nil && len(values) > 0 {
					d.Push(values)
				}
				//_ = pool.Insert(func() {
				//topic := fmt.Sprintf("device/%s/property", dev.Id)
				//mqtt.Publish(topic, values)
			}

			now := time.Now().Unix()
			interval := opts.Int64("poller_interval", 60) //默认5分钟轮询一次
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

func (adapter *Adapter) Mount(device string) error {
	var dev Device
	has, err := db.Engine.ID(device).Get(&dev)
	if err != nil {
		return err
	}
	if !has {
		return errors.New("找不到设备")
	}

	found := false
	for i, d := range adapter.devices {
		if d.Id == device {
			adapter.devices[i] = &dev
			adapter.index[device] = &dev
			found = true
		}
	}
	if !found {
		adapter.devices = append(adapter.devices, &dev)
		adapter.index[device] = &dev
	}
	return nil
}

func (adapter *Adapter) Unmount(device string) error {
	delete(adapter.index, device)
	for i, d := range adapter.devices {
		if d.Id == device {
			slices.Delete(adapter.devices, i, i+1)
			return nil
		}
	}
	return nil
}

func (adapter *Adapter) Get(id, name string) (any, error) {
	dev := device.Get(id)
	if dev == nil {
		return nil, errors.New("设备未上线")
	}
	station := adapter.index[id].Station.Slave

	prod, err := GetProduct(dev.ProductId, dev.ProductVersion)
	if err != nil {
		return nil, err
	}

	mapper := prod.Lookup(name)
	if mapper == nil {
		return nil, errors.New("找不到数据点")
	}

	//此处全部读取了，有些冗余
	data, err := adapter.modbus.Read(station, mapper.Code, mapper.Address, 2)
	if err != nil {
		return nil, err
	}

	return mapper.Parse(mapper.Address, data)
}

func (adapter *Adapter) Set(id, name string, value any) error {
	dev := device.Get(id)
	if dev == nil {
		return errors.New("设备未上线")
	}
	station := adapter.index[id].Station.Slave

	prod, err := GetProduct(dev.ProductId, dev.ProductVersion)
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
	return adapter.modbus.Write(station, mapper.Code, mapper.Address, data)
}

func (adapter *Adapter) Sync(id string) (map[string]any, error) {
	dev := device.Get(id)
	if dev == nil {
		return nil, errors.New("设备未上线")
	}
	station := adapter.index[id].Station.Slave

	prod, err := GetProduct(dev.ProductId, dev.ProductVersion)
	if err != nil {
		return nil, err
	}

	values := make(map[string]any)
	for _, poller := range prod.Pollers {
		data, err := adapter.modbus.Read(station, poller.Code, poller.Address, poller.Length)
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
