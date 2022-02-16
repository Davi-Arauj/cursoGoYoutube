package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {

	totalgoroutines := 500

	wg.Add(totalgoroutines)

	go route01(totalgoroutines)

	wg.Wait()
}

func route01(totalgoroutines int) {
	for i := 0; i < totalgoroutines; i++ {
		fmt.Println("A go routine é ->:", i)
		runtime.Gosched()
	}

	wg.Done()
}
