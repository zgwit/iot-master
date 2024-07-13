package timer

import (
	"fmt"
	"github.com/god-jason/bucket/log"
	"github.com/god-jason/bucket/pkg/exception"
	"github.com/god-jason/bucket/pool"
	"github.com/robfig/cron/v3"
	"github.com/zgwit/iot-master/v5/base"
	"github.com/zgwit/iot-master/v5/device"
	"github.com/zgwit/iot-master/v5/project"
	"github.com/zgwit/iot-master/v5/space"
	"strconv"
	"strings"
	"time"
)

// Timer 定时场景
type Timer struct {
	Id        string         `json:"_id" bson:"_id"`
	ProjectId string         `json:"project_id" bson:"project_id"`
	SpaceId   string         `json:"space_id" bson:"space_id"`
	Name      string         `json:"name"`
	Clock     int            `json:"clock"`   //启动时间 每天的分钟 1440
	Weekday   []int          `json:"weekday"` //0 1 2 3 4 5 6
	Actions   []*base.Action `json:"actions"` //动作
	Disabled  bool           `json:"disabled"`

	deviceContainer base.DeviceContainer
	entry           cron.EntryID
}

func (s *Timer) Open() (err error) {

	if s.SpaceId != "" {
		spc := space.Get(s.SpaceId)
		if spc != nil {
			s.deviceContainer = spc
		} else {
			return exception.New("找不到空间")
		}
	} else if s.ProjectId != "" {
		prj := project.Get(s.ProjectId)
		if prj != nil {
			s.deviceContainer = prj
		} else {
			return exception.New("找不到项目")
		}
	} else {
		return exception.New("无效场景")
	}

	for _, a := range s.Actions {
		err = a.Init()
		if err != nil {
			return err
		}
	}

	//星期处理
	w := "*"
	if len(s.Weekday) > 0 {
		var ws []string
		for _, day := range s.Weekday {
			if day >= 0 && day <= 7 {
				ws = append(ws, strconv.Itoa(day))
			} else {
				//error
			}
		}
		w = strings.Join(ws, ",")
	}

	//分 时 日 月 星期
	spec := fmt.Sprintf("%d %d * * %s", s.Clock%60, s.Clock/60, w)
	s.entry, err = _cron.AddFunc(spec, func() {
		//池化 避免拥堵
		_ = pool.Insert(s.ExecuteIgnoreError)
	})
	return
}

func (s *Timer) Close() (err error) {
	_cron.Remove(s.entry)
	return
}

func (s *Timer) execute(id string, action *base.Action) error {
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

func (s *Timer) executeParallel(id string, action *base.Action) {
	_ = pool.Insert(func() {
		err := s.execute(id, action)
		if err != nil {
			log.Error(err)
			return
		}
	})
}

func (s *Timer) ExecuteIgnoreError() {
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

func (s *Timer) Execute() error {
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
