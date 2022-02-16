package main

import (
	"fmt"
	"sort"
)

func main() {
	ss := []string{"c", "b", "a", "e", "d"}
	si := []int{9, 8, 5, 6, 7, 4, 25, 36, 58}

	fmt.Println(si)
	sort.Ints(si)
	fmt.Println(si)
	fmt.Print("============/////=========\n")
	fmt.Println(ss)
	sort.Strings(ss)
	fmt.Println(ss)

}
