package main

import (
	"fmt"
)

type pessoa struct {
	nome      string
	idade     int
	profissao string
	hobbs     []hobb
}
type hobb struct {
	nome       string
	modalidade string
}

func main() {
	artehobb := hobb{
		nome:       "pintar",
		modalidade: "arte",
	}
	esportehobb := hobb{
		nome:       "futebol",
		modalidade: "esporte",
	}
	nathane := pessoa{
		nome:      "Nathane",
		idade:     26,
		profissao: "estagiária",
		hobbs:     []hobb{artehobb, esportehobb},
	}

	fmt.Println(nathane.hobbs[1])
}
