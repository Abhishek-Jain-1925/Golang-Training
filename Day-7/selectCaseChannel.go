package main

import (
	"fmt"
)

func main() {

	fmt.Println("Im in main")

	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		//time.Sleep(time.Second)
		ch1 <- "Task 1 completed"
	}()

	go func() {
		//time.Sleep(2 * time.Second)
		ch2 <- "Task 2 completed"
	}()

	select {
	case msg1 := <-ch1:
		fmt.Println(msg1)
	case msg2 := <-ch2:
		fmt.Println(msg2)
	}

	fmt.Println("end of main")
}
