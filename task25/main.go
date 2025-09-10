package main

import (
	"fmt"
	"time"
)



func mySleep(t time.Duration) {
		timer := time.NewTimer(t)
		<-timer.C
		timer.Stop()
}

func main() {
	fmt.Println("begin")
	mySleep(3 * time.Second)
	fmt.Println("finish")
}
