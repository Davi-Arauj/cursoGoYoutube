package main

import "fmt"

func main() {
	x := 387

	func(x int) {
		fmt.Println(x, "vezes 87345 é :")
		fmt.Println(x * 87345)

	}(x)
}
