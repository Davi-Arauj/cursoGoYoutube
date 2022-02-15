package main

import "fmt"

func main() {
	amigos := map[string]int{
		"alfredo":5555555,
		"joana":4444444,
	}

	// fmt.Println(amigos)
	// fmt.Println(amigos["joana"])


	amigos["gopher"] = 333

	// fmt.Println(amigos)

	// fmt.Println(amigos["joana"])


	//comma ok idiom
	if sera, ok := amigos["joana"]; !ok{
		fmt.Println("não tem!")
		}else{
			fmt.Println(sera)
		}
}