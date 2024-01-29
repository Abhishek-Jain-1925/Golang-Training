package main

import (
	"fmt"
)

type myData int

func (d1 myData) multiply(d2 myData) myData {
	return d1 * d2
}

func main() {
	val1 := myData(10)
	val2 := myData(20)
	fmt.Println(val1.multiply(val2))
}
