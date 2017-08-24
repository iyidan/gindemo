package log

import (
	"path/filepath"

	"github.com/iyidan/goutils/mise"
	"github.com/iyidan/log"

	"github.com/iyidan/gindemo/conf"
)

var logger = log.New()

// Startup log
// wrap the log package
func Startup() {
	appName := conf.GetAPPName()
	if len(appName) == 0 {
		appName = "app"
	}
	logfile := filepath.Join(conf.String("logdir"), appName+".log")
	err := logger.SetOutputByName(logfile)
	if err != nil {
		mise.PanicOnError(err, "log.Startup")
	}
	logger.SetRotateByDay()
	logger.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func Println(v ...interface{}) {
	logger.Info(v...)
}

func Printf(format string, v ...interface{}) {
	logger.Infof(format, v...)
}

func Info(v ...interface{}) {
	logger.Info(v...)
}

func Infof(format string, v ...interface{}) {
	logger.Infof(format, v...)
}

func Debug(v ...interface{}) {
	logger.Debug(v...)
}

func Debugf(format string, v ...interface{}) {
	logger.Debugf(format, v...)
}

func Warn(v ...interface{}) {
	logger.Warning(v...)
}

func Warnf(format string, v ...interface{}) {
	logger.Warningf(format, v...)
}

func Warning(v ...interface{}) {
	logger.Warning(v...)
}

func Warningf(format string, v ...interface{}) {
	logger.Warningf(format, v...)
}

func Error(v ...interface{}) {
	logger.Error(v...)
}

func Errorf(format string, v ...interface{}) {
	logger.Errorf(format, v...)
}

func Fatal(v ...interface{}) {
	logger.Fatal(v...)
}

func Fatalf(format string, v ...interface{}) {
	logger.Fatalf(format, v...)
}

func SetLevelByString(level string) {
	logger.SetLevelByString(level)
}
