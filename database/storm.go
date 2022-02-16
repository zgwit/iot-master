package database

import (
	"github.com/asdine/storm/v3"
	"os"
	"path/filepath"
)

// Master 基础数据库
var Master *storm.DB

// History 历史数据库
var History *storm.DB

// Error 错误数据库
var Error *storm.DB

//Open 打开数据库
func Open(cfg *Options) error {
	if cfg == nil {
		cfg = DefaultOptions()
	}

	err := os.MkdirAll(cfg.Path, os.ModePerm)
	if err != nil {
		return err
	}

	//基础数据
	Master, err = storm.Open(filepath.Join(cfg.Path, "master.db"))
	if err != nil {
		return err
	}

	//历史数据
	History, err = storm.Open(filepath.Join(cfg.Path, "history.db"))
	if err != nil {
		return err
	}

	//错误日志
	Error, err = storm.Open(filepath.Join(cfg.Path, "error.db"))
	if err != nil {
		return err
	}

	return nil
}

//Close 关闭数据库
func Close() error {
	err := Master.Close()
	if err != nil {
		return err
	}

	err = History.Close()
	if err != nil {
		return err
	}

	err = Error.Close()
	if err != nil {
		return err
	}

	return nil
}
