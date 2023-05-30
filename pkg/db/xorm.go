package db

import (
	"xorm.io/xorm"
	"xorm.io/xorm/log"

	//按需加载数据库驱动

	_ "github.com/denisenkom/go-mssqldb" //Sql Server
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"           //PostgreSQL
	_ "github.com/mattn/go-sqlite3" //CGO版本
	//_ "github.com/glebarez/go-sqlite" //纯Go版本 使用ccgo翻译的，偶有文件锁问题
	//_ "modernc.org/sqlite"
	//_ "github.com/mattn/go-oci8"         //Oracle
)

var Engine *xorm.Engine

func Open() error {
	var err error
	Engine, err = xorm.NewEngine(options.Type, options.URL)
	if err != nil {
		return err
	}
	if options.Debug {
		Engine.ShowSQL(true)
	}

	Engine.SetLogLevel(log.LogLevel(options.LogLevel))
	//Engine.SetLogger(logrus.StandardLogger())
	return nil
}

func Close() error {
	return Engine.Close()
}
