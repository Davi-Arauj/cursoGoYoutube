package main

import "fmt"

func main() {

	umaSlice := []int{1, 2, 3, 4}
	outraSlice := []int{9, 10, 11, 12}
	fmt.Println(umaSlice)
	umaSlice = append(umaSlice, 5, 6, 7, 8)
	fmt.Println(umaSlice)
	umaSlice = append(umaSlice, outraSlice...)
	fmt.Println(umaSlice)
}
