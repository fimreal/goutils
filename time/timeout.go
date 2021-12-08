package time

import (
	"fmt"
	"time"
)

/*
传递接收结束信息的管道，以及超时时间[ms]

time.After(time.Duration(duration) * time.Millisecond) = duration * 1ms

用法例：

c := make(chan string, 1)
go func(){
	dosth()
	c <- "done"
}()
if err := TimeoutMillisecond(c, 5000); err != nil {
	return err
}

*/
func TimeoutMillisecond(c <-chan interface{}, duration int) error {
	select {
	case <-c:
		return nil
	case <-time.After(time.Duration(duration) * time.Millisecond):
		return fmt.Errorf("timeout %dms", duration)
	}
}
