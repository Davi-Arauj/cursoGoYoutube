package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup
var totalgoroutines = 100000
var contador int

func main() {

	criarGoRoutines(totalgoroutines)
	wg.Wait()
	fmt.Println("Total de goroutines:\t", totalgoroutines, "\nTotal do contador\t", contador)
}

func criarGoRoutines(i int) {
	wg.Add(i)
	for j := 0; j < i; j++ {
		go func() {
			v := contador
			runtime.Gosched()
			v++
			contador = v
			wg.Done()

		}()

	}

}
