package main

import (
	"fmt"
)

type esportivo interface {
	ligarTurbo() string
}

type ferrari struct {
	modelo      string
	turboLigado bool
	velocidade  int
}

func (f ferrari) ligarTurbo() {
	f.turboLigado = true

}

func main() {
	carro1 := ferrari{"F40", false, 0}
	carro1.ligarTurbo()
	fmt.Println(carro1)
}
