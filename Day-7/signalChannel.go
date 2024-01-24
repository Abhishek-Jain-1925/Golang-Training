package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Im in main !!")

	signalChannel := make(chan struct{})

	go func() {
		fmt.Println("Hello World !!")
		fmt.Println("Im Closing the channel now...")

		<-signalChannel
	}()

	time.Sleep(time.Second)
	fmt.Println("end of main !!")

}
