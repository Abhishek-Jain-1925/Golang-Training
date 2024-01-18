package main

import "fmt"

func main() {
	var m1 = map[string]string{
		"k1": "v1",
	}

	fmt.Println(m1)

	m2 := map[string]string{
		"k1": "v1",
		"k2": "v2",
	}
	fmt.Println(m2)

	//Using make keyword

	m3 := make(map[int]string, 0)
	m3[0] = "v3"
	m3[1] = "v4"

	fmt.Println(m3[1])

	var m4 = map[int]int{
		1: 2,
		3: 4,
		5: 6,
	}
	fmt.Println(m4)
}
