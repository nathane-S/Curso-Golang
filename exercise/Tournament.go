package main

import (
	"fmt"
	"strings"
)

func main() {
	entrada := `Allegoric Alaskans;Blithering Badgers;win
Devastating Donkeys;Courageous Californians;draw
Devastating Donkeys;Allegoric Alaskans;win
Courageous Californians;Blithering Badgers;loss
Blithering Badgers;Devastating Donkeys;loss
Allegoric Alaskans;Courageous Californians;win`
	jogos := strings.Split(entrada, "\n")
	fmt.Println(jogos)
	fmt.Println(jogos[0])
	resultado := strings.Split(jogos[0], ";")
	partidasA := 0
	partidasB := 0
	partidasC := 0
	partidasD := 0
	fmt.Println(resultado)
	for _, jogo := range jogos {

		switch {
		case strings.Contains(jogo, "Allegoric Alaskans"):
			partidasA += 1
		case strings.Contains(jogo, "Blithering Badgers"):
			partidasB += 1
		case strings.Contains(jogo, "Courageous Californians"):
			partidasC += 1
		case strings.Contains(jogo, "Devastating Donkeys"):
			partidasD += 1
		}
	}
	fmt.Println(partidasA)
	fmt.Println(partidasB)
	fmt.Println(partidasC)
	fmt.Println(partidasD)
}
