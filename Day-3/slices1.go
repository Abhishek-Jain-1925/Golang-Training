package main

import "fmt"

func main() {
	var s1 = []int{1, 2, 3, 4}
	fmt.Println(s1)

	arr1 := [3]int{1, 2, 3}
	var s2 []int
	s2 = arr1[:]
	s2[1] = 20
	fmt.Println(arr1, s2)

	//By Modifying size of slices it cant affect on array but not in above case
	s2 = append(s2, 4)
	fmt.Println(arr1, s2)

}
