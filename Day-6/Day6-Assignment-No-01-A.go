package main

import "fmt"

func accessSlice(slice []int, index int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("internal error:", r)
		}
	}()

	value := slice[index]
	fmt.Printf("item %d, value %d\n", index, value)
}

func main() {
	mySlice := []int{1, 2, 3, 4, 5}

	accessSlice(mySlice, 2)

	accessSlice(mySlice, 10)
}
