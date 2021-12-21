package runtime

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func GracefullExit(beforestop func()) {
	// 创建 chan 监听退出信号
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		// <-sig
		log.Println("Receive signal ", <-sig)
		beforestop()
		log.Println("stopping app...")
		os.Exit(0)
	}()

	defer func() {
		if err := recover(); err != nil {
			log.Panicln("ERROR: ", err)
		}
	}()
}
