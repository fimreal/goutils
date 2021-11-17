package ezap

import (
	"testing"
)

// func ExampleInfo() {
//	// 日期每次都不一样，所以不能用这种测试方法
// 	Info("info")
// 	Infow("birthday", "name", "bb", "time", "1996-11-06")
// 	// Output:
// 	// {"lv":"info","ts":"2021-11-17T17:27:42.305+0800","msg":"info"}
// 	// {"lv":"info","ts":"2021-11-17T17:27:42.305+0800","msg":"birthday","name":"bb","time":"1996-11-06"}
// }

func TestInfo(t *testing.T) {
	Info("info")
	Infof("%s", "info 2")
	Infow("birthday", "name", "bb", "time", "1996-11-06")
}

func TestSetLevel(t *testing.T) {
	Debug("debug 1")
	SetLevel("debug")
	Debug("debug 2")
	SetLevel("warn")
	Info("Info 1")
}

func TestEnableCaller(t *testing.T) {
	Info("DefaultLogFormat")
	EnableCaller()
	Info("EnableCaller")
	DisableCaller()
	Info("DisableCaller")
}

func TestJSONFormat(t *testing.T) {
	Info("DefaultConsoleFormat")
	EnableJSONFormat()
	Info("EnableJSONFormat")
	DisableJSONFormat()
	Info("disableJSONFormat")
}

func TestSetProjectName(t *testing.T) {
	SetProjectName("[unified backend]")
	Error("Undefined error occurred")
}
