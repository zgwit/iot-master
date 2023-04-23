package log

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

func Open() error {

	if options.Caller {
		logrus.SetReportCaller(true)
	}

	if options.Text {
		logrus.SetFormatter(&logrus.TextFormatter{TimestampFormat: options.Format})
	} else {
		logrus.SetFormatter(&logrus.JSONFormatter{TimestampFormat: options.Format})
	}

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	if options.Output.Filename == "" {
		logrus.SetOutput(os.Stdout)
	} else {
		logrus.SetOutput(&lumberjack.Logger{
			Filename:   options.Output.Filename,
			MaxSize:    options.Output.MaxSize,
			MaxAge:     options.Output.MaxAge,
			MaxBackups: options.Output.MaxBackups,
			LocalTime:  true,
		})
	}

	// Only log the warning severity or above.
	level, err := logrus.ParseLevel(options.Level)
	if err != nil {
		return err
	}
	logrus.SetLevel(level)
	return nil
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
