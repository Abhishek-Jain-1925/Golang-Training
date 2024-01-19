package main

import (
	"fmt"
)

func main() {

	arr := []string{"qwe", "wer", "ert", "rty", "tyu", "yui", "uio", "iop"}

	var index1, index2 int
	fmt.Print("Enter index1 : ")
	fmt.Scanln(&index1)

	fmt.Print("Enter index2 : ")
	fmt.Scanln(&index2)

	if index1 < 0 || index2 < 0 || index2 >= len(arr) || index1 >= len(arr) {
		fmt.Println("Incorrect Indexes")
	} else {

		var result1, result2, result3 []string
		result1 = arr[:index1+1]
		result2 = arr[index1 : index2+1]
		result3 = arr[index2:]

		fmt.Println(result1)
		fmt.Println(result2)
		fmt.Println(result3)
	}
}
