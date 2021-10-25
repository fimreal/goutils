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
