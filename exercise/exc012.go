package main

import (
	"fmt"
	"time"
)

func routine(ch chan int) {
	ch <- 1
	ch <- 2
	ch <- 3
	fmt.Println("Performed!")
	ch <- 4

}

func main() {
	ch := make(chan int, 3)
	go routine(ch)

	time.Sleep(time.Second)
	fmt.Println(<-ch)

}
