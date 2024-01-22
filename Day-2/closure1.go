package main

import (
	"fmt"
)

func sayHello() {
	fmt.Println("I'm In sayHello func for Printing Closure Func After Me :)")

	//Closure Func
	func() {
		fmt.Println("\t Im In Closure Function !! :) ")
	}()

	fmt.Println("I'm In sayHello func for Printing Closure Func Before Me :(")
}

func main() {
	sayHello()
}
