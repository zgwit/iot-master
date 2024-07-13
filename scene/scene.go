package scene

import (
	"github.com/god-jason/bucket/log"
	"github.com/god-jason/bucket/pkg/exception"
	"github.com/god-jason/bucket/pool"
	"github.com/zgwit/iot-master/v5/base"
	"github.com/zgwit/iot-master/v5/device"
	"github.com/zgwit/iot-master/v5/project"
	"github.com/zgwit/iot-master/v5/space"
	"time"
)

type Time struct {
	Start   int   `json:"start"`             //起始时间 每天的分钟
	End     int   `json:"end"`               //结束时间 每天的分钟
	Weekday []int `json:"weekday,omitempty"` //0 1 2 3 4 5 6
}

// Scene 联动场景
type Scene struct {
	Id        string  `json:"_id" bson:"_id"`
	ProjectId string  `json:"project_id" bson:"project_id"`
	SpaceId   string  `json:"space_id" bson:"space_id"`
	Name      string  `json:"name"`
	Times     []*Time `json:"times,omitempty"`

	Condition //组合条件

	Actions  []*base.Action `json:"actions"` //动作
	Disabled bool           `json:"disabled"`

	deviceContainer base.DeviceContainer
	last            bool //上一次判断结果
}

func (s *Scene) Open() error {
	if s.SpaceId != "" {
		spc := space.Get(s.SpaceId)
		if spc != nil {
			spc.WatchValues(s)
			s.deviceContainer = spc
		} else {
			return exception.New("找不到空间")
		}
	} else if s.ProjectId != "" {
		prj := project.Get(s.ProjectId)
		if prj != nil {
			prj.WatchValues(s)
			s.deviceContainer = prj
		} else {
			return exception.New("找不到项目")
		}
	} else {
		return exception.New("无效场景")
	}

	//找设备，注册变化 watch
	//for _, c := range s.Conditions {
	//	for _, cc := range c {
	//		dev := device.Get(cc.DeviceId)
	//		if dev == nil {
	//			return errors.New("设备找不到")
	//		}
	//		dev.WatchValues(s)
	//	}
	//}

	for _, a := range s.Actions {
		err := a.Init()
		if err != nil {
			return err
		}
	}

	return s.Condition.Init()
}

func (s *Scene) Close() error {
	s.last = false

	if s.SpaceId != "" {
		spc := space.Get(s.SpaceId)
		if spc != nil {
			spc.UnWatchValues(s)
		}
	}
	if s.ProjectId != "" {
		prj := project.Get(s.ProjectId)
		if prj != nil {
			prj.UnWatchValues(s)
		}
	}

	//找设备，unwatch
	//for _, c := range s.Conditions {
	//	for _, cc := range c {
	//		dev := device.Get(cc.DeviceId)
	//		if dev != nil {
	//			dev.UnWatchValues(s)
	//		}
	//	}
	//}

	return nil
}

func (s *Scene) OnProjectValuesChange(project, product, device string, values map[string]any) {
	s.OnDeviceValuesChange(product, device, values)
}

func (s *Scene) OnSpaceValuesChange(space, product, device string, values map[string]any) {
	s.OnDeviceValuesChange(product, device, values)
}

func (s *Scene) OnDeviceValuesChange(product, device string, values map[string]any) {
	//检查时间
	if len(s.Times) > 0 {
		now := time.Now()
		minute := now.Hour()*60 + now.Minute()
		weekday := now.Weekday()
		has := false
		for _, t := range s.Times {
			if t.Start < t.End {
				if minute < t.Start || minute > t.End {
					continue
				}
			} else {
				if minute < t.Start && minute > t.End {
					continue
				}
			}

			if len(t.Weekday) > 0 {
				ww := false
				for _, wd := range t.Weekday {
					if int(weekday) == wd {
						ww = true
						break
					}
				}
				if !ww {
					continue
				}
			}
			has = true
		}
		if !has {
			//没有合适的时间段
			return
		}
	}

	//检查条件
	ret, err := s.Condition.Eval()
	if err != nil {
		log.Error(err)
		return
	}

	//执行接口
	if ret && !s.last {
		//执行
		s.ExecuteIgnoreError()
	}
	s.last = ret
}

func (s *Scene) execute(id string, action *base.Action) error {
	dev := device.Get(id)
	if dev != nil {
		args, err := action.Evaluate(dev.Values())
		if err != nil {
			return err
		}
		//等待
		if action.Delay > 0 {
			time.AfterFunc(action.Delay*time.Millisecond, func() {
				_, err = dev.Action(action.Name, args)
				if err != nil {
					log.Error(err)
				}
			})
		} else {
			_, err = dev.Action(action.Name, args)
			if err != nil {
				return err
			}
		}

		return nil
	}
	return exception.New("找不到设备")
}

func (s *Scene) executeParallel(id string, action *base.Action) {
	_ = pool.Insert(func() {
		err := s.execute(id, action)
		if err != nil {
			log.Error(err)
			return
		}
	})
}

func (s *Scene) ExecuteIgnoreError() {
	for _, a := range s.Actions {
		if a.DeviceId != "" {
			s.executeParallel(a.DeviceId, a)
		} else if a.ProductId != "" {
			if s.deviceContainer != nil {
				ids, err := s.deviceContainer.Devices(a.ProductId)
				if err != nil {
					log.Error(err)
					continue
				}
				for _, d := range ids {
					s.executeParallel(d, a)
				}
			} else {
				log.Error("需要指定产品ID")
			}
		} else {
			log.Error("无效的动作")
		}
	}
}

func (s *Scene) Execute() error {
	for _, a := range s.Actions {
		if a.DeviceId != "" {
			err := s.execute(a.DeviceId, a)
			if err != nil {
				return err
			}
		} else if a.ProductId != "" {
			if s.deviceContainer != nil {
				ids, err := s.deviceContainer.Devices(a.ProductId)
				if err != nil {
					return err
				}
				for _, d := range ids {
					err = s.execute(d, a)
					if err != nil {
						return err
					}
				}
			} else {
				return exception.New("需要指定产品ID")
			}
		} else {
			return exception.New("无效的动作")
		}
	}
	return nil
}
