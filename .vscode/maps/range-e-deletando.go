package main

import "fmt"

func main() {
	qualquercoisa := map[int]string{
		123: "muito legal",
		98:  "menos legal um pouquinho",
		983: "esse é massa",
		19:  "idade de ir pra cadeia",
	}

	fmt.Println(qualquercoisa)

	delete(qualquercoisa, 19)

	fmt.Println(qualquercoisa)
}
