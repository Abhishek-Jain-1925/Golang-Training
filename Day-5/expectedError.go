package main

import (
	"errors"
	"fmt"
)

// Expected Error
var DivideByZero = errors.New("Can't Divide By Zero !!")

func Divide(a, b int) (result int, err error) {
	if b == 0 {
		return 0, DivideByZero
	}
	return b / a, nil
}

func main() {
	fmt.Println()

	//result, err := Divide(10, 0)
	result, err := Divide(10, 10)
	if err != nil {
		switch {
		case errors.Is(err, DivideByZero):
			fmt.Println("Expected Error !!")
		default:
			fmt.Println("Unexpected Error !!")
		}
		return
	}
	fmt.Println("Result : ", result)
}
