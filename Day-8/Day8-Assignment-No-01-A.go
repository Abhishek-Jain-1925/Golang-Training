package main

import (
	"fmt"
	"sync"
	"time"
)

func isEven(n int) bool {
	return n%2 == 0
}

var m sync.Mutex

func main() {
	n := 0

	go func() {
		//Change in Code Forward ***
		m.Lock()
		n--
		nIsEven := isEven(n)
		m.Unlock()
		//till here ***

		time.Sleep(5 * time.Millisecond)
		if nIsEven {
			fmt.Println(n, " is even")
			return
		}
		fmt.Println(n, "is odd")
	}()

	go func() {
		n++
	}()

	// just waiting for the goroutines to finish before exiting
	time.Sleep(time.Second)
}
