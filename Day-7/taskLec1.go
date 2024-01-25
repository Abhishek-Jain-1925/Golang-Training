package main

import (
	"fmt"
	"sync"
)

func main() {
	a := make(chan int)
	wg := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			a <- i
		}(i)

	}

	for v := range a {
		fmt.Println(v)
		wg.Done()
	}
	wg.Wait()
	close(a)

	var m sync.Mutex
	m.

}
