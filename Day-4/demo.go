package main

import "fmt"

type myType interface{}

func main() {
	var value myType
	var v int
	fmt.Println("Enter Any Value You Want : ")
	fmt.Scanf("%d%s", &v, &value)

	//output := fmt.Sprintf("You Have Entered Value : %v and Type : %T", value, value)
	fmt.Println(value, v)
}
