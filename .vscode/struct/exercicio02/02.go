package main

import "fmt"

type pessoa struct {
	Nome      string
	Sobrenome string
	sabores   []string
}

func main() {

	mepe := make(map[string]pessoa)

	mepe["Araujo"]= pessoa{
		Nome:      "Davi",
		Sobrenome: "Araujo",
		sabores:   []string{"biscoito", "chocolate"},
	}

	fmt.Println(mepe)

	for _, v := range mepe {
		fmt.Println(v)
	}
}
