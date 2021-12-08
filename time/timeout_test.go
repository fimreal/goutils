package time

import (
	"fmt"
	"testing"
	"time"
)

func TestTimeout(t *testing.T) {
	c := make(chan interface{}, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c <- 0
	}()
	// go func() {
	// 	fmt.Println("done")
	// 	c <- "done"
	// }()
	if err := TimeoutMillisecond(c, 1000); err != nil {
		fmt.Println(err)
		return
	} else {
		panic("done")
	}
}

func TestNotTimeout(t *testing.T) {
	c := make(chan interface{}, 1)
	go func() {
		fmt.Print()
		c <- "done"
	}()
	if err := TimeoutMillisecond(c, 2000); err != nil {
		panic("timeout !!")
	}
	fmt.Println("done")
}
