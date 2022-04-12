package log

import (
	"github.com/sirupsen/logrus"
	"os"
)

//Options 参数
type Options struct {
	Development bool   `json:"development" yaml:"development"`
	Format      string `json:"format" yaml:"format,omitempty"`
	Level       string `yaml:"level"`
}

func Init(opts Options) {
	if opts.Development {
		logrus.SetFormatter(&logrus.TextFormatter{TimestampFormat: opts.Format})
	} else {
		logrus.SetFormatter(&logrus.JSONFormatter{TimestampFormat: opts.Format})
	}

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	logrus.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	level, _ := logrus.ParseLevel(opts.Level)
	logrus.SetLevel(level)
}

type Fields = logrus.Fields

var WithFields = logrus.WithFields

var Trace = logrus.Trace

var TraceF = logrus.Tracef

var Warn = logrus.Warn

var WarnF = logrus.Warnf

var Info = logrus.Info

var InfoF = logrus.Infof

var Error = logrus.Error

var ErrorF = logrus.Errorf

var Fatal = logrus.Fatal

var FatalF = logrus.Fatalf
