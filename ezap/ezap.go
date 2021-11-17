package ezap

import (
	"github.com/fimreal/goutils/ezap/logger"
)

var Logger *logger.Logger

func init() {
	Logger = logger.New()
}

/* 设置输出的日志等级
OPTIONS: debug info warn error fatal panic
*/
func SetLevel(lv string) {
	Logger.SetLevel(lv)
}

func SetProjectName(projectname string) {
	Logger.SetProjectName(projectname)
}

func SetLogFile(filename string) {
	Logger.DisableCaller()
	Logger.SetLogFile(filename)
}

func SetLogrotate(maxsize int, maxage int, maxbackups int, compress bool) {
	Logger.SetLogrotate(maxsize, maxage, maxbackups, compress)
}

func Debug(args ...interface{}) {
	Logger.Debug(args...)
}
func Debugf(template string, args ...interface{}) {
	Logger.Debugf(template, args...)
}
func Debugw(msg string, keysAndValues ...interface{}) {
	Logger.Debugw(msg, keysAndValues...)
}

func Info(args ...interface{}) {
	Logger.Info(args...)
}
func Infof(template string, args ...interface{}) {
	Logger.Infof(template, args...)
}
func Infow(msg string, keysAndValues ...interface{}) {
	Logger.Infow(msg, keysAndValues...)
}

func Warn(args ...interface{}) {
	Logger.Warn(args...)
}
func Warnf(template string, args ...interface{}) {
	Logger.Warnf(template, args...)
}
func Warnw(msg string, keysAndValues ...interface{}) {
	Logger.Warnw(msg, keysAndValues...)
}

func Error(args ...interface{}) {
	Logger.Error(args...)
}
func Errorf(template string, args ...interface{}) {
	Logger.Errorf(template, args...)
}
func Errorw(msg string, keysAndValues ...interface{}) {
	Logger.Errorw(msg, keysAndValues...)
}

func Fatal(args ...interface{}) {
	Logger.Fatal(args...)
}
func Fatalf(template string, args ...interface{}) {
	Logger.Fatalf(template, args...)
}
func Fatalw(msg string, keysAndValues ...interface{}) {
	Logger.Fatalw(msg, keysAndValues...)
}

func Panic(args ...interface{}) {
	Logger.Panic(args...)
}
func Panicf(template string, args ...interface{}) {
	Logger.Panicf(template, args...)
}
func Panicw(msg string, keysAndValues ...interface{}) {
	Logger.Panicw(msg, keysAndValues...)
}
