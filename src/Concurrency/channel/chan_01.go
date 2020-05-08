// https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/02.7.md
package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // sum send to channel c, arrows point to left
}

func main() {
	s := []int{7, 2, 8, -4, 9, 0};
	c := make(chan int) // define channel named c, default can send and read

	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c

	fmt.Println(x, y, x+y)
}
