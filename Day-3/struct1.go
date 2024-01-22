package main

import "fmt"

type myStruct struct {
	name   string
	rollno int
	mark   float64
}

func main() {
	fmt.Println()

	s := myStruct{
		name:   "Abhishek Jain",
		rollno: 1925,
		mark:   92.70,
	}

	fmt.Printf("You Have Entered : %#v", s)

}
