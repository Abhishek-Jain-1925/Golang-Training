package main

import "fmt"

type rectangle struct {
	length  int
	breadth int
}

func (area rectangle) Area() int {
	return area.length * area.breadth
}

func (perimeter rectangle) Perimeter() int {
	return 2 * (perimeter.length + perimeter.breadth)
}

func main() {
	var length, breadth int
	fmt.Println("Enter Length And Breadth in Sequence Respectively : ")
	fmt.Scanf("%d%d", &length, &breadth)

	val := rectangle{
		length:  length,
		breadth: breadth,
	}

	fmt.Println("Area Of Rectangle : ", val.Area(), "\n Perimeter of Rectangle : ", val.Perimeter())

}
