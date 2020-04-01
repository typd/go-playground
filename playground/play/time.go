package play

import (
	"time"
	"fmt"
)

func TryTime() {
	printlnf("now is: %v", time.Now())
	println("exe 'time.After(2 * time.Second)'")
	time.After(2 * time.Second)
	printlnf("now is: %v", time.Now())

	println("")

	printlnf("now is: %v", time.Now())
	println("exe 'time.Sleep(2 * time.Second)'")
	time.Sleep(2 * time.Second)
	printlnf("now is: %v", time.Now())
}

func TryTimeout() {
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "got result 1"
	}()
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout after 1 sec")
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "got result 2"
	}()
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout after 3 secs")
	}
}
