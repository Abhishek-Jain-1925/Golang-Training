package main

import (
	"fmt"
)

type stud struct {
	sr_no  int
	name   string
	height float32
}

func main() {
	fmt.Println()

	var s stud

	s.sr_no = 1925
	s.name = "Mr. Abhishek Jain"
	s.height = 156.25

	//fmt.Printf("Details of Student : %v", s)
	fmt.Printf("Details of Student : %#v", s)
	fmt.Printf("\nDetails of Student : Sr.No : %d, Name : %s, Height : %f\n\n", s.sr_no, s.name, s.height)

	var st []stud

	st[1].sr_no = 2
	st[1].name = "Mr.PrajjwalKumar Paanzade"
	st[1].height = 145.90

	st[2].sr_no = 3
	st[2].name = "Mr. Yash Khartade"
	st[2].height = 160.60

	fmt.Println(len(st))
	// for i := 0; i < len(st); i++ {
	// 	fmt.Printf("\n %#v", st[i])
	// }

}
