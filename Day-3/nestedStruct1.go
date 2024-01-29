package main

import (
	"fmt"
)

type addr struct {
	village string
	tal     string
	dist    string
	pincode int64
}

type student struct {
	rollno  int
	name    string
	address addr
}

func main() {
	fmt.Println()
	aavej := student{
		rollno:  4264,
		name:    "Mr.Aavej Shaikh",
		address: addr{"Baramati", "Baramati", "Pune", 413102},
	}

	fmt.Printf("Student Info : %#v", aavej)
}
