package main

import "fmt"

//In for loop, initial and last Condition is optional. act as While loop
func main() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
