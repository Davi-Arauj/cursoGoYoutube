package main

import "fmt"

func main() {
	slice := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	fmt.Println(slice)
	outraSlice := append(slice[:3],slice[6:]...)

	fmt.Println(outraSlice)

}
