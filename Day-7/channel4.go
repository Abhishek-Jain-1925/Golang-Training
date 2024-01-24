//Buffered Channel

package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Welcome in demo of buffered channel")
	var wg sync.WaitGroup

	bufChan := make(chan int, 5)

	wg.Add(1)
	go func() {
		bufChan <- 1
		bufChan <- 2
		bufChan <- 3
		bufChan <- 4
		bufChan <- 5

		wg.Done()
	}()

	wg.Add(1)
	go func() {
		for val := range bufChan {
			fmt.Println(val)
		}

		wg.Done()
	}()

	wg.Wait()
	//time.Sleep(time.Second * 5)
}
