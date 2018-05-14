package logex

import (
	"fmt"
	"path"
	"runtime"

	lcf "github.com/ranhuan/logrus-custom-formatter"
	"github.com/sirupsen/logrus"
)

const (
	// PanicLevel level, highest level of severity. Logs and then calls panic with the
	// message passed to Debug, Info, ...
	PanicLevel logrus.Level = iota
	// FatalLevel level. Logs and then calls `os.Exit(1)`. It will exit even if the
	// logging level is set to Panic.
	FatalLevel
	// ErrorLevel level. Logs. Used for errors that should definitely be noted.
	// Commonly used for hooks to send errors to an error tracking service.
	ErrorLevel
	// WarnLevel level. Non-critical entries that deserve eyes.
	WarnLevel
	// InfoLevel level. General operational entries about what's going on inside the
	// application.
	InfoLevel
	// DebugLevel level. Usually only enabled when debugging. Very verbose logging.
	DebugLevel
)

var log *logrus.Logger

func init() {
	Caller := func(e *logrus.Entry, f *lcf.CustomFormatter) (interface{}, error) {
		text := ""
		if _, ok := e.Data["file"]; ok {
			text = fmt.Sprintf("%s:%d", path.Base(e.Data["file"].(string)), e.Data["line"])
			delete(e.Data, "file")
			delete(e.Data, "line")
		}
		return text, nil
	}

	log = logrus.New()
	log.SetLevel(logrus.DebugLevel)
	lcf.WindowsEnableNativeANSI(true)
	template := "%[shortLevelName]s %[ascTime]s %[caller]18s | %-45[message]s%[fields]s\n"
	log.Formatter = lcf.NewFormatter(template, lcf.CustomHandlers{"caller": Caller})
}

func SetLevel(level logrus.Level) {
	logrus.SetLevel(level)
	log.SetLevel(level)
}

func LogDir() string {
	return "/tmp/"
}

func Debug(args ...interface{}) {
	_, file, line, ok := runtime.Caller(1)
	if ok {
		log.WithField("file", file).WithField("line", line).Debug(args...)
	} else {
		log.Debug(args)
	}
}

func Debugf(format string, args ...interface{}) {
	_, file, line, ok := runtime.Caller(1)
	if ok {
		log.WithField("file", file).WithField("line", line).Debugf(format, args...)
	} else {
		log.Debugf(format, args)
	}
}

func Debugln(args ...interface{}) {
	_, file, line, ok := runtime.Caller(1)
	if ok {
		log.WithField("file", file).WithField("line", line).Debugln(args...)
	} else {
		log.Debugln(args)
	}
}

func Info(args ...interface{}) {
	_, file, line, ok := runtime.Caller(1)
	if ok {
		log.WithField("file", file).WithField("line", line).Info(args...)
	} else {
		log.Info(args)
	}
}

func Infof(format string, args ...interface{}) {
	_, file, line, ok := runtime.Caller(1)
	if ok {
		log.WithField("file", file).WithField("line", line).Infof(format, args...)
	} else {
		log.Infof(format, args)
	}
}

func Infoln(args ...interface{}) {
	_, file, line, ok := runtime.Caller(1)
	if ok {
		log.WithField("file", file).WithField("line", line).Infoln(args...)
	} else {
		log.Infoln(args)
	}
}

func Warn(args ...interface{}) {
	_, file, line, ok := runtime.Caller(1)
	if ok {
		log.WithField("file", file).WithField("line", line).Warn(args...)
	} else {
		log.Warn(args)
	}
}

func Warnf(format string, args ...interface{}) {
	_, file, line, ok := runtime.Caller(1)
	if ok {
		log.WithField("file", file).WithField("line", line).Warnf(format, args...)
	} else {
		log.Warnf(format, args)
	}
}

func Warnln(args ...interface{}) {
	_, file, line, ok := runtime.Caller(1)
	if ok {
		log.WithField("file", file).WithField("line", line).Warnln(args...)
	} else {
		log.Warnln(args)
	}
}

func Warning(args ...interface{}) {
	_, file, line, ok := runtime.Caller(1)
	if ok {
		log.WithField("file", file).WithField("line", line).Warning(args...)
	} else {
		log.Warning(args)
	}
}

func Warningf(format string, args ...interface{}) {
	_, file, line, ok := runtime.Caller(1)
	if ok {
		log.WithField("file", file).WithField("line", line).Warningf(format, args...)
	} else {
		log.Warningf(format, args)
	}
}

func Warningln(args ...interface{}) {
	_, file, line, ok := runtime.Caller(1)
	if ok {
		log.WithField("file", file).WithField("line", line).Warningln(args...)
	} else {
		log.Warningln(args)
	}
}

func Error(args ...interface{}) {
	_, file, line, ok := runtime.Caller(1)
	if ok {
		log.WithField("file", file).WithField("line", line).Error(args...)
	} else {
		log.Error(args)
	}
}

func Errorf(format string, args ...interface{}) {
	_, file, line, ok := runtime.Caller(1)
	if ok {
		log.WithField("file", file).WithField("line", line).Errorf(format, args...)
	} else {
		log.Errorf(format, args)
	}
}

func Errorln(args ...interface{}) {
	_, file, line, ok := runtime.Caller(1)
	if ok {
		log.WithField("file", file).WithField("line", line).Errorln(args...)
	} else {
		log.Errorln(args)
	}
}

func Exit(args ...interface{}) {
	log.Fatal(args)
}

func Exitln(args ...interface{}) {
	log.Fatalln(args)
}

func Fatal(args ...interface{}) {
	_, file, line, ok := runtime.Caller(1)
	if ok {
		log.WithField("file", file).WithField("line", line).Fatal(args)
	} else {
		log.Warning(args...)
	}
	log.Fatal(args...)
}

func Fatalf(format string, args ...interface{}) {
	_, file, line, ok := runtime.Caller(1)
	if ok {
		log.WithField("file", file).WithField("line", line).Fatalf(format, args...)
	} else {
		log.Fatalf(format, args...)
	}
}

func Fatalln(args ...interface{}) {
	_, file, line, ok := runtime.Caller(1)
	if ok {
		log.WithField("file", file).WithField("line", line).Fatalln(args...)
	} else {
		log.Fatalln(args...)
	}
}

func Panic(args ...interface{}) {
	_, file, line, ok := runtime.Caller(1)
	if ok {
		log.WithField("file", file).WithField("line", line).Panic(args...)
	} else {
		log.Panic(args...)
	}
}

func Panicf(format string, args ...interface{}) {
	_, file, line, ok := runtime.Caller(1)
	if ok {
		log.WithField("file", file).WithField("line", line).Panicf(format, args...)
	} else {
		log.Panicf(format, args...)
	}
}

func Panicln(args ...interface{}) {
	_, file, line, ok := runtime.Caller(1)
	if ok {
		log.WithField("file", file).WithField("line", line).Panicln(args...)
	} else {
		log.Panicln(args...)
	}
}

func V(level int) *logrus.Entry {
	_, file, line, ok := runtime.Caller(1)
	if ok {
		return log.WithField("file", file).WithField("line", line)
	} else {
		return log.WithField("file", "unknown").WithField("line", 0)
	}
}

func WithError(err error) *logrus.Entry {
	return log.WithError(err)
}

func WithField(key string, value interface{}) *logrus.Entry {
	return log.WithField(key, value)
}

func WithFields(fields logrus.Fields) *logrus.Entry {
	return log.WithFields(fields)
}

func Serialize(args ...interface{}) {
}

func Serializef(args ...interface{}) {

}

func Flush() {
}

func FlushExit() {

}
