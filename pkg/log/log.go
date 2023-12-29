package log

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/zgwit/iot-master/v4/config"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

func Open() error {

	if config.GetBool(MODULE, "caller") {
		logrus.SetReportCaller(true)
	}

	if config.GetBool(MODULE, "text") {
		logrus.SetFormatter(&formatter{})
	}

	fn := config.GetString(MODULE, "filename")

	if fn == "" {
		//标准输出
		logrus.SetOutput(os.Stdout)
	} else {
		//日志文件
		logFile := &lumberjack.Logger{
			Filename:   fn,
			MaxSize:    config.GetInt(MODULE, "max_size"), // MB
			MaxBackups: config.GetInt(MODULE, "max_backups"),
			MaxAge:     config.GetInt(MODULE, "max_age"), // days
			Compress:   config.GetBool(MODULE, "compress"),
		}

		defer logFile.Close()

		logrus.SetOutput(logFile)
	}

	level, err := logrus.ParseLevel(config.GetString(MODULE, "level"))
	if err != nil {
		return err
	}
	logrus.SetLevel(level)

	return nil
}

type formatter struct {
}

func (f *formatter) Format(entry *logrus.Entry) ([]byte, error) {
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	//打印时间
	b.WriteString(fmt.Sprintf("%s [%s] ", entry.Time.Format("2006-01-02 15:04:05"), entry.Level.String()))

	//打印内容
	b.WriteString(entry.Message)

	//打印值
	for k, v := range entry.Data {
		b.WriteByte(' ')
		b.WriteString(k)
		b.WriteString("=>")
		b.WriteString(fmt.Sprint(v))
	}

	//打印文件
	if entry.Caller != nil {
		b.WriteString(fmt.Sprintf(" %s:%d", entry.Caller.File, entry.Caller.Line))
	}

	//换行
	b.WriteByte('\n')

	return b.Bytes(), nil
}

type Fields = logrus.Fields

var WithFields = logrus.WithFields

var Trace = logrus.Trace

var Tracef = logrus.Tracef

var Warn = logrus.Warn

var Warnf = logrus.Warnf

var Info = logrus.Info

var Infof = logrus.Infof

var Error = logrus.Error

var Errorf = logrus.Errorf

var Fatal = logrus.Fatal

var Fatalf = logrus.Fatalf

var Println = logrus.Println

var Print = logrus.Print

var Printf = logrus.Printf
