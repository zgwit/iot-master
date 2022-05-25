package db

import (
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/zgwit/iot-master/model"
	"xorm.io/xorm"
)

var Engine *xorm.Engine

func Open() error {
	var err error
	Engine, err = xorm.NewEngine("sqlite3", "./test.db")
	if err != nil {
		return err
	}
	Engine.ShowSQL(true)

	//同步表
	err = Engine.Sync2(
		new(model.User), new(model.Password),
		new(model.Tunnel), new(model.Link),
		new(model.Device), new(model.Element),
		new(model.Project), new(model.Template),
		new(model.HMI), new(model.Component),
		new(model.Event),
		new(model.DeviceAlarm), new(model.ProjectAlarm),

	)
	if err != nil {
		return err
	}

	return nil
}

func Close() error {
	return Engine.Close()
}
