package model

import "time"

type Camera struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`

	Url  string `json:"url"`
	H264 bool   `json:"h264"`

	Disabled bool      `json:"disabled"`
	Updated  time.Time `json:"updated" xorm:"updated"`
	Created  time.Time `json:"created" xorm:"created"`
}

type CameraEx struct {
	Camera  `xorm:"extends"`
	Running bool `json:"running"`
}

func (s *CameraEx) TableName() string {
	return "camera"
}
