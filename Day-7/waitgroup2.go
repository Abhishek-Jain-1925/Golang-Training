package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup //For Synchronization

func main() {
	fmt.Println("Im in Main")

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(inp int) {
			fmt.Println(inp)
			//wg.Done()
		}(i)
	}
	wg.Wait() //wait for all goroutines to end their execution and then call below statements

	fmt.Println("end of main !!")
}

//As Done is not called no(add) > no(done) So,Deadlock Occurred !!
