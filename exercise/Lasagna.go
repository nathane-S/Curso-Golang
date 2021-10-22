package main

import (
	"fmt"
)

const OvenTime int = 40

func RemainingOvenTime(time int) int {
	return OvenTime - time
}

func PreparationTime(layers int) int {

	return 2 * layers
}

func ElapsedTime(PreparationTime, expected int) int {
	return PreparationTime*2 + expected
}

func main() {
	fmt.Println(RemainingOvenTime(30))
	fmt.Println(PreparationTime(2))
	fmt.Println(ElapsedTime(PreparationTime(3), 20))
}
