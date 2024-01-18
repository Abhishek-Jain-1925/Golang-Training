package main

import "fmt"

func main() {
	fmt.Println("Hello World...!!")

	var a bool = true
	b := false

	c, d := true, false //Also be valid declaration.

	fmt.Println("The Value of A=", a, " and Value of B=", b)

	fmt.Println("Now Make AND Gate by A && B : ", (a && b))

	fmt.Println("Now Make OR Gate by A || B : ", (a || b))

	fmt.Printf("The Value of C = %b, D = %b", c, d)
}
