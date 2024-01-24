package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("Im in Main")

	ch := make(chan int) //unbuffered channel
	defer close(ch)

	go func() {
		wg.Add(1)

		v := <-ch
		time.Sleep(2 * time.Second)

		fmt.Println(v)
		//fmt.Println(<-ch)   //Not Working
		ch <- 2

		wg.Done()
	}()

	go func() {
		wg.Add(1)

		ch <- 5
		//fmt.Println(<-ch)   //Not Working
		v := <-ch
		time.Sleep(2 * time.Second)

		fmt.Println(v)

		wg.Done()
	}()

	wg.Wait()
	time.Sleep(5 * time.Second)

	fmt.Println("end of main")
}
