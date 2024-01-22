package main

import "fmt"

func Divide(a, b int) (result int, err error) {
	if b == 0 {
		return 0, fmt.Errorf("Can't Divide By Zero !!")
	}
	return b / a, nil
}

func main() {
	val, err := Divide(10, 0)
	if err != nil {
		fmt.Println("Error Occured !! : ", err)
		return
	}
	fmt.Println("Result : ", val)
}
