package main

import "fmt"

func main() {
	ss := [][][]int{
		[][]int{
			[]int{1, 2, 3, 4},
		},
		[][]int{
			[]int{5, 6, 7, 8},
		},
		[][]int{
			[]int{9, 10, 11, 12},
		},
	}

	fmt.Println(ss[0][0][0])
}
