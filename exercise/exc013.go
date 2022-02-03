package main

import "fmt"

func main() {
	n := 5
	a := []int{}
	for i := n; i <= n; i-- {
		if i >= 1 {
			a = append(a, i)
			fmt.Println(a)
		}
	}

}