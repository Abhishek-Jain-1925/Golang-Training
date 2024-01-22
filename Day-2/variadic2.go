package main

import (
	"fmt"
)

func FindFactorial(values ...int) (result int, err error) {
	fact := 1

	if fact == 0 {
		err = fmt.Errorf("Factorial of Zero is One !!")
		return
	}
	for _, val := range values {
		fact *= val
	}
	return fact, err
}

func main() {
	//fmt.Println(FindFactorial(10, 20, 30))
	fmt.Println(FindFactorial(1, 2, 3, 4, 5))
}
