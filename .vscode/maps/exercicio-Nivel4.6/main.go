package main

import "fmt"

func main() {

	ss := [][]string{
		[]string{
			"miu",
			"milton",
			"encher o saco",
		},
		[]string{
			"mimi",
			"martha",
			"pedir comida",
		},
		[]string{
			"meus alunnos",
			"estudam",
			"exercicios",
		},
	}

	for _, v := range ss {
		fmt.Println(v[0])
		for _, item := range v {
			fmt.Println("\t",item)
		}
	}

}
