package main

import (
	"fmt"
	"time"
)

func giveValues(s int, messages chan int) {
	//messages := make(chan int)
	s = 11
	for i := 0; i < 10; i++ {
		messages <- i
	}
	//messages <- s
	//return messages
}
func main() {
	c := make(chan int)
	go giveValues(10, c)



	for {
		select {
		case s := <-c:
			fmt.Println(s)
		case <-time.After(10 * time.Second):
			fmt.Println("you are too slow...")
			return
		}
	}

	//fmt.Println(<-messages)
}
