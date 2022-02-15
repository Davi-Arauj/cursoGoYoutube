package main

import "fmt"

type pessoa struct {
	Nome      string
	Sobrenome string
	sabores   []string
}

func main() {

	pessoa1 := pessoa{
		Nome:      "Davi",
		Sobrenome: "Araujo",
		sabores:   []string{"biscoito", "chocolate"},
	}

	fmt.Println(pessoa1.Nome, pessoa1.Sobrenome)
	for _, v := range pessoa1.sabores {
		fmt.Println("\t", v)
	}
}
