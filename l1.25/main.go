package main

import (
	"fmt"
	"time"
)

func mySleep(duration time.Duration) {
	<-time.After(duration)
}

func main() {
	fmt.Println("Start")
	mySleep(2 * time.Second)
	fmt.Println("End after 2 seconds")
}
