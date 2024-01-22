package main

import "fmt"

func fact(n int) (result int) {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	var val int
	fmt.Println("Enter Value to find Factorial : ")
	_, err := fmt.Scan(&val)

	if err != nil {
		fmt.Println(fmt.Errorf("Enter Valid Input !!"))
		return
	}
	fmt.Println("Factorial : ", fact(val))
}
