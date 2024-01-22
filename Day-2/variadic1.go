package main

import (
	"fmt"
	"strconv"
	"strings"
)

func FindSum(values ...string) (result int, err error) {
	sum := 0
	for _, val := range values {
		value, _ := strconv.Atoi(val)
		sum += value
	}
	return sum, err
}

func main() {
	fmt.Println("Hello World !! Welcome in GO World !!")

	var inp string
	fmt.Println("Enter Your Numbers to Sum Seperated by Comma : ")
	fmt.Scan(&inp)

	values := strings.Split(inp, ",")

	result, err := FindSum(values...)

	if err != nil {
		fmt.Println(fmt.Errorf("Error !! Plz,Try Again With valid Credentials !!"))
		return
	}
	fmt.Println("Sum of Entered Digits : ", result)
}
