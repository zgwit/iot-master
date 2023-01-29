package db

import (
	"github.com/zgwit/iot-master/v3/model"
	"xorm.io/xorm"
	"xorm.io/xorm/log"

	//加载数据库驱动
	//_ "github.com/mattn/go-sqlite3" //CGO版本
	//_ "github.com/glebarez/go-sqlite" //纯Go版本 使用ccgo翻译的，有文件锁问题
	_ "github.com/go-sql-driver/mysql"
	//_ "github.com/lib/pq" //PostgreSQL
	//_ "github.com/mattn/go-oci8" //Oracle
	// _ "github.com/denisenkom/go-mssqldb" //Sql Server
)

var Engine *xorm.Engine

func Open(cfg Options) error {
	var err error
	Engine, err = xorm.NewEngine(cfg.Type, cfg.URL)
	if err != nil {
		return err
	}
	if cfg.Debug {
		Engine.ShowSQL(true)
	}

	Engine.SetLogLevel(log.LogLevel(cfg.LogLevel))
	//Engine.SetLogger(logrus.StandardLogger())

	//同步表
	if cfg.Sync {
		err = Sync()
		if err != nil {
			return err
		}
	}

	return nil
}

func Close() error {
	return Engine.Close()
}

func Sync() error {
	return Engine.Sync2(
		new(model.User), new(model.Password),
		new(model.Entrypoint), new(model.Model),
		new(model.Device), new(model.Plugin),
	)
}
