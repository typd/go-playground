package play

import "time"

func show(msg string, delay time.Duration) chan int {
	c := make(chan int)
	go func() {
		time.Sleep(delay)
		println(msg)
		c <- 1
	}()
	return c
}

func TryGor() {
	c1 := show("msg 1", time.Second*3)
	c2 := show("msg 2", time.Second*3)
	c3 := show("msg 3", time.Second*3)
	c4 := show("msg 4", time.Second*3)
	<-c1
	<-c2
	<-c3
	<-c4
	println("all done")
}
