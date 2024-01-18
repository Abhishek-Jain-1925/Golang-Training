package main

import (
	"fmt"
	"math"
)

func main() {
	//var radius float64 = 5.566
	var radius float64
	fmt.Println("Enter Radius of the Circle : ")
	fmt.Scan(&radius)

	area := math.Pi * math.Pow(radius, 2)

	result := fmt.Sprintf("%.2f", area)

	fmt.Println("Area Of Circle : ", result)
}
