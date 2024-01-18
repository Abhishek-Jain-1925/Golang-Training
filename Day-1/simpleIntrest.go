package main

import (
	"fmt"
	"strings"
)

func main() {
	// principle,rate,time = 450000.00,9.5,6.6
	fmt.Print("Enter 3 Values Seperated by Comma : ")

	var input string
	fmt.Scanln(&input)

	values := strings.Split(input, ",")

	var principle, rate, time float64
	fmt.Sscan(values[0], &principle)
	fmt.Sscan(values[1], &rate)
	fmt.Sscan(values[2], &time)

	si := (principle * rate * time) / 100
	result := fmt.Sprintf("%.2f", si)
	fmt.Println("Simple Intrest : ", result)
}
