package main

import "fmt"

func main() {
	ch := make(chan int, 1)

	ch <- 2 // enviando valor dados (escrita)
	<-ch    // recebendo dados (leitura)
	ch <- 3
	fmt.Println(<-ch)
}