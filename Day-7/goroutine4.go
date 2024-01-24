package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Im in Main")

	go func() {
		fmt.Println("1")
	}()

	go func() {
		fmt.Println("2")
	}()

	go func() {
		fmt.Println("3")
	}()

	go func() {
		fmt.Println("4")
	}()
	time.Sleep(1 * time.Second)
	fmt.Println("end of main")
}
