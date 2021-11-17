#### zap 日志模块的封装
旧版本每次都要 new 创建 logger 很麻烦，新版本参考大佬的用法，在包内 init 完，暴露需要的日志级别，用法接近 logrus。
参考：https://github.com/blessmylovexy/log


用法举例：
```go


```





用法举例（旧版本，废弃）：
```go
package main

import "github.com/fimreal/goutils/ezap"

func main() {
    // 创建 logger
	log := ezap.New()

	log.Info("info")
	log.Infof("%s", "info 2")

	log.Debug("debug 1")
	log.SetLevel("debug")
	log.Debug("debug 2")

	log.SetProjectName("myapp")
	log.DisableJSONFormat()
	log.Info("hello world")
}
```
