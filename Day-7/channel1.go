package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("Im in Main")

	ch := make(chan int) //unbuffered channel
	defer close(ch)

	wg.Add(1)
	go func() {

		fmt.Println(<-ch)
		ch <- 2

		wg.Done()
	}()

	wg.Add(1)
	go func() {

		ch <- 5
		fmt.Println(<-ch)

		wg.Done()
	}()

	wg.Wait()

	fmt.Println("Number of goroutines :", runtime.NumGoroutine())
	fmt.Println("Number of CPUs :", runtime.NumCPU())

	fmt.Println("end of main")
}
