package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup //For Synchronization

func main() {
	fmt.Println("Im in Main")

	for i := 0; i < 3; i++ {
		wg.Add(1) //Tell about adding goroutine in wg
		go func(inp int) {
			fmt.Println(inp)
			wg.Done() //after that execution of goroutine is done
		}(i)
	}
	wg.Wait() //wait for all goroutines to end their execution and then call below statements

	fmt.Println("end of main !!")
}
