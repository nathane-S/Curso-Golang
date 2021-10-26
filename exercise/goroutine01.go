package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, quantidade int) {
	for i := 0; i < quantidade; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteração %d)\n", pessoa, texto, i+1)
	}
}
func main() {
	//fale("Nathane", "Está tudo bem?", 3)
	//fale("Pessoa:", "Tudo certo!", 3)

	go fale("Nathane", "Está tudo bem?", 3)
	go fale("Pessoa:", "Tudo certo!", 3)
	time.Sleep(time.Second * 4)
}
