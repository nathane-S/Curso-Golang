package main

import (
	"fmt"
	"strings"
)

func main() {

	palavra := "casa"
	letras := strings.Split(strings.ToUpper(palavra), "")
	fmt.Println(letras)
	soma := 0
	for _, letra := range letras {
		switch letra {
		case "A", "E", "I", "O", "U", "L", "N", "R", "S", "T":
			soma += 1
		case "D", "G":
			soma += 2
		case "B", "C", "M", "P":
			soma += 3
		case "F", "H", "V", "W", "Y":
			soma += 4
		case "K":
			soma += 5
		case "J", "X":
			soma += 8
		case "Q", "Z":
			soma += 10
		}

	}
	fmt.Println(soma)
}

/*
Instruções
Dada uma palavra, calcule a pontuação Scrabble para essa palavra.

Valores de Carta

Letter                           Value
A, E, I, O, U, L, N, R, S, T       1
D, G                               2
B, C, M, P                         3
F, H, V, W, Y                      4
K                                  5
J, X                               8
Q, Z                               10
*/
