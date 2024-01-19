package main

import "fmt"

type rectangle struct {
	length  int
	breadth int
}

func (area rectangle) Area() int {
	return area.length * area.breadth
}

func (perimeter rectangle) Perimeter() (result int) {
	result = 2 * (perimeter.length + perimeter.breadth)
	return
}

func main() {
	var length, breadth int
	fmt.Println("Enter Length And Width For Rectangle in Resp. Sequence : ")
	fmt.Scanf("%d%d", &length, &breadth)

	val := rectangle{
		length:  length,
		breadth: breadth,
	}

	fmt.Println("Area Of Rectangle : ", val.Area(), "\n Perimeter of Rectangle : ", val.Perimeter())
}
