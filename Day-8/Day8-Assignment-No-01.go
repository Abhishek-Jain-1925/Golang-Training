package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func isEven(n int) bool {
	return n%2 == 0
}

func main() {
	n := 0

	wg.Add(1)
	go func() {
		defer wg.Done()
		nIsEven := isEven(n)
		time.Sleep(5 * time.Millisecond)
		if nIsEven {
			fmt.Println(n, " is even")
			return
		}
		fmt.Println(n, "is odd")
	}()

	go func() {
		wg.Wait()
		n++
	}()

	// just waiting for the goroutines to finish before exiting
	time.Sleep(time.Second)
}
