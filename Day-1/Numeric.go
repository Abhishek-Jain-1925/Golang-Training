package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println("Jay Jinendra !!")

	var a int = 1925
	b := 2225

	fmt.Println("The Value of A =", a, " and Value of B=", b)
	fmt.Printf("The Type of A is %T and Size of A is %d", a, unsafe.Sizeof(a))

	c := 10.10
	fmt.Printf("\n Value of C = %f and type of c = %T", c, c)

	d := 2 + 4i // or d := complex(2,4)
	fmt.Printf("\n Value of D = %f and type of D = %T", d, d)

	//Rounding
	e := 122.1000000
	no := fmt.Sprintf(" %.2f", e)
	fmt.Println(no)
}
