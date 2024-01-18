package main

import "fmt"

func main() {
	var arr1 []int

	arr1[0] = 1
	arr1[1] = 2
	fmt.Println(arr1)

	//Define and  Initialize
	var arr2 = []int{1, 2}
	fmt.Println(arr2)

	arr3 := [3]int{1, 2, 3}
	fmt.Println(arr3)

	//Compiler will calculate len
	arr4 := [...]int{1, 2, 3, 4, 4, 5, 4, 6}
	fmt.Println(arr4)
}
