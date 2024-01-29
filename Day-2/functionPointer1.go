package main

import "fmt"

func add5(val *int) int {
	*val = *val + 5
	return *val
}

// Pass by Ref.
func main() {
	fmt.Println("\n Hello !!")

	var n int = 10

	fmt.Println("Value of n : ", n)

	add5(&n)

	fmt.Println("Value of n : ", n)

}
