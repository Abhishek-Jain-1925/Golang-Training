package main

import (
	"fmt"
)

func main() {
	arr := [3][3]string{{"C #", "C", "Python"}, {"Java", "Scala", "Perl"},
		{"C++", "Go", "HTML"}}

	fmt.Println("Elements of Array 1")
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {

			fmt.Print("\t", arr[x][y])
		}
		fmt.Println()
	}
}
