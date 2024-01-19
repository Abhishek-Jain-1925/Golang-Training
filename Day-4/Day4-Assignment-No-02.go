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

	fmt.Println("Enter Length And Width For Rectangle in Resp. Sequence : ")

	_, err := fmt.Scanf("%d%d", &length, &breadth)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	if length <= 0 || breadth <= 0 {
		fmt.Println("Error: Length and breadth must be positive values")
		return
	}

	val := rectangle{
		length:  length,
		breadth: breadth,
	}

	fmt.Println("Area Of Rectangle:", val.Area(), "\nPerimeter of Rectangle:", val.Perimeter())
}
