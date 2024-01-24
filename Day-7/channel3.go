//Buffered Channel

package main

import (
	"fmt"
	"sync"
	"time"
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
		fmt.Println(<-bufChan) //if we print one time then it will print only one value 1
		fmt.Println(<-bufChan)
		fmt.Println(<-bufChan)
		fmt.Println(<-bufChan)

		wg.Done()
	}()

	wg.Wait()
	time.Sleep(time.Second * 5)
}
