package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Seu OS é: ", runtime.GOOS)
	fmt.Println("Sua ARCH é: ", runtime.GOARCH)
}