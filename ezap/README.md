#### zap 日志模块的封装

旧版本每次都要 new 创建 logger 很麻烦，新版本参考大佬的用法，在包内 init 完，暴露需要的日志级别，用法接近 logrus。

参考：https://github.com/blessmylovexy/log


用法🌰：

```go
package main

import log "github.com/fimreal/goutils/ezap"

func main() {
	log.Info("info")
	log.Infof("%s", "info 2")
	log.Infow("birthday", "name", "bb", "time", "1996-11-06")

	log.Debug("debug 1")
	log.SetProjectName("[unified backend]")
	log.SetLevel("debug")
	log.Debug("debug 2")
	log.SetLevel(int8(-1))
	log.Debug("debug 3")

	log.DisableJSONFormat()
	log.Error("Undefined error occurred")
}

```
