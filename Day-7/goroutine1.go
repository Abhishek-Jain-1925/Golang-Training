package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello World !! \n")

	//Every Prog Always should have atleast one goroutine by default => main()

	//Creating Single goroutine

	time.Sleep(2 * time.Second)

	go func() {
		fmt.Println("I'm in Go-Routine and Executing Independantly !!")
	}()

	time.Sleep(1 * time.Second)
}
