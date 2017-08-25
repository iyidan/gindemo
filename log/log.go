package log

import (
	"path/filepath"

	"github.com/iyidan/goutils/mise"
	"github.com/iyidan/log"

	"github.com/iyidan/gindemo/conf"
)

var (
	DefaultLogger = log.New()
)

// Startup log
// wrap the log package
func Startup() {
	appName := conf.GetAPPName()
	if len(appName) == 0 {
		appName = "app"
	}
	logfile := filepath.Join(conf.String("logdir"), appName+".log")
	err := DefaultLogger.SetOutputByName(logfile)
	if err != nil {
		mise.PanicOnError(err, "log.Startup")
	}
	DefaultLogger.SetRotateByDay()
	DefaultLogger.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func Println(v ...interface{}) {
	DefaultLogger.Info(v...)
}

func Printf(format string, v ...interface{}) {
	DefaultLogger.Infof(format, v...)
}

func Info(v ...interface{}) {
	DefaultLogger.Info(v...)
}

func Infof(format string, v ...interface{}) {
	DefaultLogger.Infof(format, v...)
}

func Debug(v ...interface{}) {
	DefaultLogger.Debug(v...)
}

func Debugf(format string, v ...interface{}) {
	DefaultLogger.Debugf(format, v...)
}

func Warn(v ...interface{}) {
	DefaultLogger.Warning(v...)
}

func Warnf(format string, v ...interface{}) {
	DefaultLogger.Warningf(format, v...)
}

func Warning(v ...interface{}) {
	DefaultLogger.Warning(v...)
}

func Warningf(format string, v ...interface{}) {
	DefaultLogger.Warningf(format, v...)
}

func Error(v ...interface{}) {
	DefaultLogger.Error(v...)
}

func Errorf(format string, v ...interface{}) {
	DefaultLogger.Errorf(format, v...)
}

func Fatal(v ...interface{}) {
	DefaultLogger.Fatal(v...)
}

func Fatalf(format string, v ...interface{}) {
	DefaultLogger.Fatalf(format, v...)
}

func SetLevelByString(level string) {
	DefaultLogger.SetLevelByString(level)
}
