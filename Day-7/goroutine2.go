package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello World !!")

	go func() {
		fmt.Println("im inside goroutine 1 !! \n")

		for i := 0; i < 10; i++ {
			go func(i int) {
				fmt.Println(i)
			}(i)
		}
	}()
	time.Sleep(time.Second * 2)

	fmt.Println("End of Main !!")
}
