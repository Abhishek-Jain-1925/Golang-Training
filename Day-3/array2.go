package main

import (
	"fmt"
)

func main() {
	fmt.Println()
	var no int
	fmt.Print("How many Elements You want in Array : ")
	fmt.Scanf("%d", &no)
	var arr1 = make([]int, no)

	for i := 0; i < no; i++ {
		fmt.Printf("\nEnter no %d:", i+1)
		fmt.Scanf("%d", &arr1[i])
	}

	fmt.Print("\nYou Have Entered : ")
	for i := 0; i < no; i++ {
		fmt.Printf("\t%d", arr1[i])
	}

}
