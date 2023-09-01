package log

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
)

func Open() error {

	if options.Caller {
		logrus.SetReportCaller(true)
	}

	logrus.SetFormatter(&formatter{})

	logrus.SetOutput(os.Stdout)
	//TODO 日志文件

	level, err := logrus.ParseLevel(options.Level)
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
	b.WriteString(fmt.Sprintf("%s [%s]", entry.Time.Format("2006-01-02 15:04:05"), entry.Level.String()))

	//打印文件
	if entry.Caller != nil {
		b.WriteString(fmt.Sprintf("%s:%d", entry.Caller.File, entry.Caller.Line))
	}

	//打印值
	for k, v := range entry.Data {
		b.WriteByte(' ')
		b.WriteString(k)
		b.WriteString("=>")
		b.WriteString(fmt.Sprint(v))
	}

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
