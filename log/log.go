package log

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

//Options 参数
type Options struct {
	Development bool   `yaml:"development"`
	Format      string `yaml:"format,omitempty"`
	Level       string `yaml:"level"`
	Output      struct {
		Filename   string `yaml:"filename"`
		MaxSize    int    `yaml:"max_size"`
		MaxAge     int    `yaml:"max_age"`
		MaxBackups int    `yaml:"max_backups"`
	} `yaml:"output"`
}

func Init(opts Options) {
	if opts.Development {
		logrus.SetFormatter(&logrus.TextFormatter{TimestampFormat: opts.Format})
	} else {
		logrus.SetFormatter(&logrus.JSONFormatter{TimestampFormat: opts.Format})
	}

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	if opts.Output.Filename == "" {
		logrus.SetOutput(os.Stdout)
	} else {
		logrus.SetOutput(&lumberjack.Logger{
			Filename:   opts.Output.Filename,
			MaxSize:    opts.Output.MaxSize,
			MaxAge:     opts.Output.MaxAge,
			MaxBackups: opts.Output.MaxBackups,
			LocalTime:  true,
		})
	}

	// Only log the warning severity or above.
	level, _ := logrus.ParseLevel(opts.Level)
	logrus.SetLevel(level)
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
