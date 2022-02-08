package database

import (
	"github.com/asdine/storm/v3"
	"github.com/zgwit/iot-master/config"
	"path/filepath"
)

// Storm 基础数据库
var Storm *storm.DB

// History 历史数据库
var History *storm.DB

// Error 错误数据库
var Error *storm.DB

// Project 项目
var Project storm.Node

// Device 设备
var Device storm.Node

// Service 服务
var Service storm.Node

// Link 连接
var Link storm.Node

// User 用户
var User storm.Node

// Password 密码
var Password storm.Node

// ProjectHistory 项目历史
var ProjectHistory storm.Node

// ProjectHistoryAlarm 项目报警历史
var ProjectHistoryAlarm storm.Node

// ProjectHistoryJob 项目任务历史
var ProjectHistoryJob storm.Node

// DeviceHistory 设备历史
var DeviceHistory storm.Node

// DeviceHistoryAlarm 设备报警历史
var DeviceHistoryAlarm storm.Node

// DeviceHistoryJob 设备任务历史
var DeviceHistoryJob storm.Node

// DeviceHistoryCommand 设备命令历史
var DeviceHistoryCommand storm.Node

// ServiceHistory 服务历史
var ServiceHistory storm.Node

// LinkHistory 连接历史
var LinkHistory storm.Node

// UserHistory 用户历史
var UserHistory storm.Node

func Open() error {
	var err error
	cfg := &config.Config.Database

	//基础数据
	Storm, err = storm.Open(filepath.Join(cfg.Path, "storm.db"))
	if err != nil {
		return err
	}
	Project = Storm.From("project")
	Device = Storm.From("device")
	Service = Storm.From("service")
	Link = Storm.From("link")
	User = Storm.From("user")
	Password = Storm.From("password")

	//历史数据
	History, err = storm.Open(filepath.Join(cfg.Path, "history.db"))
	if err != nil {
		return err
	}
	ProjectHistory = History.From("project")
	ProjectHistoryAlarm = History.From("project", "alarm")
	ProjectHistoryJob = History.From("project", "job")

	DeviceHistory = History.From("device")
	DeviceHistoryAlarm = History.From("device", "alarm")
	DeviceHistoryJob = History.From("device", "job")
	DeviceHistoryCommand = History.From("device", "command")

	ServiceHistory = History.From("service")
	LinkHistory = History.From("link")
	UserHistory = History.From("user")

	//错误日志
	Error, err = storm.Open(filepath.Join(cfg.Path, "error.db"))
	if err != nil {
		return err
	}

	return nil
}

func Close() error {
	err := Storm.Close()
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
